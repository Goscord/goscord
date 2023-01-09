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
