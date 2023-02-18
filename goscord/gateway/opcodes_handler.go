package gateway

import "github.com/Goscord/goscord/goscord/gateway/event"

type ReadyHandler struct{}

func (_ *ReadyHandler) Handle(s *Session, data []byte) {
	ev, err := event.NewReady(data)

	if err != nil {
		return
	}

	s.connMu.Lock()
	s.user = ev.Data.User
	s.sessionID = ev.Data.SessionID
	s.status = StatusReady
	s.connMu.Unlock()

	for _, guild := range ev.Data.Guilds {
		s.State().AddGuild(guild)
	}

	s.Publish(event.EventReady)
}

type ResumedHandler struct{}

func (_ *ResumedHandler) Handle(s *Session, _ []byte) {
	s.Publish(event.EventResumed)
}
