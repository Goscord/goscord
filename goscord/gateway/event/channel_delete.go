package event

import (
	"github.com/Goscord/goscord/goscord/discord"
	"github.com/Goscord/goscord/goscord/rest"
	"github.com/bytedance/sonic"
)

type ChannelDelete struct {
	Data *discord.Channel `json:"d"`
}

func NewChannelDelete(rest *rest.Client, data []byte) (*ChannelDelete, error) {
	pk := new(ChannelDelete)

	err := sonic.Unmarshal(data, pk)

	if err != nil {
		return nil, err
	}

	return pk, nil
}
