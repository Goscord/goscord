package gateway

import "github.com/Goscord/goscord/gateway/event"

type GuildBanRemoveHandler struct{}

func (h *GuildBanRemoveHandler) Handle(s *Session, data []byte) {
	ev, err := event.NewGuildBanRemove(s.rest, data)

	if err != nil {
		return
	}

	guild, err := s.state.Guild(ev.Data.GuildId)
	user := ev.Data.User

	if err != nil {
		return
	}

	s.bus.Publish("guildBanRemove", guild, user)
}
