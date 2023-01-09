package event

import (
	"github.com/Goscord/goscord/goscord/discord"
	"github.com/Goscord/goscord/goscord/rest"
	"github.com/goccy/go-json"
)

type ThreadUpdate struct {
	Data *discord.Channel `json:"d"`
}

func NewThreadUpdate(rest *rest.Client, data []byte) (*ThreadUpdate, error) {
	pk := new(ThreadUpdate)

	err := json.Unmarshal(data, pk)

	if err != nil {
		return nil, err
	}

	return pk, nil
}
