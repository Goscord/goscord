package gateway

import (
	"errors"
	"github.com/Goscord/goscord/discord"
	"sync"
)

type State struct {
	sync.RWMutex

	Guilds   map[string]*discord.Guild
	Channels map[string]*discord.Channel
	Members  map[string][]*discord.Member
}

func NewState() *State {
	return &State{
		Guilds:   map[string]*discord.Guild{},
		Channels: map[string]*discord.Channel{},
		Members:  map[string][]*discord.Member{},
	}
}

func (s *State) AddGuild(guild *discord.Guild) {
	s.Lock()
	defer s.Unlock()

	// TODO : Members

	if _, ok := s.Guilds[guild.Id]; ok {
		s.UpdateGuild(guild)

		return
	}

	if guild.Channels != nil {
		for _, channel := range guild.Channels {
			s.AddChannel(channel)
		}
	}

	s.Guilds[guild.Id] = guild
}

func (s *State) UpdateGuild(guild *discord.Guild) {
	s.Lock()
	defer s.Unlock()

	if guild.Channels != nil {
		for _, channel := range guild.Channels {
			s.AddChannel(channel)
		}
	}

	s.Guilds[guild.Id] = guild
}

func (s *State) RemoveGuild(guild *discord.Guild) {
	s.Lock()
	defer s.Unlock()

	if g, ok := s.Guilds[guild.Id]; ok {
		if guild.Channels != nil {
			for _, channel := range guild.Channels {
				s.RemoveChannel(channel)
			}
		}

		delete(s.Guilds, g.Id)
	}
}

func (s *State) Guild(id string) (*discord.Guild, error) {
	s.RLock()
	defer s.RUnlock()

	if guild, ok := s.Guilds[id]; ok {
		return guild, nil
	}

	return nil, errors.New("Guild not found")
}

func (s *State) AddChannel(channel *discord.Channel) {
	s.Lock()
	defer s.Unlock()

	if _, ok := s.Channels[channel.Id]; ok {
		s.UpdateChannel(channel)

		return
	}

	s.Channels[channel.Id] = channel
}

func (s *State) UpdateChannel(channel *discord.Channel) {
	s.Lock()
	defer s.Unlock()

	s.Channels[channel.Id] = channel
}

func (s *State) RemoveChannel(channel *discord.Channel) {
	s.Lock()
	defer s.Unlock()

	if c, ok := s.Channels[channel.Id]; ok {
		delete(s.Channels, c.Id)
	}
}

func (s *State) Channel(id string) (*discord.Channel, error) {
	s.RLock()
	defer s.RUnlock()

	if channel, ok := s.Channels[id]; ok {
		return channel, nil
	}

	return nil, errors.New("Channel not found")
}
