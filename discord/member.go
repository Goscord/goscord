package discord

import (
	"time"
)

type Member struct {
	User         *User     `json:"user"`
	Nick         string    `json:"nick,omitempty"`
	Roles        []string  `json:"roles"`
	JoinedAt     time.Time `json:"joined_at"`
	PremiumSince time.Time `json:"premium_since,omitempty"`
	Deaf         bool      `json:"deaf"`
	Mute         bool      `json:"mute"`
}
