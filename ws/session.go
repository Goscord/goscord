package ws

import (
	"fmt"
	"sync"
	"time"
	"github.com/Seyz123/yalis/rest"
	"github.com/Seyz123/yalis/ws/packet"
	"github.com/Seyz123/yalis/ws/event"
	"github.com/gorilla/websocket"
	ev "github.com/asaskevich/EventBus"
)

type Session struct {
	sync.Mutex
	token string
	bus *ev.EventBus
	connMu sync.Mutex
	conn *websocket.Conn
	sessionID string
	heartbeatInterval time.Duration
	lastSequence int
	handlers map[string]EventHandler
	close chan bool
}

func NewSession(token string, bus *ev.EventBus) *Session {
	s := &Session{}

	s.token = token
	s.bus = bus
	s.close = make(chan bool)
	
	s.registerHandlers()

	return s
}

func (s *Session) registerHandlers() {
    s.handlers = map[string]EventHandler{
        event.EventReady: &ReadyHandler{},
    }
}

func (s *Session) onMessage(msg []byte) {
	pk, err := packet.NewPacket(msg)

	if err != nil {
		panic(err)
	}

	opcode, event := pk.Opcode, pk.Event
	
	s.Lock()
	s.lastSequence = pk.Sequence
	s.Unlock()

	switch opcode {
	case packet.OpHello:
		hello, err := packet.NewHello(msg)

		if err != nil {
			panic(err)
		}
		
		s.Lock()
		s.heartbeatInterval = hello.Data.HeartbeatInterval
		s.Unlock()
		
		go s.startHeartbeat()
		
		identify := packet.NewIdentify(s.token)
		
		if err := s.Send(identify); err != nil {
		    panic("Cannot identify")
		}
	}

	if event != "" {
		handler, exists := s.handlers[event]
		
		if exists {
		    handler.Handle(s, msg)
		} else {
		    fmt.Println("Unhandled event : " + event)
		}
	}
}

func (s *Session) startHeartbeat() {
    for {
        s.Lock()
        ticker := time.NewTicker(s.heartbeatInterval)
        s.Unlock()
        
        defer ticker.Stop()
        
        heartbeat := packet.NewHeartbeat(s.lastSequence)
        
        err := s.Send(heartbeat)
        
        if err != nil {
            // ToDo : Try resume session
        }
        
        select {
            case <-ticker.C:
                // loop
            
            case <-s.close:
                return
        }
    }
}

func (s *Session) Login() error {
    s.Lock()
	defer s.Unlock()
	
    conn, _, err := websocket.DefaultDialer.Dial(rest.GatewayUrl, nil)
	if err != nil {
		return err
	}
	
	conn.SetCloseHandler(func (code int, text string) error {
	    return nil
	})

	s.conn = conn
	
	go func() {
		for {
			select {
			case <-s.close:
				return

			default:
				_, msg, err := s.conn.ReadMessage()

				if err != nil {
					return
				}

				s.onMessage(msg)
			}
		}
	}()

	return nil
}

func (s *Session) Send(v interface{}) error {
    s.connMu.Lock()
    defer s.connMu.Unlock()
    
    return s.conn.WriteJSON(v)
}

func (s *Session) Close() {
	_ = s.conn.Close()
	s.close <- true

	fmt.Println("Connection closed")
}

func (s *Session) Bus() *ev.EventBus {
    return s.bus
}