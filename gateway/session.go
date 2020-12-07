package gateway

import (
	"errors"
	"fmt"
	"github.com/Goscord/goscord/gateway/event"
	"github.com/Goscord/goscord/gateway/packet"
	"github.com/Goscord/goscord/rest"
	"github.com/Goscord/goscord/user"
	ev "github.com/asaskevich/EventBus"
	"github.com/gorilla/websocket"
	"sync"
	"time"
)

type Session struct {
	sync.Mutex
	user              *user.User
	token             string
	rest              *rest.Client
	bus               *ev.EventBus
	connMu            sync.Mutex
	conn              *websocket.Conn
	sessionID         string
	heartbeatInterval time.Duration
	lastSequence      int64
	handlers          map[string]EventHandler
	close             chan bool
}

func NewSession(token string) *Session {
	s := &Session{}

	s.token = token
	s.rest = rest.NewClient(token)
	s.bus = ev.New().(*ev.EventBus)
	s.close = make(chan bool)

	s.registerHandlers()

	return s
}

func (s *Session) registerHandlers() {
	s.handlers = map[string]EventHandler{
		event.EventReady:         &ReadyHandler{},
		event.EventGuildCreate:   &GuildCreateHandler{},
		event.EventMessageCreate: &MessageCreateHandler{},
	}
}

func (s *Session) Login() error {
	conn, _, err := websocket.DefaultDialer.Dial(rest.GatewayUrl, nil)
	if err != nil {
		return err
	}

	conn.SetCloseHandler(func(code int, text string) error {
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
		return errors.New("Expecting op 10")
	}

	// ToDo : Handle heartbeat ack

	s.Lock()
	defer s.Unlock()

	sessionID := s.sessionID
	sequence := s.lastSequence

	if sequence == 0 && sessionID == "" {
		identify := packet.NewIdentify(s.token)

		if err = s.Send(identify); err != nil {
			return err
		}
	} else {
		resume := packet.NewResume(s.token, sessionID, sequence)

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
		defer s.Unlock()

		s.sessionID = ""
		s.lastSequence = 0

		s.Close()
		s.reconnect()

	case packet.OpReconnect:
		panic("Gateway want a reconnect")

	}

	if e != "" {
		s.Lock()
		s.lastSequence = pk.Sequence
		s.Unlock()

		handler, exists := s.handlers[e]

		if exists {
			handler.Handle(s, msg)
		} else {
			fmt.Println("Unhandled e : " + e)
		}
	}

	return pk, nil
}

func (s *Session) startHeartbeat() {
	s.Lock()
	ticker := time.NewTicker(s.heartbeatInterval)
	s.Unlock()

	defer ticker.Stop()

	for {
		heartbeat := packet.NewHeartbeat(s.lastSequence)

		err := s.Send(heartbeat)

		if err != nil {
			s.Close()
			s.reconnect()

			return
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
		_, msg, err := s.conn.ReadMessage()

		if err != nil {
			s.Close()
			s.reconnect()

			return
		}

		_, _ = s.onMessage(msg)
	}
}

func (s *Session) reconnect() {
	wait := time.Duration(5)

	for {
		fmt.Println("Reconnecting")

		err := s.Login()

		if err == nil {
			fmt.Println("Reconnected")

			// ToDo : Reconnect to voice connections

			return
		}

		fmt.Println("Retrying to reconnect...")

		<-time.After(wait * time.Second)
	}
}

func (s *Session) Send(v interface{}) error {
	s.connMu.Lock()
	defer s.connMu.Unlock()

	return s.conn.WriteJSON(v)
}

func (s *Session) Close() {
	_ = s.conn.Close()
	s.close <- true
}

func (s *Session) Bus() *ev.EventBus {
	return s.bus
}

func (s *Session) User() *user.User {
	return s.user
}

func (s *Session) On(ev string, fn interface{}) error {
	return s.bus.SubscribeAsync(ev, fn, false)
}