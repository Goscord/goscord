package event

import (
	"github.com/Goscord/goscord/goscord/discord"
	"github.com/Goscord/goscord/goscord/rest"
	"github.com/goccy/go-json"
)

// GuildRoleUpdate Is sent when a guild role is updated.
type GuildRoleUpdate struct {
	Data *discord.GuildRoleUpdateEventFields `json:"d"`
}

func NewGuildRoleUpdate(rest *rest.Client, data []byte) (*GuildRoleUpdate, error) {
	pk := new(GuildRoleUpdate)

	err := json.Unmarshal(data, pk)

	if err != nil {
		return nil, err
	}

	return pk, nil
}
