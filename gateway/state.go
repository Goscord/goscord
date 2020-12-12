package gateway

import (
	"errors"
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

func (s *State) UpdateGuild(g *discord.Guild) {
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

		/*
		if guild.Roles == nil {
			guild.Roles = g.Roles
		}

		if guild.Emojis == nil {
			guild.Emojis = g.Emojis
		}

		if guild.VoiceStates == nil {
			guild.VoiceStates = g.VoiceStates
		}
		*/

		*guild = *g
	}
}

func (s *State) RemoveGuild(guild *discord.Guild) {
	if guild, ok := s.guilds[guild.Id]; ok {
		s.RLock()
		defer s.RUnlock()

		delete(s.guilds, guild.Id)
	}
}

func (s *State) Guild(id string) (*discord.Guild, error) {
	s.RLock()
	defer s.RUnlock()

	if guild, ok := s.guilds[id]; ok {
		return guild, nil
	}

	return nil, errors.New("Guild not found")
}