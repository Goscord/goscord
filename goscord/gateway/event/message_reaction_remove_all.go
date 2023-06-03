package event

import (
	"github.com/Goscord/goscord/goscord/discord"
	"github.com/Goscord/goscord/goscord/rest"
	"github.com/bytedance/sonic"
)

type MessageReactionRemoveAll struct {
	Data *discord.Reaction `json:"d"`
}

func NewMessageReactionRemoveAll(rest *rest.Client, data []byte) (*MessageReactionRemoveAll, error) {
	pk := new(MessageReactionRemoveAll)

	err := sonic.Unmarshal(data, pk)

	if err != nil {
		return nil, err
	}

	return pk, nil
}
