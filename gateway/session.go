package gateway

import (
	"errors"
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/Goscord/goscord/discord"
	"github.com/Goscord/goscord/gateway/event"
	"github.com/Goscord/goscord/gateway/packet"
	"github.com/Goscord/goscord/rest"
	ev "github.com/asaskevich/EventBus"
	"github.com/gorilla/websocket"
)

type Session struct {
	sync.Mutex
	options           *Options
	status            *packet.UpdateStatus
	user              *discord.User
	rest              *rest.Client
	bus               *ev.EventBus
	state             *State
	connMu            sync.Mutex
	conn              *websocket.Conn
	sessionID         string
	heartbeatInterval time.Duration
	lastHeartbeatAck  time.Time
	lastHeartbeatSent time.Time
	lastSequence      int64

	Channel  *rest.ChannelHandler
	Emoji    *rest.EmojiHandler
	Guild    *rest.GuildHandler
	Invite   *rest.InviteHandler
	Template *rest.TemplateHandler
	User     *rest.UserHandler
	Voice    *rest.VoiceHandler
	Webhook  *rest.WebhookHandler

	handlers map[string]EventHandler
	close    chan bool
}

func NewSession(options *Options) *Session {
	s := &Session{}

	s.options = options
	s.status = packet.NewUpdateStatus(nil, "")
	s.rest = rest.NewClient(options.Token)
	s.bus = ev.New().(*ev.EventBus)
	s.state = NewState(s)
	s.close = make(chan bool)

	s.Channel = rest.NewChannelHandler(s.rest)
	s.Emoji = rest.NewEmojiHandler(s.rest)
	s.Guild = rest.NewGuildHandler(s.rest)
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
		event.EventReady:             &ReadyHandler{},
		event.EventResumed:           &ResumedHandler{},
		event.EventGuildCreate:       &GuildCreateHandler{},
		event.EventGuildUpdate:       &GuildUpdateHandler{},
		event.EventGuildDelete:       &GuildDeleteHandler{},
		event.EventGuildBanAdd:       &GuildBanAddHandler{},
		event.EventGuildBanRemove:    &GuildBanRemoveHandler{},
		event.EventGuildEmojisUpdate: &GuildEmojisUpdateHandler{},
		event.EventMessageCreate:     &MessageCreateHandler{},
		event.EventChannelCreate:     &ChannelCreateHandler{},
		event.EventChannelUpdate:     &ChannelUpdateHandler{},
		event.EventChannelDelete:     &ChannelDeleteHandler{},
		event.EventPresenceUpdate:    &PresenceUpdateHandler{},
		event.EventGuildMemberAdd:    &GuildMemberAddHandler{},
	}
}

func (s *Session) Login() error {
	header := http.Header{}
	header.Add("accept-encoding", "zlib")

	s.connMu.Lock()
	conn, _, err := websocket.DefaultDialer.Dial(rest.GatewayUrl, header)
	s.connMu.Unlock()

	if err != nil {
		return err
	}

	s.connMu.Lock()
	conn.SetCloseHandler(func(code int, text string) error {
		if code == 4004 {
			panic(errors.New("authentication failed"))
		}

		return nil
	})
	s.connMu.Unlock()

	s.Lock()
	s.conn = conn
	s.Unlock()

	s.connMu.Lock()
	_, msg, err := s.conn.ReadMessage()
	s.connMu.Unlock()

	if err != nil {
		return err
	}

	pk, err := s.onMessage(msg)

	if err != nil {
		return err
	} else if pk.Opcode != 10 {
		return errors.New("expecting op 10")
	}

	s.Lock()
	s.lastHeartbeatAck = time.Now().UTC()
	sessionID := s.sessionID
	sequence := s.lastSequence
	token := s.options.Token
	intents := s.options.Intents
	cclose := s.close
	s.Unlock()

	if sequence == 0 && sessionID == "" {
		identify := packet.NewIdentify(token, intents)

		if err = s.Send(identify); err != nil {
			return err
		}
	} else {
		resume := packet.NewResume(token, sessionID, sequence)

		if err = s.Send(resume); err != nil {
			return err
		}
	}

	go s.startHeartbeat(cclose)
	go s.listen(cclose)

	return nil
}

func (s *Session) onMessage(msg []byte) (*packet.Packet, error) {
	pk, err := packet.NewPacket(msg)

	if err != nil {
		return nil, err
	}

	opcode, e := pk.Opcode, pk.Event

	switch opcode {
	case packet.OpHello:
		hello, err := packet.NewHello(msg)

		if err != nil {
			return nil, err
		}

		s.Lock()
		s.heartbeatInterval = hello.Data.HeartbeatInterval
		s.Unlock()

	case packet.OpInvalidSession:
		s.Lock()
		s.sessionID = ""
		s.lastSequence = 0
		s.Unlock()

		s.Close()
		s.reconnect()

	case packet.OpReconnect:
		s.Close()
		s.reconnect()

	case packet.OpHeartbeatAck:
		s.Lock()
		s.lastHeartbeatAck = time.Now().UTC()
		s.Unlock()
	}

	if e != "" {
		s.Lock()
		s.lastSequence = pk.Sequence
		handler, exists := s.handlers[e]
		s.Unlock()

		if exists {
			go handler.Handle(s, msg)
		} else {
			fmt.Println("Unhandled event : " + e)
		}
	}

	return pk, nil
}

func (s *Session) startHeartbeat(closed <-chan bool) {
	s.Lock()
	heartbeatInterval := s.heartbeatInterval
	s.Unlock()

	ticker := time.NewTicker(heartbeatInterval * time.Millisecond)
	defer ticker.Stop()

	for {
		s.Lock()
		lastSequence := s.lastSequence
		lastHeartbeatAck := s.lastHeartbeatAck
		s.lastHeartbeatSent = time.Now().UTC()
		s.Unlock()

		heartbeat := packet.NewHeartbeat(lastSequence)
		err := s.Send(heartbeat)

		if err != nil || time.Now().UTC().Sub(lastHeartbeatAck) > (heartbeatInterval*5*time.Millisecond) {
			s.Close()
			s.reconnect()

			return
		}

		select {
		case <-ticker.C:
			// loop

		case <-closed:
			return
		}
	}
}

func (s *Session) listen(closed <-chan bool) {
	for {
		select {
		default:
			s.connMu.Lock()
			_, msg, err := s.conn.ReadMessage()
			s.connMu.Unlock()

			if err != nil {
				s.Close()
				s.reconnect()

				return
			}

			_, _ = s.onMessage(msg)

		case <-closed:
			return
		}
	}
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

		if wait > 600 {
			wait = 600
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
	defer s.Unlock()

	s.status.Data.Game = activity

	return s.Send(s.status)
}

func (s *Session) SetStatus(status string) error {
	s.Lock()
	defer s.Unlock()

	s.status.Data.Status = status

	return s.Send(s.status)
}

func (s *Session) UpdatePresence(status *packet.UpdateStatus) error {
	s.Lock()
	defer s.Unlock()

	s.status = status

	return s.Send(status)
}

func (s *Session) Latency() time.Duration {
	s.Lock()
	defer s.Unlock()

	return s.lastHeartbeatAck.Sub(s.lastHeartbeatSent)
}

func (s *Session) Close() {
	s.close <- true

	s.connMu.Lock()
	_ = s.conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
	s.connMu.Unlock()

	time.Sleep(1 * time.Second)

	_ = s.conn.Close()
}

func (s *Session) Bus() *ev.EventBus {
	s.Lock()
	defer s.Unlock()

	return s.bus
}

func (s *Session) Me() *discord.User {
	s.Lock()
	defer s.Unlock()

	return s.user
}

func (s *Session) State() *State {
	s.Lock()
	defer s.Unlock()

	return s.state
}

func (s *Session) On(ev string, fn interface{}) error {
	s.Lock()
	defer s.Unlock()

	return s.bus.SubscribeAsync(ev, fn, false)
}
