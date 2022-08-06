package discord

import (
	"time"
)

type ChannelType int

const (
	ChannelTypeText ChannelType = iota
	ChannelTypeDM
	ChannelTypeVoice
	ChannelTypeGroupDM
	ChannelTypeCategory
	ChannelTypeNews
	ChannelTypeNewsThread
	ChannelTypePublicThread
	ChannelTypePrivateThread
	ChannelTypeStageVoice
	ChannelTypeDirectory
	ChannelTypeForum // Still in development
)

type VideoQuality int

const (
	_                VideoQuality = iota
	VideoQualityAuto              // Discord chooses the quality for optimal performance
	VideoQualityFull              // 720p, 1080p, or higher quality
)

type ChannelFlags int

const (
	ChannelFlagsPinned ChannelFlags = 1 << 1
)

type FollowedChannel struct {
	ChannelId string `json:"channel_id"`
	WebhookId string `json:"webhook_id"`
}

type Overwrite struct {
	Id    string `json:"id"`    // role or user id
	Type  int    `json:"type"`  // 0 for role, 1 for user
	Allow int    `json:"allow"` // permission bit set
	Deny  int    `json:"deny"`  // permission bit set
}

type Channel struct {
	Id       string      `json:"id"`
	Type     ChannelType `json:"type"`
	GuildId  string      `json:"guild_id,omitempty"`
	Position int         `json:"position,omitempty"`
	//PermissionOverwrites []PermissionOverwrite `json:"permission_overwrites,omitempty"`
	Name                       string         `json:"name,omitempty"`
	Topic                      string         `json:"topic,omitempty"`
	Nsfw                       bool           `json:"nsfw,omitempty"`
	LastMessageId              string         `json:"last_message_id,omitempty"`
	Bitrate                    int            `json:"bitrate,omitempty"`
	UserLimit                  int            `json:"user_limit,omitempty"`
	RateLimitPerUser           int            `json:"rate_limit_per_user,omitempty"`
	Recipients                 []User         `json:"recipients,omitempty"`
	Icon                       string         `json:"icon,omitempty"`
	OwnerId                    string         `json:"owner_id,omitempty"`
	ApplicationId              string         `json:"application_id,omitempty"`
	ParentId                   string         `json:"parent_id,omitempty"`
	LastPinTimestamp           *time.Time     `json:"last_pin_timestamp,omitempty"`
	RtcRegion                  string         `json:"rtc_region,omitempty"` // 	voice region id for the voice channel, automatic when set to null
	VideoQualityMode           VideoQuality   `json:"video_quality_mode,omitempty"`
	MessageCount               int            `json:"message_count,omitempty"`
	MemberCount                int            `json:"member_count,omitempty"`
	ThreadMetadata             ThreadMetadata `json:"thread_metadata,omitempty"`
	Member                     *ThreadMember  `json:"member,omitempty"`
	DefaultAutoArchiveDuration int            `json:"default_auto_archive_duration,omitempty"`
	Permissions                string         `json:"permissions,omitempty"`
	Flags                      ChannelFlags   `json:"flags,omitempty"`
	TotalMessageSent           int            `json:"total_message_sent,omitempty"`
}
