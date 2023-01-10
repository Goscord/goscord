package gateway

import (
	"errors"
	"fmt"
	"github.com/Goscord/goscord/goscord/discord"
	"github.com/Goscord/goscord/goscord/gateway/event"
	"github.com/Goscord/goscord/goscord/gateway/packet"
	"github.com/Goscord/goscord/goscord/rest"
	"io"
	"net"
	"sync"
	"syscall"
	"time"

	"github.com/goccy/go-json"

	ev "github.com/asaskevich/EventBus"
	"github.com/gorilla/websocket"
)

type Status int

const (
	StatusUnconnected Status = iota
	StatusConnecting
	StatusWaitingForHello
	StatusWaitingForReady
	StatusIdentifying
	StatusReady
	StatusResuming
	StatusDisconnected
)

type Session struct {
	sync.RWMutex

	options  *Options
	rest     *rest.Client
	presence *packet.PresenceUpdate
	user     *discord.User
	bus      *ev.EventBus
	state    *State
	status   Status

	// ws conn
	connMu   sync.Mutex
	conn     *websocket.Conn
	handlers map[string]EventHandler

	// Discord gateway fields
	sessionID         string
	heartbeatTicker   *time.Ticker
	heartbeatInterval time.Duration
	lastHeartbeatAck  time.Time
	lastHeartbeatSent time.Time
	lastSequence      int64

	// Rest handlers
	Application *rest.ApplicationHandler
	Channel     *rest.ChannelHandler
	Emoji       *rest.EmojiHandler
	Guild       *rest.GuildHandler
	Interaction *rest.InteractionHandler
	Invite      *rest.InviteHandler
	Template    *rest.TemplateHandler
	User        *rest.UserHandler
	Voice       *rest.VoiceHandler
	Webhook     *rest.WebhookHandler
}

func NewSession(options *Options) *Session {
	s := new(Session)

	s.options = options
	s.presence = packet.NewPresenceUpdate(nil, discord.StatusTypeOnline)
	s.user = new(discord.User)
	s.rest = rest.NewClient(options.Token)
	s.bus = ev.New().(*ev.EventBus)
	s.state = NewState(s)
	s.status = StatusUnconnected

	s.Application = rest.NewApplicationHandler(s.rest)
	s.Channel = rest.NewChannelHandler(s.rest)
	s.Emoji = rest.NewEmojiHandler(s.rest)
	s.Guild = rest.NewGuildHandler(s.rest)
	s.Interaction = rest.NewInteractionHandler(s.rest)
	s.Invite = rest.NewInviteHandler(s.rest)
	s.Template = rest.NewTemplateHandler(s.rest)
	s.User = rest.NewUserHandler(s.rest)
	s.Voice = rest.NewVoiceHandler(s.rest)
	s.Webhook = rest.NewWebhookHandler(s.rest)

	s.registerHandlers()

	return s
}

func (s *Session) registerHandlers() {
	s.handlers = map[string]EventHandler{
		event.EventReady:   &ReadyHandler{},
		event.EventResumed: &ResumedHandler{},
		// Application events
		event.EventApplicationCommandPermissionsUpdate: &ApplicationCommandPermissionsUpdateHandler{},
		// AutoModeration events
		event.EventAutoModerationRuleCreate:      &AutoModerationRuleCreateHandler{},
		event.EventAutoModerationRuleDelete:      &AutoModerationRuleDeleteHandler{},
		event.EventAutoModerationRuleUpdate:      &AutoModerationRuleUpdateHandler{},
		event.EventAutoModerationActionExecution: &AutoModerationActionExecutionHandler{},
		event.EventChannelCreate:                 &ChannelCreateHandler{},
		event.EventChannelUpdate:                 &ChannelUpdateHandler{},
		event.EventChannelDelete:                 &ChannelDeleteHandler{},
		event.EventChannelPinsUpdate:             &ChannelPinsUpdateHandler{},
		event.EventThreadCreate:                  &ThreadCreateHandler{},
		event.EventThreadUpdate:                  &ThreadUpdateHandler{},
		event.EventThreadDelete:                  &ThreadDeleteHandler{},
		event.EventThreadListSync:                &ThreadListSyncHandler{},
		event.EventThreadMemberUpdate:            &ThreadMemberUpdateHandler{},
		event.EventThreadMembersUpdate:           &ThreadMembersUpdateHandler{},
		// GuildStickersUpdate
		// GuildIntegrationsUpdate
		event.EventGuildMemberAdd:    &GuildMemberAddHandler{},
		event.EventGuildMemberRemove: &GuildMemberRemoveHandler{},
		event.EventGuildMemberUpdate: &GuildMemberUpdateHandler{},

		event.EventGuildCreate:       &GuildCreateHandler{},
		event.EventGuildUpdate:       &GuildUpdateHandler{},
		event.EventGuildDelete:       &GuildDeleteHandler{},
		event.EventGuildBanAdd:       &GuildBanAddHandler{},
		event.EventGuildBanRemove:    &GuildBanRemoveHandler{},
		event.EventGuildEmojisUpdate: &GuildEmojisUpdateHandler{},
		event.EventMessageCreate:     &MessageCreateHandler{},
		event.EventPresenceUpdate:    &PresenceUpdateHandler{},
		event.EventInteractionCreate: &InteractionCreateHandler{},
	}
}

func (s *Session) Login() error {
	s.connMu.Lock()
	defer s.connMu.Unlock()

	if s.conn != nil {
		return errors.New("session is already connected")
	}

	s.status = StatusConnecting
	s.lastHeartbeatSent = time.Now().UTC()

	conn, rs, err := websocket.DefaultDialer.Dial(rest.GatewayUrl, nil)
	if err != nil {
		body := "null"

		if rs != nil && rs.Body != nil {
			defer func() {
				_ = rs.Body.Close()
			}()

			rawBody, bErr := io.ReadAll(rs.Body)
			if bErr != nil {
				return err
			}

			body = string(rawBody)
		}

		return fmt.Errorf("error while connecting to the gateway : %s", body)
	}

	conn.SetCloseHandler(func(code int, text string) error {
		closeCode := packet.CloseEventCode(code)

		if !closeCode.ShouldReconnect() {
			panic(fmt.Errorf("error connecting to gateway : %d %s", code, text))
		}

		return nil
	})

	s.conn = conn
	s.status = StatusWaitingForHello

	go s.listen(conn)

	return nil
}

func (s *Session) listen(conn *websocket.Conn) {
loop:
	for {
		_, msg, err := conn.ReadMessage()

		if err != nil {
			s.connMu.Lock()
			sameConnection := s.conn == conn
			s.connMu.Unlock()

			if !sameConnection {
				return
			}

			reconnect := true

			if closeError, ok := err.(*websocket.CloseError); ok {
				closeCode := packet.CloseEventCode(closeError.Code)
				reconnect = closeCode.ShouldReconnect()
			} else if errors.Is(err, net.ErrClosed) {
				reconnect = false
			}

			s.CloseWithCode(websocket.CloseServiceRestart, "reconnecting")
			if reconnect {
				go s.reconnect()

				break loop
			}
		}

		pk, err := packet.NewPacket(msg)

		if err != nil {
			return
		}

		opcode, e := pk.Opcode, pk.Event

		switch opcode {
		case packet.OpHello:
			s.connMu.Lock()
			s.lastHeartbeatAck = time.Now().UTC()
			s.connMu.Unlock()

			hello, err := packet.NewHello(msg)

			if err != nil {
				return
			}

			go s.startHeartbeat()

			s.connMu.Lock()
			s.heartbeatInterval = time.Duration(hello.Data.HeartbeatInterval) * time.Millisecond
			lastSequence := s.lastSequence
			sessionID := s.sessionID

			token := s.options.Token
			intents := s.options.Intents
			s.connMu.Unlock()

			if lastSequence == 0 || sessionID == "" {
				s.connMu.Lock()
				s.status = StatusIdentifying
				s.connMu.Unlock()

				identify := packet.NewIdentify(token, intents)

				if err = s.Send(identify); err != nil {
					return
				}

				s.connMu.Lock()
				s.status = StatusWaitingForReady
				s.connMu.Unlock()
			} else {
				resume := packet.NewResume(token, sessionID, lastSequence)

				if err = s.Send(resume); err != nil {
					return
				}
			}

		case packet.OpDispatch:
			s.connMu.Lock()
			s.lastSequence = pk.Sequence
			s.connMu.Unlock()

			if e != "" {
				s.connMu.Lock()
				s.lastSequence = pk.Sequence
				handler, exists := s.handlers[e]
				s.connMu.Unlock()

				if exists {
					go handler.Handle(s, msg)
				} else {
					fmt.Println("Unhandled event : " + e)
				}
			}

		case packet.OpHeartbeat:
			s.sendHeartbeat()

		case packet.OpReconnect:
			s.CloseWithCode(websocket.CloseServiceRestart, "reconnecting")
			go s.reconnect()

			break loop

		case packet.OpInvalidSession:
			var shouldResume = false

			err = json.Unmarshal(pk.Data, &shouldResume)
			if err != nil {
				shouldResume = false
			}

			code := websocket.CloseNormalClosure
			if shouldResume {
				code = websocket.CloseServiceRestart
			} else {
				s.connMu.Lock()
				s.sessionID = ""
				s.lastSequence = 0
				s.connMu.Unlock()
			}

			s.CloseWithCode(code, "invalid session")

			go s.reconnect()

			break loop

		case packet.OpHeartbeatAck:
			s.connMu.Lock()
			s.lastHeartbeatAck = time.Now().UTC()
			s.connMu.Unlock()
		}
	}
}

func (s *Session) startHeartbeat() {
	s.connMu.Lock()
	heartbeatTicker := time.NewTicker(s.heartbeatInterval)
	s.heartbeatTicker = heartbeatTicker
	s.connMu.Unlock()

	defer heartbeatTicker.Stop()

	for range heartbeatTicker.C {
		s.sendHeartbeat()
	}
}

func (s *Session) sendHeartbeat() {
	s.connMu.Lock()
	lastSequence := s.lastSequence
	s.connMu.Unlock()

	heartbeat := packet.NewHeartbeat(lastSequence)

	if err := s.Send(heartbeat); err != nil {
		if errors.Is(err, syscall.EPIPE) {
			return
		}

		s.CloseWithCode(websocket.CloseServiceRestart, "heartbeat timeout")

		go s.reconnect()

		return
	}

	s.connMu.Lock()
	s.lastHeartbeatSent = time.Now().UTC()
	s.connMu.Unlock()
}

func (s *Session) reconnect() {
	wait := time.Duration(5)

	for {
		fmt.Println("Reconnecting")

		err := s.Login()

		if err == nil {
			// ToDo : Reconnect to voice connections

			fmt.Println("Reconnected")

			return
		}

		<-time.After(wait)

		wait *= 2

		if wait > 300 {
			wait = 300
		}
	}
}

func (s *Session) Send(v interface{}) error {
	s.connMu.Lock()
	defer s.connMu.Unlock()

	return s.conn.WriteJSON(v)
}

func (s *Session) SetActivity(activity *discord.Activity) error {
	s.Lock()
	s.presence.Data.Activities[0] = activity
	s.Unlock()

	s.RLock()
	defer s.RUnlock()

	return s.Send(s.presence)
}

func (s *Session) SetStatus(status discord.StatusType) error {
	s.Lock()
	s.presence.Data.Status = status
	s.Unlock()

	s.RLock()
	defer s.RUnlock()

	return s.Send(s.presence)
}

func (s *Session) UpdatePresence(status *packet.PresenceUpdate) error {
	s.Lock()
	s.presence = status
	s.Unlock()

	return s.Send(status)
}

func (s *Session) Latency() time.Duration {
	s.connMu.Lock()
	lastHeartbeatAck := s.lastHeartbeatAck
	lastHeartbeatSent := s.lastHeartbeatSent
	s.connMu.Unlock()

	return lastHeartbeatAck.Sub(lastHeartbeatSent)
}

func (s *Session) Close() {
	s.CloseWithCode(websocket.CloseNormalClosure, "Shutting down")
}

func (s *Session) CloseWithCode(code int, message string) {
	s.connMu.Lock()
	heartbeatTicker := s.heartbeatTicker
	s.connMu.Unlock()

	if heartbeatTicker != nil {
		heartbeatTicker.Stop()
		heartbeatTicker = nil
	}

	s.connMu.Lock()
	defer s.connMu.Unlock()

	if s.conn != nil {
		s.conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(code, message))

		_ = s.conn.Close()

		s.conn = nil

		if code == websocket.CloseNormalClosure || code == websocket.CloseGoingAway {
			s.sessionID = ""
			s.lastSequence = 0
		}
	}
}

func (s *Session) Bus() *ev.EventBus {
	s.RLock()
	defer s.RUnlock()

	return s.bus
}

func (s *Session) Me() *discord.User {
	s.RLock()
	defer s.RUnlock()

	return s.user
}

func (s *Session) State() *State {
	s.RLock()
	defer s.RUnlock()

	return s.state
}

func (s *Session) Status() Status {
	s.RLock()
	defer s.RUnlock()

	return s.status
}

func (s *Session) On(ev string, fn interface{}) error {
	s.Lock()
	defer s.Unlock()

	return s.bus.SubscribeAsync(ev, fn, false)
}
