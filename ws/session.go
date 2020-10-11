package ws

import (
	"fmt"
	"sync"
	"github.com/Seyz123/yalis/rest"
	"github.com/Seyz123/yalis/ws/packet"
	"github.com/gorilla/websocket"
)

type Session struct {
	*sync.Mutex
	token string
	conn *websocket.Conn
	sessionID string
	close chan bool
}

func NewSession(token string) *Session {
	s := &Session{}

	conn, _, err := websocket.DefaultDialer.Dial(rest.GatewayUrl, nil)
	if err != nil {
		panic(err)
	}

	s.conn = conn
	s.token = token
	s.close = make(chan bool)

	return s
}

func (s *Session) Login() error {
	go func() {
		for {
			select {
			case _ = <-s.close:
				return
			break

			default:
				_, msg, err := s.conn.ReadMessage()

				if err != nil {
					panic(err)
				}

				s.handleWs(msg)
			break
			}
		}
	}()

	// ToDo : Identify

	return nil
}

func (s *Session) handleWs(msg []byte) {
	pk, err := packet.NewPacket(msg)

	if err != nil {
		panic(err)
	}

	opcode, event := pk.Opcode, pk.Event

	switch opcode {
	case packet.OpHello:
		hello, err := packet.NewHello(msg)

		if err != nil {
			panic(err)
		}
		
		fmt.Println("Got Hello")
		fmt.Println(hello.Data)
	break
	}

	if event != "" {
		// ToDo
	}
}

func (s *Session) Send(v interface{}) error {
	// ToDo
	return nil
}

func (s *Session) Close() {
	s.conn.Close()
	s.close <- true

	fmt.Println("Closed")
	// ToDo : Make this cleaner
}