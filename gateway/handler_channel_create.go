package gateway

import (
	"github.com/Goscord/goscord/gateway/event"
)

type ChannelCreateHandler struct{}

func (h *ChannelCreateHandler) Handle(s *Session, data []byte) {
	ev, err := event.NewChannelCreate(s.rest, data)

	if err != nil {
		return
	}

	if _, ok := s.state.Channels[ev.Data.Id]; !ok {
		s.state.AddChannel(ev.Data)
		s.bus.Publish("channelCreate", ev.Data)
	} else {
		s.state.UpdateChannel(ev.Data)
	}
}
