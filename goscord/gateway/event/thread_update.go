package event

import (
	"github.com/Goscord/goscord/goscord/discord"
	"github.com/Goscord/goscord/goscord/rest"
	"github.com/bytedance/sonic"
)

type ThreadUpdate struct {
	Data *discord.Channel `json:"d"`
}

func NewThreadUpdate(rest *rest.Client, data []byte) (*ThreadUpdate, error) {
	pk := new(ThreadUpdate)

	err := sonic.Unmarshal(data, pk)

	if err != nil {
		return nil, err
	}

	return pk, nil
}
