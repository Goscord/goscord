package gateway

import (
	"github.com/Goscord/goscord/gateway/event"
)

type MessageCreateHandler struct{}

func (h *MessageCreateHandler) Handle(s *Session, data []byte) {
	ev, err := event.NewMessageCreate(s.rest, data)

	if err != nil {
		return
	}

	s.Bus().Publish("message", ev.Data)
}
