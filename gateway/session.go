package gateway

import (
	"errors"
	"fmt"
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
	sync.RWMutex
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
	}
}

func (s *Session) Login() error {
	conn, _, err := websocket.DefaultDialer.Dial(rest.GatewayUrl, nil)

	if err != nil {
		return err
	}

	conn.SetCloseHandler(func(code int, text string) error {
		if code == 4004 {
			panic(errors.New("authentication failed"))
		}

		return nil
	})

	s.conn = conn

	_, msg, err := s.conn.ReadMessage()

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
	s.Unlock()

	sessionID := s.sessionID
	sequence := s.lastSequence

	if sequence == 0 && sessionID == "" {
		identify := packet.NewIdentify(s.options.Token, s.options.Intents)

		if err = s.Send(identify); err != nil {
			return err
		}
	} else {
		resume := packet.NewResume(s.options.Token, sessionID, sequence)

		if err = s.Send(resume); err != nil {
			return err
		}
	}

	go s.startHeartbeat()
	go s.listen()

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
		s.Unlock()

		handler, exists := s.handlers[e]

		if exists {
			go handler.Handle(s, msg)
		} else {
			fmt.Println("Unhandled event : " + e)
		}
	}

	return pk, nil
}

func (s *Session) startHeartbeat() {
	s.RLock()
	heartbeatInterval := s.heartbeatInterval
	s.RUnlock()

	ticker := time.NewTicker(s.heartbeatInterval)
	defer ticker.Stop()

	for {
		s.RLock()
		lastSequence := s.lastSequence
		lastHeartbeatAck := s.lastHeartbeatAck
		s.RUnlock()

		heartbeat := packet.NewHeartbeat(lastSequence)

		err := s.Send(heartbeat)

		s.Lock()
		s.lastHeartbeatSent = time.Now().UTC()
		s.Unlock()

		if err != nil || time.Now().UTC().Sub(lastHeartbeatAck) > (heartbeatInterval*5*time.Millisecond) {
			s.Close()
			s.reconnect()
			break
		}

		select {
		case <-ticker.C:
			// loop

		case <-s.close:
			return
		}
	}
}

func (s *Session) listen() {
	for {
		select {
		default:
			_, msg, err := s.conn.ReadMessage()

			if err != nil {
				s.Close()
				s.reconnect()

				break
			}

			_, _ = s.onMessage(msg)

		case <-s.close:
			return
		}
	}
}

func (s *Session) reconnect() {
	wait := time.Duration(5)

	for {
		err := s.Login()

		if err == nil {
			// ToDo : Reconnect to voice connections

			break
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
	s.status.Data.Game = activity

	return s.Send(s.status)
}

func (s *Session) SetStatus(status string) error {
	s.status.Data.Status = status

	return s.Send(s.status)
}

func (s *Session) UpdatePresence(status *packet.UpdateStatus) error {
	s.status = status

	return s.Send(status)
}

func (s *Session) Latency() time.Duration {
	return s.lastHeartbeatAck.Sub(s.lastHeartbeatSent)
}

func (s *Session) Close() {
	_ = s.conn.Close()
	s.close <- true
}

func (s *Session) Bus() *ev.EventBus {
	return s.bus
}

func (s *Session) Me() *discord.User {
	return s.user
}

func (s *Session) State() *State {
	return s.state
}

func (s *Session) On(ev string, fn interface{}) error {
	return s.bus.SubscribeAsync(ev, fn, false)
}
