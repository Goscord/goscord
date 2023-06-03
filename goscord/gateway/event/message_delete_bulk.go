package event

import (
	"github.com/Goscord/goscord/goscord/discord"
	"github.com/Goscord/goscord/goscord/rest"
	"github.com/bytedance/sonic"
)

type MessageDeleteBulk struct {
	Data *discord.MessageDeleteBulk `json:"d"`
}

func NewMessageDeleteBulk(_ *rest.Client, data []byte) (*MessageDeleteBulk, error) {
	pk := new(MessageDeleteBulk)

	err := sonic.Unmarshal(data, pk)

	if err != nil {
		return nil, err
	}

	return pk, nil
}
