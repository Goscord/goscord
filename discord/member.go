package discord

import (
	"encoding/json"
	"github.com/Goscord/goscord/rest"
	"github.com/Goscord/goscord/user"
	"time"
)

type Member struct {
	Rest         *rest.Client `json:"-"`
	User         *user.User   `json:"user"`
	Nick         string       `json:"nick,omitempty"`
	Roles        []string     `json:"roles"`
	JoinedAt     time.Time    `json:"joined_at"`
	PremiumSince time.Time    `json:"premium_since,omitempty"`
	Deaf         bool         `json:"deaf"`
	Mute         bool         `json:"mute"`
}

func NewMember(rest *rest.Client, data []byte) (*Member, error) {
	u := new(Member)

	err := json.Unmarshal(data, u)

	if err != nil {
		return nil, err
	}

	u.Rest = rest

	return u, nil
}
