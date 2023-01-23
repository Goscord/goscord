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

	s.Bus().Publish("autoModerationRuleCreate", ev.Data)
}

type AutoModerationRuleDeleteHandler struct{}

func (_ *AutoModerationRuleDeleteHandler) Handle(s *Session, data []byte) {
	ev, err := event.NewAutoModerationRuleDelete(s.rest, data)

	if err != nil {
		return
	}

	s.Bus().Publish("autoModerationRuleDelete", ev.Data)
}

type AutoModerationRuleUpdateHandler struct{}

func (_ *AutoModerationRuleUpdateHandler) Handle(s *Session, data []byte) {
	ev, err := event.NewAutoModerationRuleUpdate(s.rest, data)

	if err != nil {
		return
	}

	s.Bus().Publish("autoModerationRuleUpdate", ev.Data)
}

type AutoModerationActionExecutionHandler struct{}

func (_ *AutoModerationActionExecutionHandler) Handle(s *Session, data []byte) {
	ev, err := event.NewAutoModerationActionExecution(s.rest, data)

	if err != nil {
		return
	}

	s.Bus().Publish("autoModerationActionExecution", ev.Data)
}
