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
	members map[string][]*discord.Member
}

func NewState() *State {
	return &State{
		guilds:   map[string]*discord.Guild{},
		channels: map[string]*discord.Channel{},
		members: map[string][]*discord.Member{},
	}
}

func (s *State) AddGuild(guild *discord.Guild) {
	s.Lock()
	defer s.Unlock()

	// TODO : Members

	if _, ok := s.guilds[guild.Id]; ok {
		s.UpdateGuild(guild)

		return
	}

	s.guilds[guild.Id] = guild
}

func (s *State) UpdateGuild(guild *discord.Guild) {
	s.Lock()
	defer s.Unlock()

	s.guilds[guild.Id] = guild
}

func (s *State) RemoveGuild(guild *discord.Guild) {
	s.Lock()
	defer s.Unlock()

	if g, ok := s.guilds[guild.Id]; ok {
		delete(s.guilds, g.Id)
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

func (s *State) AddChannel(channel *discord.Channel) {
	s.Lock()
	defer s.Unlock()

	if _, ok := s.channels[channel.Id]; ok {
		s.UpdateChannel(channel)

		return
	}

	s.channels[channel.Id] = channel
}

func (s *State) UpdateChannel(channel *discord.Channel) {
	s.Lock()
	defer s.Unlock()

	s.channels[channel.Id] = channel
}

func (s *State) RemoveChannel(id string) {
	s.Lock()
	defer s.Unlock()

	if c, ok := s.channels[id]; ok {
		delete(s.channels, c.Id)
	}
}

func (s *State) Channel(id string) (*discord.Channel, error) {
	s.RLock()
	defer s.RUnlock()

	if channel, ok := s.channels[id]; ok {
		return channel, nil
	}

	return nil, errors.New("Channel not found")
}