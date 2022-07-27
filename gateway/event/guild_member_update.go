package event

import (
	"encoding/json"

	"github.com/Goscord/goscord/discord"
	"github.com/Goscord/goscord/rest"
)

type GuildMemberUpdate struct {
	Data *discord.Member `json:"d"`
}

func NewGuildMemberUpdate(rest *rest.Client, data []byte) (*GuildMemberUpdate, error) {
	pk := new(GuildMemberUpdate)

	err := json.Unmarshal(data, pk)

	if err != nil {
		return nil, err
	}

	return pk, nil
}
