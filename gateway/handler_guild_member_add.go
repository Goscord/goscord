package gateway

import (
	"github.com/Goscord/goscord/gateway/event"
)

type GuildMemberAddHandler struct{}

func (_ *GuildMemberAddHandler) Handle(s *Session, data []byte) {
	ev, err := event.NewGuildMemberAdd(s.rest, data)

	if err != nil {
		return
	}

	s.State().AddMember(ev.Data.GuildId, ev.Data)

	s.Bus().Publish("guildMemberAdd", ev.Data)
}
