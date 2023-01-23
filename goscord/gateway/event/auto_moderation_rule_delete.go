package event

import (
	"github.com/Goscord/goscord/goscord/discord"
	"github.com/Goscord/goscord/goscord/rest"
	"github.com/goccy/go-json"
)

type AutoModerationRuleDelete struct {
	Data *discord.AutoModerationRule `json:"d"`
}

func NewAutoModerationRuleDelete(rest *rest.Client, data []byte) (*AutoModerationRuleDelete, error) {
	pk := new(AutoModerationRuleDelete)

	err := json.Unmarshal(data, pk)

	if err != nil {
		return nil, err
	}

	return pk, nil
}
