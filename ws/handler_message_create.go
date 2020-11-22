package ws

import (
	"fmt"
	"github.com/Seyz123/yalis/ws/event"
)

type MessageCreateHandler struct{}

func (h *MessageCreateHandler) Handle(s *Session, data []byte) {
	ev, err := event.NewMessageCreate(data)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	s.Bus().Publish("message", ev.Data)
}
