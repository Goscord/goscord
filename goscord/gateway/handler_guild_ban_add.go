package gateway

import (
	"github.com/Goscord/goscord/goscord/gateway/event"
)

type GuildBanAddHandler struct{}

func (_ *GuildBanAddHandler) Handle(s *Session, data []byte) {
	ev, err := event.NewGuildBanAdd(s.rest, data)

	if err != nil {
		return
	}

	guild, err := s.State().Guild(ev.Data.GuildId)
	user := ev.Data.User

	if err != nil {
		return
	}

	s.Bus().Publish("guildBanAdd", guild, user)
}
