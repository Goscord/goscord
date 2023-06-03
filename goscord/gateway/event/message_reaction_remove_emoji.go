package event

import (
	"github.com/Goscord/goscord/goscord/discord"
	"github.com/Goscord/goscord/goscord/rest"
	"github.com/bytedance/sonic"
)

type MessageReactionRemoveEmoji struct {
	Data *discord.Reaction `json:"d"`
}

func NewMessageReactionRemoveEmoji(rest *rest.Client, data []byte) (*MessageReactionRemoveEmoji, error) {
	pk := new(MessageReactionRemoveEmoji)

	err := sonic.Unmarshal(data, pk)

	if err != nil {
		return nil, err
	}

	return pk, nil
}
