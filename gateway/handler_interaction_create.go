package gateway

import (
	"github.com/Goscord/goscord/gateway/event"
)

type InteractionCreateHandler struct{}

func (_ *InteractionCreateHandler) Handle(s *Session, data []byte) {
	ev, err := event.NewInteractionCreate(s.rest, data)

	if err != nil {
		return
	}

	s.bus.Publish("interactionCreate", ev.Data)
}
