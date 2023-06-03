package gateway

import "github.com/Goscord/goscord/goscord/gateway/event"

type MessageCreateHandler struct{}

// Handle handles the message create event
func (_ *MessageCreateHandler) Handle(s *Session, data []byte) {
	ev, err := event.NewMessageCreate(s.rest, data)

	if err != nil {
		return
	}

	s.Publish(event.EventMessageCreate, ev.Data)
}

type MessageUpdateHandler struct{}

// Handle handles the message update event
func (_ *MessageUpdateHandler) Handle(s *Session, data []byte) {
	ev, err := event.NewMessageUpdate(s.rest, data)

	if err != nil {
		return
	}

	s.Publish(event.EventMessageUpdate, ev.Data)
}

type MessageDeleteHandler struct{}

// Handle handles the message delete event
func (_ *MessageDeleteHandler) Handle(s *Session, data []byte) {
	ev, err := event.NewMessageDelete(s.rest, data)

	if err != nil {
		return
	}

	s.Publish(event.EventMessageDelete, ev.Data)
}
