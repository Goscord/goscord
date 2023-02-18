package gateway

import (
	"github.com/Goscord/goscord/goscord/gateway/event"
)

type AutoModerationRuleCreateHandler struct{}

func (_ *AutoModerationRuleCreateHandler) Handle(s *Session, data []byte) {
	ev, err := event.NewAutoModerationRuleCreate(s.rest, data)

	if err != nil {
		return
	}

	s.Publish(event.EventAutoModerationRuleCreate, ev.Data)
}

type AutoModerationRuleDeleteHandler struct{}

func (_ *AutoModerationRuleDeleteHandler) Handle(s *Session, data []byte) {
	ev, err := event.NewAutoModerationRuleDelete(s.rest, data)

	if err != nil {
		return
	}

	s.Publish(event.EventAutoModerationRuleDelete, ev.Data)
}

type AutoModerationRuleUpdateHandler struct{}

func (_ *AutoModerationRuleUpdateHandler) Handle(s *Session, data []byte) {
	ev, err := event.NewAutoModerationRuleUpdate(s.rest, data)

	if err != nil {
		return
	}

	s.Publish(event.EventAutoModerationRuleUpdate, ev.Data)
}

type AutoModerationActionExecutionHandler struct{}

func (_ *AutoModerationActionExecutionHandler) Handle(s *Session, data []byte) {
	ev, err := event.NewAutoModerationActionExecution(s.rest, data)

	if err != nil {
		return
	}

	s.Publish(event.EventAutoModerationActionExecution, ev.Data)
}
