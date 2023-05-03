package event

import (
	"github.com/Goscord/goscord/goscord/discord"
	"github.com/Goscord/goscord/goscord/rest"
	"github.com/bytedance/sonic"
)

type GuildCreate struct {
	Data *discord.Guild `json:"d"`
}

func NewGuildCreate(rest *rest.Client, data []byte) (*GuildCreate, error) {
	pk := new(GuildCreate)

	err := sonic.Unmarshal(data, pk)

	if err != nil {
		return nil, err
	}

	return pk, nil
}
