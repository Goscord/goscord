package gateway

import (
	"github.com/Goscord/goscord/goscord/gateway/event"
)

type ChannelCreateHandler struct{}

func (_ *ChannelCreateHandler) Handle(s *Session, data []byte) {
	ev, err := event.NewChannelCreate(s.rest, data)

	if err != nil {
		return
	}

	s.State().AddChannel(ev.Data)

	s.Bus().Publish("channelCreate", ev.Data)
}
