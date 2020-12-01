package guild

import (
	"github.com/Seyz123/yalis/rest"
	"github.com/Seyz123/yalis/user"
	"time"
)

type Member struct {
	Rest *rest.RestClient `json:"-"`
	User         *user.User `json:"user"`
	Nick         string     `json:"nick,omitempty"`
	Roles        []string   `json:"roles"`
	JoinedAt     time.Time  `json:"joined_at"`
	PremiumSince time.Time  `json:"premium_since,omitempty"`
	Deaf         bool       `json:"deaf"`
	Mute         bool       `json:"mute"`
}
