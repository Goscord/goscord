package event

import (
	"github.com/Goscord/goscord/goscord/discord"
	"github.com/Goscord/goscord/goscord/rest"
	"github.com/goccy/go-json"
)

type GuildMemberUpdate struct {
	Data *discord.GuildMember `json:"d"`
}

func NewGuildMemberUpdate(rest *rest.Client, data []byte) (*GuildMemberUpdate, error) {
	pk := new(GuildMemberUpdate)

	err := json.Unmarshal(data, pk)

	if err != nil {
		return nil, err
	}

	return pk, nil
}
