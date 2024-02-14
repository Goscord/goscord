package event

import (
	"github.com/bytedance/sonic"
	"github.com/Goscord/goscord/goscord/discord"
	"github.com/Goscord/goscord/goscord/rest"
)

type MessageReactionAdd struct {
	Data *discord.MessageReaction `json:"d"`
}

func NewMessageReactionAdd(rest *rest.Client, data []byte) (*MessageReactionAdd, error) {
	pk := new(MessageReactionAdd)

	err := sonic.Unmarshal(data, pk)

	if err != nil {
		return nil, err
	}
	return pk, nil
}