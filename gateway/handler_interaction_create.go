package gateway

import (
	"fmt"

	"github.com/Goscord/goscord/gateway/event"
)

type InteractionCreateHandler struct{}

func (_ *InteractionCreateHandler) Handle(s *Session, data []byte) {
	ev, err := event.NewInteractionCreate(s.rest, data)

	if err != nil {
		fmt.Println(err)
		return
	}

	s.Bus().Publish("interactionCreate", ev.Data)
}
