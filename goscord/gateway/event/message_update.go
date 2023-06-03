package event

import (
	"github.com/Goscord/goscord/goscord/discord"
	"github.com/Goscord/goscord/goscord/rest"
	"github.com/bytedance/sonic"
)

type MessageUpdate struct {
	Data *discord.Message `json:"d"`
}

func NewMessageUpdate(rest *rest.Client, data []byte) (*MessageUpdate, error) {
	pk := new(MessageUpdate)

	err := sonic.Unmarshal(data, pk)

	if err != nil {
		return nil, err
	}

	return pk, nil
}
