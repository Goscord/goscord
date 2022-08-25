package discord

import "time"

type Invite struct {
	Code                     string               `json:"code"`
	Guild                    *Guild               `json:"guild,omitempty"`
	Channel                  *Channel             `json:"channel"`
	Inviter                  *User                `json:"inviter,omitempty"`
	TargetType               int                  `json:"target_type,omitempty"`
	TargetUser               *User                `json:"target_user,omitempty"`
	TargetApplication        *Application         `json:"target_application,omitempty"`
	ApproximatePresenceCount int                  `json:"approximate_presence_count,omitempty"`
	ApproximateMemberCount   int                  `json:"approximate_member_count,omitempty" `
	ExpiresAt                *time.Time           `json:"expires_at,omitempty"`
	StageInstance            *StageInstance       `json:"stage_instance,omitempty"`
	GuildScheduledEvent      *GuildScheduledEvent `json:"guild_scheduled_event,omitempty"`
}
