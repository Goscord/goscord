package gateway

import (
	"github.com/Goscord/goscord/goscord/gateway/event"
)

type ChannelUpdateHandler struct{}

func (_ *ChannelUpdateHandler) Handle(s *Session, data []byte) {
	ev, err := event.NewChannelUpdate(s.rest, data)

	if err != nil {
		return
	}

	s.state.AddChannel(ev.Data)

	s.Bus().Publish("channelUpdate", ev.Data)
}
