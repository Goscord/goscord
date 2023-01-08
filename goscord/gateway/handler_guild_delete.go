package gateway

import (
	"github.com/Goscord/goscord/goscord/gateway/event"
)

type GuildDeleteHandler struct{}

func (_ *GuildDeleteHandler) Handle(s *Session, data []byte) {
	ev, err := event.NewGuildDelete(s.rest, data)

	if err != nil {
		return
	}

	s.State().RemoveGuild(ev.Data)

	s.Bus().Publish("guildDelete", ev.Data)
}
