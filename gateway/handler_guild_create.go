package gateway

import (
	"github.com/Goscord/goscord/gateway/event"
)

type GuildCreateHandler struct{}

func (h *GuildCreateHandler) Handle(s *Session, data []byte) {
	ev, err := event.NewGuildCreate(s.rest, data)

	if err != nil {
		return
	}

	if _, ok := s.state.Guilds[ev.Data.Id]; !ok {
		s.state.AddGuild(ev.Data)
		s.bus.Publish("guildCreate", ev.Data)
	} else {
		s.state.AddGuild(ev.Data)
	}
}
