package event

import (
	"github.com/Goscord/goscord/goscord/discord"
	"github.com/Goscord/goscord/goscord/rest"
	"github.com/goccy/go-json"
)

type AutoModerationActionExecution struct {
	Data *discord.AutoModerationActionExecutionEventFields `json:"d"`
}

func NewAutoModerationActionExecution(rest *rest.Client, data []byte) (*AutoModerationActionExecution, error) {
	pk := new(AutoModerationActionExecution)

	err := json.Unmarshal(data, pk)

	if err != nil {
		return nil, err
	}

	return pk, nil
}
