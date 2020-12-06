package gateway

import (
	"github.com/Seyz123/yalis/gateway/event"
)

type MessageCreateHandler struct{}

func (h *MessageCreateHandler) Handle(s *Session, data []byte) {
	ev, err := event.NewMessageCreate(s.rest, data)

	if err != nil {
		return
	}

	s.Bus().Publish("message", ev.Data)
}
