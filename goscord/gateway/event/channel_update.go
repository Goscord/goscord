package event

import (
	"github.com/Goscord/goscord/goscord/discord"
	"github.com/Goscord/goscord/goscord/rest"
	"github.com/bytedance/sonic"
)

type ChannelUpdate struct {
	Data *discord.Channel `json:"d"`
}

func NewChannelUpdate(rest *rest.Client, data []byte) (*ChannelUpdate, error) {
	pk := new(ChannelUpdate)

	err := sonic.Unmarshal(data, pk)

	if err != nil {
		return nil, err
	}

	return pk, nil
}
