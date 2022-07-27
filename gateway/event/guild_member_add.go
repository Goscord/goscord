package event

import (
	"encoding/json"

	"github.com/Goscord/goscord/discord"
	"github.com/Goscord/goscord/rest"
)

type GuildMemberAdd struct {
	Data *discord.Member `json:"d"`
}

func NewGuildMemberAdd(rest *rest.Client, data []byte) (*GuildMemberAdd, error) {
	pk := new(GuildMemberAdd)

	err := json.Unmarshal(data, pk)

	if err != nil {
		return nil, err
	}

	return pk, nil
}
