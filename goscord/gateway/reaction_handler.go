package gateway

import "github.com/Goscord/goscord/goscord/gateway/event"

type MessageReactionAddHandler struct{}

func (_ *MessageReactionAddHandler) Handle(s *Session, data []byte) {
	ev, err := event.NewMessageReactionAdd(s.rest, data)

	if err != nil {
		return
	}

	s.Publish(event.EventMessageReactionAdd, ev.Data)
}

type MessageReactionRemoveHandler struct{}

func (_ *MessageReactionRemoveHandler) Handle(s *Session, data []byte) {
	ev, err := event.NewMessageReactionRemove(s.rest, data)

	if err != nil {
		return
	}

	s.Publish(event.EventMessageReactionRemove, ev.Data)
}

type MessageReactionRemoveAllHandler struct{}

func (_ *MessageReactionRemoveAllHandler) Handle(s *Session, data []byte) {
	ev, err := event.NewMessageReactionRemoveAll(s.rest, data)

	if err != nil {
		return
	}

	s.Publish(event.EventMessageReactionRemoveAll, ev.Data)
}

type MessageReactionRemoveEmojiHandler struct{}

func (_ *MessageReactionRemoveEmojiHandler) Handle(s *Session, data []byte) {
	ev, err := event.NewMessageReactionRemoveEmoji(s.rest, data)

	if err != nil {
		return
	}

	s.Publish(event.EventMessageReactionRemoveEmoji, ev.Data)
}
