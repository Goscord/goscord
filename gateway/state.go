package gateway

import (
	"github.com/Goscord/goscord/discord"
	"sync"
)

type State struct {
	sync.RWMutex

	guilds   map[string]*discord.Guild
	channels map[string]*discord.Channel
}

func NewState() *State {
	return &State{
		guilds: map[string]*discord.Guild{},
		channels: map[string]*discord.Channel{},
	}
}

func (s *State) AddGuild(guild *discord.Guild) {
	s.RLock()
	defer s.RUnlock()

	// TODO : Channels & members

	if _, ok := s.guilds[guild.Id]; ok {
		s.UpdateGuild(guild)

		return
	}

	s.guilds[guild.Id] = guild
}

func (s *State) RemoveGuild(guild *discord.Guild) {
	if guild, ok := s.guilds[guild.Id]; ok {
		s.RLock()
		defer s.RUnlock()

		delete(s.guilds, guild.Id)
	}
}

func (s *State) UpdateGuild(g *discord.Guild) {
	s.RLock()
	defer s.RUnlock()

	if guild, ok := s.guilds[g.Id]; ok {
		if guild.MemberCount == 0 {
			guild.MemberCount = g.MemberCount
		}

		if guild.Members == nil {
			guild.Members = g.Members
		}

		if guild.Channels == nil {
			guild.Channels = g.Channels
		}

		// TODO : Roles & emotes
	} else {
		s.AddGuild(guild)
	}
}
