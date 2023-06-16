package gateway

import (
	"log"

	"github.com/Goscord/goscord/goscord/gateway/event"
)

const (
	EventAutoModerationRuleCreate = "auto_moderation_rule_create"
	EventAutoModerationRuleDelete = "auto_moderation_rule_delete"
	EventAutoModerationRuleUpdate = "auto_moderation_rule_update"
	EventAutoModerationActionExecution = "auto_moderation_action_execution"
)

type EventHandler interface {
	Handle(s *Session, data []byte)
}

type AutoModerationRuleCreateHandler struct{}
func (_ *AutoModerationRuleCreateHandler) Handle(s *Session, data []byte) {
	ev, err := event.NewAutoModerationRuleCreate(s.rest, data)
	if err != nil {
		log.Println("Error when creating the event AutoModerationRuleCreate:", err)
		return
	}

	s.Publish(EventAutoModerationRuleCreate, ev.Data)
}

type AutoModerationRuleDeleteHandler struct{}
func (_ *AutoModerationRuleDeleteHandler) Handle(s *Session, data []byte) {
	ev, err := event.NewAutoModerationRuleDelete(s.rest, data)
	if err != nil {
		log.Println("Error when creating the event AutoModerationRuleDelete:", err)
		return
	}

	s.Publish(EventAutoModerationRuleDelete, ev.Data)
}

type AutoModerationRuleUpdateHandler struct{}
func (_ *AutoModerationRuleUpdateHandler) Handle(s *Session, data []byte) {
	ev, err := event.NewAutoModerationRuleUpdate(s.rest, data)
	if err != nil {
		log.Println("Error when creating the event AutoModerationRuleUpdate:", err)
		return
	}

	s.Publish(EventAutoModerationRuleUpdate, ev.Data)
}

type AutoModerationActionExecutionHandler struct{}
func (_ *AutoModerationActionExecutionHandler) Handle(s *Session, data []byte) {
	ev, err := event.NewAutoModerationActionExecution(s.rest, data)
	if err != nil {
		log.Println("Error when creating the event AutoModerationActionExecution:", err)
		return
	}

	s.Publish(EventAutoModerationActionExecution, ev.Data)
}
