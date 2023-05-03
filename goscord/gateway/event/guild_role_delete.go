package event

import (
	"github.com/Goscord/goscord/goscord/discord"
	"github.com/Goscord/goscord/goscord/rest"
	"github.com/bytedance/sonic"
)

// GuildRoleDelete Is sent when a guild role is deleted.
type GuildRoleDelete struct {
	Data *discord.GuildRoleDeleteEventFields `json:"d"`
}

func NewGuildRoleDelete(rest *rest.Client, data []byte) (*GuildRoleDelete, error) {
	pk := new(GuildRoleDelete)

	err := sonic.Unmarshal(data, pk)

	if err != nil {
		return nil, err
	}

	return pk, nil
}
