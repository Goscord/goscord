package ws

import (
	"fmt"
	"sync"
	"time"
	"github.com/Seyz123/yalis/rest"
	"github.com/Seyz123/yalis/ws/packet"
	"github.com/gorilla/websocket"
)

type Session struct {
	sync.Mutex
	token string
	connMu sync.Mutex
	conn *websocket.Conn
	sessionID string
	heartbeatInterval time.Duration
	lastSequence int
	close chan bool
}

func NewSession(token string) *Session {
	s := &Session{}

	s.token = token
	s.close = make(chan bool)

	return s
}

func (s *Session) Login() error {
    s.Lock()
	defer s.Unlock()
	
    conn, _, err := websocket.DefaultDialer.Dial(rest.GatewayUrl, nil)
	if err != nil {
		return err
	}
	
	conn.SetCloseHandler(s.onClose)

	s.conn = conn
	
	go func() {
		for {
			select {
			case _ = <-s.close:
				break
			break

			default:
				_, msg, err := s.conn.ReadMessage()

				if err != nil {
					return
				}

				s.onMessage(msg)
			break
			}
		}
	}()

	return nil
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

		break
	}

	if event != "" {
		fmt.Println("GOT EVENT : " + event)
	}
}

func (s *Session) onClose(code int, text string) error {
    // ToDo
    
    return nil
}

func (s *Session) startHeartbeat() {
    for {
        s.Lock()
        ticker := time.NewTicker(s.heartbeatInterval)
        s.Unlock()
        
        defer ticker.Stop()
        
        fmt.Println("Sending heartbeat...")
        
        heartbeat := packet.NewHeartbeat(s.lastSequence)
        
        err := s.Send(heartbeat)
        
        if err != nil {
            panic(err)
        }
        
        select {
            case <-ticker.C:
                // loop
            break
            
            case <-s.close:
                return
            break
        }
    }
}

func (s *Session) Send(v interface{}) error {
    s.connMu.Lock()
    defer s.connMu.Unlock()
    
    return s.conn.WriteJSON(v)
}

func (s *Session) Close() {
	s.conn.Close()
	s.close <- true

	fmt.Println("Closed")
	// ToDo : Make this cleaner
}