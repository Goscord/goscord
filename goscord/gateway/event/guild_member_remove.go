package event

import (
	"github.com/Goscord/goscord/goscord/discord"
	"github.com/Goscord/goscord/goscord/rest"
	"github.com/bytedance/sonic"
)

type GuildMemberRemove struct {
	Data *discord.GuildMemberRemoveEventFields `json:"d"`
}

func NewGuildMemberRemove(rest *rest.Client, data []byte) (*GuildMemberRemove, error) {
	pk := new(GuildMemberRemove)

	err := sonic.Unmarshal(data, pk)

	if err != nil {
		return nil, err
	}

	return pk, nil
}
