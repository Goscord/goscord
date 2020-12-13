package event

import (
	"encoding/json"
	"github.com/Goscord/goscord/discord"
	"github.com/Goscord/goscord/rest"
)

type GuildEmojisUpdate struct {
	Data struct {
		GuildId string           `json:"guild_id"`
		Emojis  []*discord.Emoji `json:"emojis"`
	} `json:"d"`
}

func NewGuildEmojisUpdate(rest *rest.Client, data []byte) (*GuildEmojisUpdate, error) {
	pk := new(GuildEmojisUpdate)

	err := json.Unmarshal(data, pk)

	if err != nil {
		return nil, err
	}

	return pk, nil
}
