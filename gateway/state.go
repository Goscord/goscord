package gateway

import (
	"errors"
	"github.com/Goscord/goscord/discord"
	"sync"
)

type State struct {
	session  *Session
	mut      *sync.RWMutex
	Guilds   map[string]*discord.Guild
	Channels map[string]*discord.Channel
	Members  map[string][]*discord.Member
}

func NewState(session *Session) *State {
	return &State{
		session:  session,
		mut:       new(sync.RWMutex),
		Guilds:   map[string]*discord.Guild{},
		Channels: map[string]*discord.Channel{},
		Members:  map[string][]*discord.Member{},
	}
}

func (s *State) AddGuild(guild *discord.Guild) {
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

	s.mut.Lock()
	s.Guilds[guild.Id] = guild
	s.mut.Unlock()
}

func (s *State) UpdateGuild(guild *discord.Guild) {
	if guild.Channels != nil {
		for _, channel := range guild.Channels {
			s.AddChannel(channel)
		}
	}

	s.mut.Lock()
	s.Guilds[guild.Id] = guild
	s.mut.Unlock()
}

func (s *State) RemoveGuild(guild *discord.Guild) {
	if g, ok := s.Guilds[guild.Id]; ok {
		if guild.Channels != nil {
			for _, channel := range guild.Channels {
				s.RemoveChannel(channel)
			}
		}

		s.mut.Lock()
		delete(s.Guilds, g.Id)
		s.mut.Unlock()
	}
}

func (s *State) Guild(id string) (*discord.Guild, error) {
	s.mut.RLock()
	defer s.mut.RUnlock()

	if guild, ok := s.Guilds[id]; ok {
		return guild, nil
	}

	return nil, errors.New("Guild not found")
}

func (s *State) AddChannel(channel *discord.Channel) {
	if _, ok := s.Channels[channel.Id]; ok {
		s.UpdateChannel(channel)

		return
	}

	s.mut.Lock()
	s.Channels[channel.Id] = channel
	s.mut.Unlock()
}

func (s *State) UpdateChannel(channel *discord.Channel) {
	s.mut.Lock()
	s.Channels[channel.Id] = channel
	s.mut.Unlock()
}

func (s *State) RemoveChannel(channel *discord.Channel) {
	s.mut.Lock()
	defer s.mut.Unlock()

	if c, ok := s.Channels[channel.Id]; ok {
		delete(s.Channels, c.Id)
	}
}

func (s *State) Channel(id string) (*discord.Channel, error) {
	s.mut.RLock()

	if channel, ok := s.Channels[id]; ok {
		return channel, nil
	}

	s.mut.RUnlock()

	channel, _ := s.session.Channel.GetChannel(id)

	if channel != nil {
		s.AddChannel(channel)

		return channel, nil
	}

	return nil, errors.New("Channel not found")
}
