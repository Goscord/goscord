package discord

import (
	"time"
)

type Channel struct {
	Id       string `json:"id"`
	Type     int    `json:"type"`
	GuildId  string `json:"guild_id"`
	Position int    `json:"position"`
	//PermissionOverwrites []PermissionOverwrite `json:"permission_overwrites"`
	Name             string     `json:"name"`
	Topic            string     `json:"topic"`
	Nsfw             bool       `json:"nsfw"`
	LastMessageId    string     `json:"last_message_id"`
	Bitrate          int        `json:"bitrate"`
	UserLimit        int        `json:"user_limit"`
	RateLimitPerUser int        `json:"rate_limit_per_user"`
	Recipients       []User     `json:"recipients"`
	Icon             string     `json:"icon"`
	OwnerId          string     `json:"owner_id"`
	ApplicationId    string     `json:"application_id"`
	ParentId         string     `json:"parent_id"`
	LastPinTimestamp *time.Time `json:"last_pin_timestamp"`
}
