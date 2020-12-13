package gateway

import "github.com/Goscord/goscord/gateway/event"

type ChannelUpdateHandler struct{}

func (h *ChannelUpdateHandler) Handle(s *Session, data []byte) {
	ev, err := event.NewChannelUpdate(s.rest, data)

	if err != nil {
		return
	}

	s.state.UpdateChannel(ev.Data)

	s.bus.Publish("channelUpdate", ev.Data)
}
