package event

import (
	"encoding/json"

	"github.com/Goscord/goscord/discord"
	"github.com/Goscord/goscord/rest"
)

type GuildCreate struct {
	Data *discord.Guild `json:"d"`
}

func NewGuildCreate(rest *rest.Client, data []byte) (*GuildCreate, error) {
	pk := new(GuildCreate)

	err := json.Unmarshal(data, pk)

	if err != nil {
		return nil, err
	}

	return pk, nil
}
