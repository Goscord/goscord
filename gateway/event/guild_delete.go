package event

import (
	"encoding/json"
	"github.com/Goscord/goscord/discord"
	"github.com/Goscord/goscord/rest"
)

type GuildDelete struct {
	Data *discord.Guild `json:"d"`
}

func NewGuildDelete(rest *rest.Client, data []byte) (*GuildDelete, error) {
	pk := new(GuildDelete)

	err := json.Unmarshal(data, pk)

	if err != nil {
		return nil, err
	}

	pk.Data.Rest = rest

	return pk, nil
}
