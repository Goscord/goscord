package user

import (
	"encoding/json"
	"fmt"
	"github.com/Goscord/goscord/rest"
)

type User struct {
	Rest          *rest.Client `json:"-"`
	Id            string       `json:"id"`
	Username      string       `json:"username"`
	Discriminator string       `json:"discriminator"`
	Avatar        string       `json:"avatar"`
	Bot           bool         `json:"bot"`
	System        bool         `json:"system"`
	MfaEnabled    bool         `json:"mfa_enabled"`
	Locale        string       `json:"locale"`
	Verified      bool         `json:"verified"`
	Email         string       `json:"email"`
	Flags         int          `json:"flags"`
	PremiumType   int          `json:"premium_type"`
	PublicFlags   int          `json:"public_flags"`
}

func NewUser(rest *rest.Client, data []byte) (*User, error) {
	u := new(User)

	err := json.Unmarshal(data, u)

	if err != nil {
		return nil, err
	}

	u.Rest = rest

	return u, nil
}

func (u *User) Tag() string {
	return fmt.Sprintf("%s#%s", u.Username, u.Discriminator)
}
