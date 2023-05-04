package gateway

import (
	"errors"
	"github.com/Goscord/goscord/goscord/discord"
	"github.com/Goscord/goscord/goscord/gateway/event"
	"sync"
)

type State struct {
	sync.RWMutex

	guilds   map[string]*discord.Guild
	channels map[string]*discord.Channel
	members  map[string]map[string]*discord.GuildMember
}

// NewState creates a new state
func NewState() *State {
	return &State{
		guilds:   make(map[string]*discord.Guild),
		channels: make(map[string]*discord.Channel),
		members:  make(map[string]map[string]*discord.GuildMember),
	}
}

// GUILDS

// AddGuild adds a guild to the state
func (s *State) AddGuild(guild *discord.Guild) {
	s.Lock()
	defer s.Unlock()

	// ToDo : Maybe set the guild id for channels and voice states

	for _, member := range guild.Members {
		member.GuildId = guild.Id
	}

	for _, c := range guild.Channels {
		s.channels[c.Id] = c
	}

	for _, t := range guild.Threads {
		s.channels[t.Id] = t
	}

	if guild.Members != nil {
		s.registerMembers(guild)
	} else if _, ok := s.members[guild.Id]; !ok {
		s.members[guild.Id] = make(map[string]*discord.GuildMember)
	}

	if g, ok := s.guilds[guild.Id]; ok {
		if guild.MemberCount == 0 {
			guild.MemberCount = g.MemberCount
		}

		if guild.Roles == nil {
			guild.Roles = g.Roles
		}

		if guild.Emojis == nil {
			guild.Emojis = g.Emojis
		}

		if guild.Members == nil {
			guild.Members = g.Members
		}

		if guild.Presences == nil {
			guild.Presences = g.Presences
		}

		if guild.Channels == nil {
			guild.Channels = g.Channels
		}

		if guild.Threads == nil {
			guild.Threads = g.Threads
		}

		if guild.VoiceStates == nil {
			guild.VoiceStates = g.VoiceStates
		}

		*g = *guild

		return
	}

	s.guilds[guild.Id] = guild
}

func (s *State) registerMembers(guild *discord.Guild) {
	members := make(map[string]*discord.GuildMember)

	for _, m := range guild.Members {
		members[m.User.Id] = m
	}

	s.members[guild.Id] = members
}

// RemoveGuild removes a guild from the state
func (s *State) RemoveGuild(guild *discord.Guild) error {
	_, err := s.Guild(guild.Id)
	if err != nil {
		return err
	}

	s.Lock()
	defer s.Unlock()

	delete(s.guilds, guild.Id)

	return nil
}

// Guild returns a guild from the state
func (s *State) Guild(guildID string) (*discord.Guild, error) {
	s.RLock()
	defer s.RUnlock()

	if g, ok := s.guilds[guildID]; ok {
		return g, nil
	}

	return nil, errors.New("guild not found")
}

// CHANNELS

// AddChannel adds a channel to the state
func (s *State) AddChannel(channel *discord.Channel) {
	s.Lock()
	defer s.Unlock()

	if c, ok := s.channels[channel.Id]; ok {
		if channel.ThreadMetadata == nil {
			channel.ThreadMetadata = c.ThreadMetadata
		}

		*c = *channel
	}

	if channel.Type == discord.ChannelTypeDM || channel.Type == discord.ChannelTypeGroupDM {
		s.channels[channel.Id] = channel
		return
	}

	guild, ok := s.guilds[channel.GuildId]
	if !ok {
		return
	}

	if channel.Type == discord.ChannelTypePublicThread || channel.Type == discord.ChannelTypePrivateThread {
		guild.Threads = append(guild.Threads, channel)
	} else {
		guild.Channels = append(guild.Channels, channel)
	}

	s.channels[channel.Id] = channel
}

// RemoveChannel removes a channel from the state
func (s *State) RemoveChannel(channel *discord.Channel) {
	_, err := s.Channel(channel.Id)
	if err != nil {
		return
	}

	if channel.Type == discord.ChannelTypeDM || channel.Type == discord.ChannelTypeGroupDM {
		s.Lock()
		defer s.Unlock()

		delete(s.channels, channel.Id)

		return
	}

	guild, err := s.Guild(channel.GuildId)
	if err != nil {
		return
	}

	s.Lock()
	defer s.Unlock()

	if channel.Type == discord.ChannelTypePublicThread || channel.Type == discord.ChannelTypePrivateThread {
		for i, t := range guild.Threads {
			if t.Id == channel.Id {
				guild.Threads = append(guild.Threads[:i], guild.Threads[i+1:]...)
				break
			}
		}
	} else {
		for i, c := range guild.Channels {
			if c.Id == channel.Id {
				guild.Channels = append(guild.Channels[:i], guild.Channels[i+1:]...)
				break
			}
		}
	}

	delete(s.channels, channel.Id)
}

// Channel returns a channel from the state
func (s *State) Channel(id string) (*discord.Channel, error) {
	s.RLock()
	defer s.RUnlock()

	if c, ok := s.channels[id]; ok {
		return c, nil
	}

	return nil, errors.New("channel not found")
}

// MEMBERS

// AddMember adds a member to the state
func (s *State) AddMember(guildID string, member *discord.GuildMember) {
	s.Lock()
	defer s.Unlock()

	guild, ok := s.guilds[member.GuildId]
	if !ok {
		return
	}

	members, ok := s.members[member.GuildId]
	if !ok {
		return
	}

	m, ok := members[member.User.Id]
	if !ok {
		members[member.User.Id] = member
		guild.Members = append(guild.Members, member)
	} else {
		if member.JoinedAt.IsZero() {
			member.JoinedAt = m.JoinedAt
		}

		*m = *member
	}
}

// RemoveMember removes a member from the state
func (s *State) RemoveMember(guildId string, memberId string) {
	guild, err := s.Guild(guildId)
	if err != nil {
		return
	}

	s.Lock()
	defer s.Unlock()

	members, ok := s.members[guildId]
	if !ok {
		return
	}

	_, ok = members[memberId]
	if !ok {
		return
	}

	delete(members, memberId)

	for i, m := range guild.Members {
		if m.User.Id == memberId {
			guild.Members = append(guild.Members[:i], guild.Members[i+1:]...)

			return
		}
	}
}

// Member returns a member from the state
func (s *State) Member(guildID string, userID string) (*discord.GuildMember, error) {
	s.RLock()
	defer s.RUnlock()

	members, ok := s.members[guildID]
	if !ok {
		return nil, errors.New("guild members not found")
	}

	m, ok := members[userID]
	if ok {
		return m, nil
	}

	// ToDo : Get member from the API

	return nil, errors.New("guild member not found")
}

// EMOJIS

// AddEmoji adds an emoji to the state
func (s *State) AddEmoji(guildId string, emoji *discord.Emoji) error {
	guild, err := s.Guild(guildId)
	if err != nil {
		return err
	}

	s.Lock()
	defer s.Unlock()

	for i, e := range guild.Emojis {
		if e.Id == emoji.Id {
			guild.Emojis[i] = emoji

			return nil
		}
	}

	guild.Emojis = append(guild.Emojis, emoji)

	return nil
}

// Emoji returns an emoji from the state
func (s *State) Emoji(guildId, emojiId string) (*discord.Emoji, error) {
	guild, err := s.Guild(guildId)
	if err != nil {
		return nil, err
	}

	s.RLock()
	defer s.RUnlock()

	for _, e := range guild.Emojis {
		if e.Id == emojiId {
			return e, nil
		}
	}

	return nil, errors.New("emoji not found")
}

// AddEmojis adds emojis to the state
func (s *State) AddEmojis(guildId string, emojis []*discord.Emoji) error {
	for _, e := range emojis {
		if err := s.AddEmoji(guildId, e); err != nil {
			return err
		}
	}

	return nil
}

// AddRole adds a role to the state
func (s *State) AddRole(guildId string, role *discord.Role) error {
	guild, err := s.Guild(guildId)
	if err != nil {
		return err
	}

	s.Lock()
	defer s.Unlock()

	for i, r := range guild.Roles {
		if r.Id == role.Id {
			guild.Roles[i] = role

			return nil
		}
	}

	guild.Roles = append(guild.Roles, role)

	return nil
}

// RemoveRole removes a role from the state
func (s *State) RemoveRole(guildId, roleId string) error {
	guild, err := s.Guild(guildId)
	if err != nil {
		return err
	}

	s.Lock()
	defer s.Unlock()

	for i, r := range guild.Roles {
		if r.Id == roleId {
			guild.Roles = append(guild.Roles[:i], guild.Roles[i+1:]...)
			return nil
		}
	}

	return errors.New("role not found")
}

// Role returns a role from the state
func (s *State) Role(guildId, roleId string) (*discord.Role, error) {
	guild, err := s.Guild(guildId)
	if err != nil {
		return nil, err
	}

	s.RLock()
	defer s.RUnlock()

	for _, r := range guild.Roles {
		if r.Id == roleId {
			return r, nil
		}
	}

	return nil, errors.New("role not found")
}

// UpdateVoiceState updates a voice state in the state
func (s *State) UpdateVoiceState(ev *event.VoiceStateUpdate) error {
	guild, err := s.Guild(ev.Data.GuildId)
	if err != nil {
		return err
	}

	s.Lock()
	defer s.Unlock()

	if ev.Data.ChannelId == "" {
		for i, state := range guild.VoiceStates {
			if state.UserId == ev.Data.UserId {
				guild.VoiceStates = append(guild.VoiceStates[:i], guild.VoiceStates[i+1:]...)

				return nil
			}
		}
	} else {
		for i, state := range guild.VoiceStates {
			if state.UserId == ev.Data.UserId {
				guild.VoiceStates[i] = ev.Data

				return nil
			}
		}

		guild.VoiceStates = append(guild.VoiceStates, ev.Data)
	}

	return nil
}

// VoiceState returns a voice state from the state
func (s *State) VoiceState(guildId, userId string) (*discord.VoiceState, error) {
	guild, err := s.Guild(guildId)
	if err != nil {
		return nil, err
	}

	for _, state := range guild.VoiceStates {
		if state.UserId == userId {
			return state, nil
		}
	}

	return nil, errors.New("voice state not found")
}

// Guilds returns all guilds from the state
func (s *State) Guilds() map[string]*discord.Guild {
	s.RLock()
	defer s.RUnlock()

	return s.guilds
}

// Channels returns all channels from the state
func (s *State) Channels() map[string]*discord.Channel {
	s.RLock()
	defer s.RUnlock()

	return s.channels
}

// Members returns all members from the state
func (s *State) Members() map[string]map[string]*discord.GuildMember {
	s.RLock()
	defer s.RUnlock()

	return s.members
}
