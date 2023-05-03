package event

import (
	"github.com/Goscord/goscord/goscord/discord"
	"github.com/Goscord/goscord/goscord/rest"
	"github.com/bytedance/sonic"
)

// GuildRoleUpdate Is sent when a guild role is updated.
type GuildRoleUpdate struct {
	Data *discord.GuildRoleUpdateEventFields `json:"d"`
}

func NewGuildRoleUpdate(rest *rest.Client, data []byte) (*GuildRoleUpdate, error) {
	pk := new(GuildRoleUpdate)

	err := sonic.Unmarshal(data, pk)

	if err != nil {
		return nil, err
	}

	return pk, nil
}
