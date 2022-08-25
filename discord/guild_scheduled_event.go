package discord

import "time"

type GuildScheduledEventPrivacyLevel int

const (
	GuildScheduledEventPrivacyLevelGuildOnly GuildScheduledEventPrivacyLevel = 2
)

type GuildScheduledEventEntityType int

const (
	GuildScheduledEventEntityTypeStageInstance GuildScheduledEventEntityType = iota + 1
	GuildScheduledEventEntityTypeVoice
	GuildScheduledEventEntityTypeExternal
)

type GuildScheduledEventStatus int

const (
	GuildScheduledEventStatusScheduled GuildScheduledEventStatus = iota + 1
	GuildScheduledEventStatusActive
	GuildScheduledEventStatusCompleted
	GuildScheduledEventStatusCanceled
)

type GuildScheduledEventEntityMetadata struct {
	Location string `json:"location,omitempty"` // 	location of the event (1-100 characters)
}

type GuildScheduledEventUserObject struct {
	GuildScheduledEventId string       `json:"guild_scheduled_event_id"`
	User                  *User        `json:"user"`
	Member                *GuildMember `json:"member,omitempty"`
}

type GuildScheduledEvent struct {
	Id                 string                             `json:"id"`
	GuildId            string                             `json:"guild_id"`
	ChannelId          string                             `json:"channel_id"`
	CreatorId          string                             `json:"creator_id,omitempty"`
	Name               string                             `json:"name"`
	Description        string                             `json:"description,omitempty"`
	ScheduledStartTime *time.Time                         `json:"scheduled_start_time"`
	ScheduledEndTime   *time.Time                         `json:"scheduled_end_time"`
	PrivacyLevel       GuildScheduledEventPrivacyLevel    `json:"privacy_level"`
	Status             GuildScheduledEventStatus          `json:"status"`
	EntityType         GuildScheduledEventEntityType      `json:"entity_type"`
	EntityId           string                             `json:"entity_id"`
	EntityMetadata     *GuildScheduledEventEntityMetadata `json:"entity_metadata"`
	Creator            *User                              `json:"creator,omitempty"`
	UserCount          int                                `json:"user_count,omitempty"`
	Image              string                             `json:"image,omitempty"`
}
