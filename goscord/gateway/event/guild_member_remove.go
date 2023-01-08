package event

import (
	"github.com/Goscord/goscord/goscord/discord"
	"github.com/Goscord/goscord/goscord/rest"
	"github.com/goccy/go-json"
)

type GuildMemberRemove struct {
	Data *discord.User `json:"d"`
}

func NewGuildMemberRemove(rest *rest.Client, data []byte) (*GuildMemberRemove, error) {
	pk := new(GuildMemberRemove)

	err := json.Unmarshal(data, pk)

	if err != nil {
		return nil, err
	}

	return pk, nil
}
