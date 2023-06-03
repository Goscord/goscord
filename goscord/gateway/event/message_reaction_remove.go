package event

import (
	"github.com/Goscord/goscord/goscord/discord"
	"github.com/Goscord/goscord/goscord/rest"
	"github.com/bytedance/sonic"
)

type MessageReactionRemove struct {
	Data *discord.Reaction `json:"d"`
}

func NewMessageReactionRemove(rest *rest.Client, data []byte) (*MessageReactionRemove, error) {
	pk := new(MessageReactionRemove)

	err := sonic.Unmarshal(data, pk)

	if err != nil {
		return nil, err
	}

	return pk, nil
}
