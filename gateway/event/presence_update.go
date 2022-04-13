package event

import (
	"encoding/json"
	"github.com/Goscord/goscord/discord"
	"github.com/Goscord/goscord/rest"
)

type PresenceUpdate struct {
	Data struct {
		User         *discord.User       `json:"user"`
		GuildId      string              `json:"guild_id"`
		Status       string              `json:"status"`
		Activities   []*discord.Activity `json:"activities"`
		ClientStatus string              `json:"client_status"`
	} `json:"d"`
}

func NewPresenceUpdate(_ *rest.Client, data []byte) (*PresenceUpdate, error) {
	pk := new(PresenceUpdate)

	err := json.Unmarshal(data, pk)

	if err != nil {
		return nil, err
	}

	return pk, nil
}
