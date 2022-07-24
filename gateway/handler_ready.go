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

	s.Lock()
	s.user = ev.Data.User
	s.sessionID = ev.Data.SessionID
	s.Unlock()

	for _, guild := range ev.Data.Guilds {
		s.state.AddGuild(guild)
	}

	s.Bus().Publish("ready")
}
