package gateway

import "github.com/Goscord/goscord/gateway/event"

type GuildEmojisUpdateHandler struct{}

func (h *GuildEmojisUpdateHandler) Handle(s *Session, data []byte) {
	ev, err := event.NewGuildEmojisUpdate(s.rest, data)

	if err != nil {
		return
	}

	guild, err := s.state.Guild(ev.Data.GuildId)

	if err != nil {
		return
	}

	guild.Emojis = ev.Data.Emojis

	s.bus.Publish("guildEmojisUpdate", guild)
}
