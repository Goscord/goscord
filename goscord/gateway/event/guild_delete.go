package event

import (
	"github.com/Goscord/goscord/goscord/discord"
	"github.com/Goscord/goscord/goscord/rest"
	"github.com/goccy/go-json"
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

	return pk, nil
}
