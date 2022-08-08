package gateway

import "github.com/Goscord/goscord/gateway/event"

type GuildUpdateHandler struct{}

func (_ *GuildUpdateHandler) Handle(s *Session, data []byte) {
	ev, err := event.NewGuildUpdate(s.rest, data)

	if err != nil {
		return
	}

	s.State().AddGuild(ev.Data)

	s.bus.Publish("guildUpdate", ev.Data)
}
