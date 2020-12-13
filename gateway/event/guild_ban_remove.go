package event

import (
	"encoding/json"
	"github.com/Goscord/goscord/discord"
	"github.com/Goscord/goscord/rest"
)

type GuildBanRemove struct {
	Data struct {
		GuildId string        `json:"guild_id"`
		User    *discord.User `json:"user"`
	} `json:"d"`
}

func NewGuildBanRemove(rest *rest.Client, data []byte) (*GuildBanRemove, error) {
	pk := new(GuildBanRemove)

	err := json.Unmarshal(data, pk)

	if err != nil {
		return nil, err
	}

	return pk, nil
}
