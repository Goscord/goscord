package discord

type PrivacyLevel int

const (
	PrivacyLevelPublic PrivacyLevel = iota + 1
	PrivacyLevelGuildOnly
)

type StageInstance struct {
	Id                    string       `json:"id"`
	GuildId               string       `json:"guild_id"`
	ChannelId             string       `json:"channel_id"`
	Topic                 string       `json:"topic"`
	PrivacyLevel          PrivacyLevel `json:"privacy_level"`
	DiscoverableDisabled  bool         `json:"discoverable_disabled"`
	GuildScheduledEventId string       `json:"guild_scheduled_event_id"`
}
