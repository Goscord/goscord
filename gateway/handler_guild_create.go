package gateway

import (
	"github.com/Goscord/goscord/gateway/event"
)

type GuildCreateHandler struct{}

func (_ *GuildCreateHandler) Handle(s *Session, data []byte) {
	ev, err := event.NewGuildCreate(s.rest, data)

	if err != nil {
		return
	}

	s.state.AddGuild(ev.Data)

	if _, err := s.state.Guild(ev.Data.Id); err != nil {
		s.bus.Publish("guildCreate", ev.Data)
	}
}
