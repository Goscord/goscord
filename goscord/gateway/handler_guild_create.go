package gateway

import (
	"github.com/Goscord/goscord/goscord/gateway/event"
)

type GuildCreateHandler struct{}

func (_ *GuildCreateHandler) Handle(s *Session, data []byte) {
	ev, err := event.NewGuildCreate(s.rest, data)

	if err != nil {
		return
	}

	s.State().AddGuild(ev.Data)

	s.Bus().Publish("guildCreate", ev.Data)
}
