package event

import (
	"github.com/Goscord/goscord/goscord/discord"
	"github.com/Goscord/goscord/goscord/rest"
	"github.com/bytedance/sonic"
)

type AutoModerationRuleDelete struct {
	Data *discord.AutoModerationRule `json:"d"`
}

func NewAutoModerationRuleDelete(rest *rest.Client, data []byte) (*AutoModerationRuleDelete, error) {
	pk := new(AutoModerationRuleDelete)

	err := sonic.Unmarshal(data, pk)

	if err != nil {
		return nil, err
	}

	return pk, nil
}
