package event

import (
	"github.com/Goscord/goscord/goscord/discord"
	"github.com/Goscord/goscord/goscord/rest"
	"github.com/bytedance/sonic"
)

type AutoModerationActionExecution struct {
	Data *discord.AutoModerationActionExecutionEventFields `json:"d"`
}

func NewAutoModerationActionExecution(rest *rest.Client, data []byte) (*AutoModerationActionExecution, error) {
	pk := new(AutoModerationActionExecution)

	err := sonic.Unmarshal(data, pk)

	if err != nil {
		return nil, err
	}

	return pk, nil
}
