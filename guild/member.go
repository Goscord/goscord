package guild

import (
	"github.com/Seyz123/yalis/user"
	"time"
)

type Member struct {
	User         *user.User `json:"user"`
	Nick         string     `json:"nick,omitempty"`
	Roles        []int      `json:"roles"`
	JoinedAt     time.Time  `json:"joined_at"`
	PremiumSince time.Time  `json:"premium_since,omitempty"`
	Deaf         bool       `json:"deaf"`
	Mute         bool       `json:"mute"`
}
