package gateway

import "github.com/Goscord/goscord/goscord/gateway/event"

type InteractionCreateHandler struct{}

func (_ *InteractionCreateHandler) Handle(s *Session, data []byte) {
	ev, err := event.NewInteractionCreate(s.rest, data)

	if err != nil {
		return
	}

	s.Bus().Publish("interactionCreate", ev.Data)
}
