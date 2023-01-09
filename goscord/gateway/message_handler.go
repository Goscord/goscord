package gateway

import "github.com/Goscord/goscord/goscord/gateway/event"

type MessageCreateHandler struct{}

func (_ *MessageCreateHandler) Handle(s *Session, data []byte) {
	ev, err := event.NewMessageCreate(s.rest, data)

	if err != nil {
		return
	}

	s.Bus().Publish("messageCreate", ev.Data)
}
