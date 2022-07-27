package gateway

import (
	"github.com/Goscord/goscord/gateway/event"
)

type GuildMemberAddHandler struct{}

func (_ *GuildMemberAddHandler) Handle(s *Session, data []byte) {
	_, err := event.NewGuildMemberAdd(s.rest, data)

	if err != nil {
		return
	}

	// ToDo : Need some rework on State
}
