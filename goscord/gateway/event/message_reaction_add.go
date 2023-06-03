package event

import (
	"github.com/Goscord/goscord/goscord/discord"
	"github.com/Goscord/goscord/goscord/rest"
	"github.com/bytedance/sonic"
)

type MessageReactionAdd struct {
	Data *discord.Reaction `json:"d"`
}

func NewMessageReactionAdd(rest *rest.Client, data []byte) (*MessageReactionAdd, error) {
	pk := new(MessageReactionAdd)

	err := sonic.Unmarshal(data, pk)

	if err != nil {
		return nil, err
	}

	return pk, nil
}
