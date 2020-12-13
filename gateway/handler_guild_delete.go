package gateway

import (
	"github.com/Goscord/goscord/gateway/event"
)

type GuildDeleteHandler struct{}

func (h *GuildDeleteHandler) Handle(s *Session, data []byte) {
	ev, err := event.NewGuildDelete(s.rest, data)

	if err != nil {
		return
	}

	s.state.RemoveGuild(ev.Data)

	s.bus.Publish("guildDelete", ev.Data)
}
