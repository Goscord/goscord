package event

import (
	"github.com/Goscord/goscord/goscord/discord"
	"github.com/Goscord/goscord/goscord/rest"
	"github.com/bytedance/sonic"
)

type GuildMemberAdd struct {
	Data *discord.GuildMember `json:"d"`
}

func NewGuildMemberAdd(rest *rest.Client, data []byte) (*GuildMemberAdd, error) {
	pk := new(GuildMemberAdd)

	err := sonic.Unmarshal(data, pk)

	if err != nil {
		return nil, err
	}

	return pk, nil
}
