package event

import (
	"github.com/Goscord/goscord/goscord/discord"
	"github.com/Goscord/goscord/goscord/rest"
	"github.com/goccy/go-json"
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

	return pk, nil
}
