package event

import (
	"encoding/json"
	"github.com/Goscord/goscord/discord"
	"github.com/Goscord/goscord/rest"
)

type GuildUpdate struct {
	Data *discord.Guild `json:"d"`
}

func NewGuildUpdate(rest *rest.Client, data []byte) (*GuildUpdate, error) {
	pk := new(GuildUpdate)

	err := json.Unmarshal(data, pk)

	if err != nil {
		return nil, err
	}

	pk.Data.Rest = rest

	return pk, nil
}
