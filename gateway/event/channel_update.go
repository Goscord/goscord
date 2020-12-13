package event

import (
	"encoding/json"
	"github.com/Goscord/goscord/discord"
	"github.com/Goscord/goscord/rest"
)

type ChannelUpdate struct {
	Data *discord.Channel `json:"d"`
}

func NewChannelUpdate(rest *rest.Client, data []byte) (*ChannelUpdate, error) {
	pk := new(ChannelUpdate)

	err := json.Unmarshal(data, pk)

	if err != nil {
		return nil, err
	}

	return pk, nil
}
