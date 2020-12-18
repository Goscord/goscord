package gateway

import "github.com/Goscord/goscord/gateway/event"

type GuildUpdateHandler struct{}

func (h *GuildUpdateHandler) Handle(s *Session, data []byte) {
	ev, err := event.NewGuildUpdate(s.rest, data)

	if err != nil {
		return
	}

	s.state.AddGuild(ev.Data)

	s.bus.Publish("guildUpdate", ev.Data)
}
