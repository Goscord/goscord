package gateway

import (
	"github.com/Goscord/goscord/gateway/event"
)

type ChannelDeleteHandler struct{}

func (_ *ChannelDeleteHandler) Handle(s *Session, data []byte) {
	ev, err := event.NewChannelDelete(s.rest, data)

	if err != nil {
		return
	}

	s.state.RemoveChannel(ev.Data)

	s.bus.Publish("channelDelete", ev.Data)
}
