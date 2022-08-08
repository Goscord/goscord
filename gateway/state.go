package gateway

import (
	"errors"
	"sync"

	"github.com/Goscord/goscord/discord"
)

type State struct {
	sync.RWMutex

	session  *Session
	guilds   map[string]*discord.Guild
	channels map[string]*discord.Channel
	members  map[string]map[string]*discord.GuildMember
}

func NewState(session *Session) *State {
	return &State{
		session:  session,
		guilds:   map[string]*discord.Guild{},
		channels: map[string]*discord.Channel{},
		members:  map[string]map[string]*discord.GuildMember{},
	}
}

// GUILDS

func (s *State) AddGuild(guild *discord.Guild) {
	// TODO : Members

	if guild.Channels != nil {
		for _, channel := range guild.Channels {
			s.AddChannel(channel)
		}
	}

	s.Lock()
	s.guilds[guild.Id] = guild
	s.Unlock()
}

func (s *State) RemoveGuild(guild *discord.Guild) {
	if g, err := s.Guild(guild.Id); err == nil {
		if guild.Channels != nil {
			for _, channel := range guild.Channels {
				s.RemoveChannel(channel)
			}
		}

		s.Lock()
		delete(s.guilds, g.Id)
		s.Unlock()
	}
}

func (s *State) Guild(id string) (*discord.Guild, error) {
	s.RLock()
	defer s.RUnlock()

	if guild, ok := s.guilds[id]; ok {
		return guild, nil
	}

	return nil, errors.New("guild not found")
}

// CHANNELS

func (s *State) AddChannel(channel *discord.Channel) {
	s.Lock()
	s.channels[channel.Id] = channel
	s.Unlock()
}

func (s *State) RemoveChannel(channel *discord.Channel) {
	if c, err := s.Channel(channel.Id); err == nil {
		s.Lock()
		delete(s.channels, c.Id)
		s.Unlock()
	}
}

func (s *State) Channel(id string) (*discord.Channel, error) {
	s.RLock()

	if channel, ok := s.channels[id]; ok {
		s.RUnlock()
		return channel, nil
	}

	s.RUnlock()

	channel, _ := s.session.Channel.GetChannel(id)

	if channel != nil {
		s.AddChannel(channel)

		return channel, nil
	}

	return nil, errors.New("channel not found")
}

// MEMBERS

func (s *State) AddMember(guildID string, member *discord.GuildMember) {
	if _, err := s.Guild(guildID); err != nil {
		return
	}

	s.Lock()

	if _, ok := s.members[guildID]; !ok {
		s.members[guildID] = map[string]*discord.GuildMember{}
	}

	s.members[guildID][member.User.Id] = member
	s.Unlock()
}

func (s *State) RemoveMember(guildID string, member string) {
	if _, err := s.Guild(guildID); err != nil {
		return
	}

	s.Lock()
	if _, ok := s.members[guildID]; ok {
		delete(s.members[guildID], member)
	}
	s.Unlock()
}

func (s *State) Member(guildID string, userID string) (*discord.GuildMember, error) {
	if _, err := s.Guild(guildID); err != nil {
		return nil, err
	}

	s.RLock()
	if _, ok := s.members[guildID]; ok {
		if member, ok := s.members[guildID][userID]; ok {
			s.RUnlock()
			return member, nil
		}
	}
	s.RUnlock()

	member, _ := s.session.Guild.GetMember(guildID, userID)

	if member != nil {
		s.AddMember(guildID, member)

		return member, nil
	}

	return nil, errors.New("member not found")
}

func (s *State) Guilds() map[string]*discord.Guild {
	s.Lock()
	defer s.Unlock()

	return s.guilds
}

func (s *State) Channels() map[string]*discord.Channel {
	s.Lock()
	defer s.Unlock()

	return s.channels
}

func (s *State) Members() map[string]map[string]*discord.GuildMember {
	s.Lock()
	defer s.Unlock()

	return s.members
}
