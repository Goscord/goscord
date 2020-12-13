package event

import (
	"encoding/json"
	"github.com/Goscord/goscord/discord"
	"github.com/Goscord/goscord/rest"
)

type ChannelDelete struct {
	Data *discord.Channel `json:"d"`
}

func NewChannelDelete(rest *rest.Client, data []byte) (*ChannelDelete, error) {
	pk := new(ChannelDelete)

	err := json.Unmarshal(data, pk)

	if err != nil {
		return nil, err
	}

	return pk, nil
}
