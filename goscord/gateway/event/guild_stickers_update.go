package event

import (
	"github.com/Goscord/goscord/goscord/discord"
	"github.com/Goscord/goscord/goscord/rest"
	"github.com/bytedance/sonic"
)

type GuildStickersUpdate struct {
	Data *discord.GuildStickersUpdateEventFields `json:"d"`
}

func NewGuildStickersUpdate(rest *rest.Client, data []byte) (*GuildStickersUpdate, error) {
	pk := new(GuildStickersUpdate)

	err := sonic.Unmarshal(data, pk)

	if err != nil {
		return nil, err
	}

	return pk, nil
}
