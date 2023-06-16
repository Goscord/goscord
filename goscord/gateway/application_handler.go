package gateway

import (
	"log"

	"github.com/Goscord/goscord/goscord/gateway/event"
)

const (
	EventApplicationCommandPermissionsUpdate = "application_command_permissions_update"
)

type ApplicationCommandPermissionsUpdateHandler struct{}

func (_ *ApplicationCommandPermissionsUpdateHandler) Handle(s *Session, data []byte) {
	ev, err := event.NewApplicationCommandPermissionsUpdate(s.rest, data)
	if err != nil {
		log.Println("Erreur lors de la création de l'événement ApplicationCommandPermissionsUpdate:", err)
		return
	}

	s.Publish(EventApplicationCommandPermissionsUpdate, ev.Data)
}

