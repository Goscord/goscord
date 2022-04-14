package gateway

import (
	"github.com/Goscord/goscord/gateway/event"
)

type ReadyHandler struct{}

func (_ *ReadyHandler) Handle(s *Session, data []byte) {
	ev, err := event.NewReady(data)

	if err != nil {
		return
	}

	s.user = ev.Data.User
	s.sessionID = ev.Data.SessionID

	for _, guild := range ev.Data.Guilds {
		s.state.AddGuild(guild)
	}

	s.Bus().Publish("ready")
}
