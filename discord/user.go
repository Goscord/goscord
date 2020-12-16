package discord

import (
	"fmt"
	"strings"
)

type User struct {
	Id            string `json:"id"`
	Username      string `json:"username"`
	Discriminator string `json:"discriminator"`
	Avatar        string `json:"avatar"`
	Bot           bool   `json:"bot"`
	System        bool   `json:"system"`
	MfaEnabled    bool   `json:"mfa_enabled"`
	Locale        string `json:"locale"`
	Verified      bool   `json:"verified"`
	Email         string `json:"email"`
	Flags         int    `json:"flags"`
	PremiumType   int    `json:"premium_type"`
	PublicFlags   int    `json:"public_flags"`
}

func (u *User) Tag() string {
	return fmt.Sprintf("%s#%s", u.Username, u.Discriminator)
}

func (u *User) AvatarURL() string {
	if strings.HasPrefix(u.Avatar, "a_") {
		return fmt.Sprintf("https://cdn.discordapp.com/avatars/%s/%s.gif", u.Id, u.Avatar)
	}

	return fmt.Sprintf("https://cdn.discordapp.com/avatars/%s/%s.png", u.Id, u.Avatar)
}
