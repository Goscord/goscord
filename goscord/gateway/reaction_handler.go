package gateway

import (
	"github.com/Goscord/goscord/goscord/gateway/event"
)

type MessageReactionHandler struct{}

func (_ *MessageReactionHandler) Handle(s *Session, data []byte) {
	ev, err := event.NewMessageReactionAdd(s.rest, data)

	if err != nil {
		return
	}
	s.Publish(event.EventMessageReactionAdd, ev.Data)
}