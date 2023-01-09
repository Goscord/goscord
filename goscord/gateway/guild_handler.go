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

type GuildUpdateHandler struct{}

func (_ *GuildUpdateHandler) Handle(s *Session, data []byte) {
	ev, err := event.NewGuildUpdate(s.rest, data)

	if err != nil {
		return
	}

	s.State().AddGuild(ev.Data)

	s.Bus().Publish("guildUpdate", ev.Data)
}

type GuildDeleteHandler struct{}

func (_ *GuildDeleteHandler) Handle(s *Session, data []byte) {
	ev, err := event.NewGuildDelete(s.rest, data)

	if err != nil {
		return
	}

	s.State().RemoveGuild(ev.Data)

	s.Bus().Publish("guildDelete", ev.Data)
}

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

type GuildEmojisUpdateHandler struct{}

func (_ *GuildEmojisUpdateHandler) Handle(s *Session, data []byte) {
	ev, err := event.NewGuildEmojisUpdate(s.rest, data)

	if err != nil {
		return
	}

	err = s.State().AddEmojis(ev.Data.GuildId, ev.Data.Emojis)
	if err != nil {
		return
	}

	s.Bus().Publish("guildEmojisUpdate", ev.Data)
}

// ToDo : EventGuildStickersUpdate

// ToDo : EventGuildIntegrationsUpdate

type GuildMemberAddHandler struct{}

func (_ *GuildMemberAddHandler) Handle(s *Session, data []byte) {
	ev, err := event.NewGuildMemberAdd(s.rest, data)

	if err != nil {
		return
	}

	s.State().AddMember(ev.Data.GuildId, ev.Data)

	s.Bus().Publish("guildMemberAdd", ev.Data)
}
