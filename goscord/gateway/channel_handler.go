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

	s.Publish(event.EventChannelCreate, ev.Data)
}

type ChannelUpdateHandler struct{}

func (_ *ChannelUpdateHandler) Handle(s *Session, data []byte) {
	ev, err := event.NewChannelUpdate(s.rest, data)

	if err != nil {
		return
	}

	s.State().AddChannel(ev.Data)

	s.Publish(event.EventChannelUpdate, ev.Data)
}

type ChannelDeleteHandler struct{}

func (_ *ChannelDeleteHandler) Handle(s *Session, data []byte) {
	ev, err := event.NewChannelDelete(s.rest, data)

	if err != nil {
		return
	}

	ev.Data, _ = s.State().Channel(ev.Data.Id)

	s.State().RemoveChannel(ev.Data)

	s.Publish(event.EventChannelDelete, ev.Data)
}

type ChannelPinsUpdateHandler struct{}

func (_ *ChannelPinsUpdateHandler) Handle(s *Session, data []byte) {
	ev, err := event.NewChannelPinsUpdate(s.rest, data)

	if err != nil {
		return
	}

	s.Publish(event.EventChannelPinsUpdate, ev.Data)
}

type ThreadCreateHandler struct{}

func (_ *ThreadCreateHandler) Handle(s *Session, data []byte) {
	ev, err := event.NewThreadCreate(s.rest, data)

	if err != nil {
		return
	}

	s.State().AddChannel(ev.Data)

	s.Publish(event.EventThreadCreate, ev.Data)
}

type ThreadUpdateHandler struct{}

func (_ *ThreadUpdateHandler) Handle(s *Session, data []byte) {
	ev, err := event.NewThreadUpdate(s.rest, data)

	if err != nil {
		return
	}

	s.State().AddChannel(ev.Data)

	s.Publish(event.EventThreadUpdate, ev.Data)
}

type ThreadDeleteHandler struct{}

func (_ *ThreadDeleteHandler) Handle(s *Session, data []byte) {
	ev, err := event.NewThreadDelete(s.rest, data)

	if err != nil {
		return
	}

	ev.Data, _ = s.State().Channel(ev.Data.Id)

	s.State().RemoveChannel(ev.Data)

	s.Publish(event.EventThreadDelete, ev.Data)
}

type ThreadListSyncHandler struct{}

func (_ *ThreadListSyncHandler) Handle(s *Session, data []byte) {
	ev, err := event.NewThreadListSync(s.rest, data)

	if err != nil {
		return
	}

	for _, thread := range ev.Data.Threads {
		s.State().AddChannel(thread)
	}

	s.Publish(event.EventThreadListSync, ev.Data)
}

type ThreadMemberUpdateHandler struct{}

func (_ *ThreadMemberUpdateHandler) Handle(s *Session, data []byte) {
	ev, err := event.NewThreadMemberUpdate(s.rest, data)

	if err != nil {
		return
	}

	// ToDo : Update thread member?

	s.Publish(event.EventThreadMemberUpdate, ev.Data)
}

type ThreadMembersUpdateHandler struct{}

func (_ *ThreadMembersUpdateHandler) Handle(s *Session, data []byte) {
	ev, err := event.NewThreadMembersUpdate(s.rest, data)

	if err != nil {
		return
	}

	// ToDo : Update thread members?

	s.Publish(event.EventThreadMembersUpdate, ev.Data)
}
