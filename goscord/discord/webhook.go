package discord

type WebhookType int

const (
	WebhookTypeIncoming WebhookType = iota + 1
	WebhookTypeChannelFollower
	WebhookTypeApplication
)

type Webhook struct {
	Id            string      `json:"id"`
	Type          WebhookType `json:"type"`
	GuildId       string      `json:"guild_id,omitempty"`
	ChannelId     string      `json:"channel_id"`
	User          *User       `json:"user,omitempty"`
	Name          string      `json:"name"`
	Avatar        string      `json:"avatar"`
	Token         string      `json:"token,omitempty"`
	ApplicationId string      `json:"application_id"`
	SourceGuild   *Guild      `json:"source_guild,omitempty"`
	SourceChannel *Channel    `json:"source_channel,omitempty"`
	Url           string      `json:"url,omitempty"`
}
