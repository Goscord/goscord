package gateway

import (
	"github.com/Goscord/goscord/goscord/gateway/event"
)

type GuildBanRemoveHandler struct{}

func (_ *GuildBanRemoveHandler) Handle(s *Session, data []byte) {
	ev, err := event.NewGuildBanRemove(s.rest, data)

	if err != nil {
		return
	}

	guild, err := s.State().Guild(ev.Data.GuildId)
	user := ev.Data.User

	if err != nil {
		return
	}

	s.Bus().Publish("guildBanRemove", guild, user)
}
