package gateway

import "github.com/Goscord/goscord/goscord/gateway/event"

type ChannelCreateHandler struct{}

func (_ *ChannelCreateHandler) Handle(s *Session, data []byte) {
	ev, err := event.NewChannelCreate(s.rest, data)

	if err != nil {
		return
	}

	s.State().AddChannel(ev.Data)

	s.Bus().Publish("channelCreate", ev.Data)
}

type ChannelUpdateHandler struct{}

func (_ *ChannelUpdateHandler) Handle(s *Session, data []byte) {
	ev, err := event.NewChannelUpdate(s.rest, data)

	if err != nil {
		return
	}

	s.state.AddChannel(ev.Data)

	s.Bus().Publish("channelUpdate", ev.Data)
}

type ChannelDeleteHandler struct{}

func (_ *ChannelDeleteHandler) Handle(s *Session, data []byte) {
	ev, err := event.NewChannelDelete(s.rest, data)

	if err != nil {
		return
	}

	s.state.RemoveChannel(ev.Data)

	s.Bus().Publish("channelDelete", ev.Data)
}

type ChannelPinsUpdateHandler struct{}

func (_ *ChannelPinsUpdateHandler) Handle(s *Session, data []byte) {
	ev, err := event.NewChannelPinsUpdate(s.rest, data)

	if err != nil {
		return
	}

	s.Bus().Publish("channelPinsUpdate", ev.Data)
}
