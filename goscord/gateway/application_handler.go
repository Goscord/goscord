package gateway

import "github.com/Goscord/goscord/goscord/gateway/event"

type ApplicationCommandPermissionsUpdateHandler struct{}

func (_ *ApplicationCommandPermissionsUpdateHandler) Handle(s *Session, data []byte) {
	ev, err := event.NewApplicationCommandPermissionsUpdate(s.rest, data)

	if err != nil {
		return
	}

	s.Bus().Publish("applicationCommandPermissionsUpdate", ev.Data)
}
