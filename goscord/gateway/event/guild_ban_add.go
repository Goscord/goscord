package event

import (
	"github.com/Goscord/goscord/goscord/discord"
	"github.com/Goscord/goscord/goscord/rest"
	"github.com/bytedance/sonic"
)

type GuildBanAdd struct {
	Data struct {
		GuildId string        `json:"guild_id"`
		User    *discord.User `json:"user"`
	} `json:"d"`
}

func NewGuildBanAdd(rest *rest.Client, data []byte) (*GuildBanAdd, error) {
	pk := new(GuildBanAdd)

	err := sonic.Unmarshal(data, pk)

	if err != nil {
		return nil, err
	}

	return pk, nil
}
