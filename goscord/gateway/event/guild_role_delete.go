package event

import (
	"github.com/Goscord/goscord/goscord/discord"
	"github.com/Goscord/goscord/goscord/rest"
	"github.com/goccy/go-json"
)

// GuildRoleDelete Is sent when a guild role is created.
type GuildRoleDelete struct {
	Data *discord.GuildRoleDeleteEventFields `json:"d"`
}

func NewGuildRoleDelete(rest *rest.Client, data []byte) (*GuildRoleDelete, error) {
	pk := new(GuildRoleDelete)

	err := json.Unmarshal(data, pk)

	if err != nil {
		return nil, err
	}

	return pk, nil
}
