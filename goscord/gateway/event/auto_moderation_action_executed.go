package event

import (
	"github.com/Goscord/goscord/goscord/discord"
	"github.com/Goscord/goscord/goscord/rest"
	"github.com/goccy/go-json"
)

type AutoModerationActionExecution struct {
	Data struct {
		GuildId              string                                `json:"guild_id"`
		Action               *discord.AutoModerationAction         `json:"action"`
		RuleId               string                                `json:"rule_id"`
		RuleTriggerType      discord.AutoModerationRuleTriggerType `json:"rule_trigger_type"`
		UserId               string                                `json:"user_id"`
		ChannelId            string                                `json:"channel_id,omitempty"`
		MessageId            string                                `json:"message_id,omitempty"`
		AlertSystemMessageId string                                `json:"alert_system_message_id,omitempty"`
		Content              string                                `json:"content"`
		MatchedKeyword       string                                `json:"matched_keyword"`
		MatchedContent       string                                `json:"matched_content"`
	} `json:"d"`
}

func NewAutoModerationActionExecution(rest *rest.Client, data []byte) (*AutoModerationActionExecution, error) {
	pk := new(AutoModerationActionExecution)

	err := json.Unmarshal(data, pk)

	if err != nil {
		return nil, err
	}

	return pk, nil
}
