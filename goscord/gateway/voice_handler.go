package gateway

import (
	"github.com/Goscord/goscord/goscord/gateway/event"
)

type VoiceServerUpdateHandler struct{}

func (_ *VoiceServerUpdateHandler) Handle(s *Session, data []byte) {
	ev, err := event.NewVoiceServerUpdate(s.rest, data)

	if err != nil {
		return
	}

	s.RLock()
	voice, ok := s.VoiceConnections[ev.Data.GuildId]
	s.RUnlock()

	if !ok {
		return
	}

	voice.Close()

	voice.Lock()
	voice.GuildId = ev.Data.GuildId
	voice.token = ev.Data.Token
	voice.endpoint = ev.Data.Endpoint
	voice.Unlock()

	s.Publish(event.EventVoiceServerUpdate, ev.Data)

	_ = voice.login()
}

type VoiceStateUpdateHandler struct{}

func (_ *VoiceStateUpdateHandler) Handle(s *Session, data []byte) {
	ev, err := event.NewVoiceStateUpdate(s.rest, data)

	if err != nil {
		return
	}

	err = s.State().UpdateVoiceState(ev)
	if err != nil {
		return
	}

	s.RLock()
	voice, ok := s.VoiceConnections[ev.Data.GuildId]
	s.RUnlock()

	if !ok {
		return
	}

	if s.Me().Id == ev.Data.UserId {
		voice.Lock()
		voice.UserId = ev.Data.UserId
		voice.sessionId = ev.Data.SessionId
		voice.ChannelId = ev.Data.ChannelId
		voice.Unlock()
	}

	s.Publish(event.EventVoiceStateUpdate, ev.Data)
}
