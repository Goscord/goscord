package gateway

import "github.com/Goscord/goscord/gateway/event"

type PresenceUpdateHandler struct{}

func (_ *PresenceUpdateHandler) Handle(s *Session, Data []byte) {
	ev, err := event.NewPresenceUpdate(s.rest, Data)

	if err != nil {
		return
	}

	// ToDo : Need to handle guildMemberAdd and guildMemberRemove events

	s.bus.Publish("presenceUpdate", ev.Data)
}
