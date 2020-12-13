package event

import (
	"encoding/json"
	"github.com/Goscord/goscord/discord"
	"github.com/Goscord/goscord/rest"
)

type GuildBanAdd struct {
	Data struct {
		GuildId string        `json:"guild_id"`
		User    *discord.User `json:"user"`
	} `json:"d"`
}

func NewGuildBanAdd(rest *rest.Client, data []byte) (*GuildBanAdd, error) {
	pk := new(GuildBanAdd)

	err := json.Unmarshal(data, pk)

	if err != nil {
		return nil, err
	}

	pk.Data.User.Rest = rest

	return pk, nil
}
