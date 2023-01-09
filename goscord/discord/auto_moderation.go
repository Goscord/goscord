package discord

type AutoModerationRuleTriggerType int

const (
	AutoModerationRuleTriggerTypeKeyword AutoModerationRuleTriggerType = iota + 1
	AutoModerationRuleTriggerTypeSpam
	AutoModerationRuleTriggerTypeKeywordPreset
	AutoModerationRuleTriggerTypeMentionSpam
)

type AutoModerationRuleKeywordPreset int

const (
	AutoModerationRuleKeywordPresetProfanity     AutoModerationRuleKeywordPreset = iota + 1 // Words that may be considered forms of swearing or cursing
	AutoModerationRuleKeywordPresetSexualContent                                            // Words that refer to sexually explicit behavior or activity
	AutoModerationRuleKeywordPresetSlurs                                                    // Personal insults or words that may be considered hate speech
)

type AutoModerationRuleEventType int

const (
	AutoModerationRuleEventTypeMessageSend AutoModerationRuleEventType = iota + 1 // when a member sends or edits a message in the guild
)

type AutoModerationRule struct {
	Id          string                        `json:"id"`
	GuildId     string                        `json:"guild_id"`
	Name        string                        `json:"name"`
	CreatorId   string                        `json:"creator_id"`
	EventType   AutoModerationRuleEventType   `json:"event_type"`
	TriggerType AutoModerationRuleTriggerType `json:"trigger_type"`
	// ToDo : TriggerMetadata
	Actions        []*AutoModerationAction `json:"actions"`
	Enabled        bool                    `json:"enabled"`
	ExemptRoles    []string                `json:"exempt_roles"`    // 	the role ids that should not be affected by the rule (Maximum of 20)
	ExemptChannels []string                `json:"exempt_channels"` //	the channel ids that should not be affected by the rule (Maximum of 50)
}

type AutoModerationActionType int

const (
	AutoModerationActionTypeBlockMessage AutoModerationActionType = iota + 1
	AutoModerationActionTypeSendAlertMessage
	AutoModerationActionTypeTimeout
)

type AutoModerationAction struct {
	Type     AutoModerationActionType      `json:"type"`
	Metadata *AutoModerationActionMetadata `json:"metadata,omitempty"`
}

type AutoModerationActionMetadata struct {
	ChannelId       string `json:"channel_id"`
	DurationSeconds int    `json:"duration_seconds"`
}

// AutoModerationActionExecutionEventFields is used by AUTO_MODERATION_ACTION_EXECUTION event
type AutoModerationActionExecutionEventFields struct {
	GuildId              string                        `json:"guild_id"`
	Action               *AutoModerationAction         `json:"action"`
	RuleId               string                        `json:"rule_id"`
	RuleTriggerType      AutoModerationRuleTriggerType `json:"rule_trigger_type"`
	UserId               string                        `json:"user_id"`
	ChannelId            string                        `json:"channel_id,omitempty"`
	MessageId            string                        `json:"message_id,omitempty"`
	AlertSystemMessageId string                        `json:"alert_system_message_id,omitempty"`
	Content              string                        `json:"content"`
	MatchedKeyword       string                        `json:"matched_keyword"`
	MatchedContent       string                        `json:"matched_content"`
}
