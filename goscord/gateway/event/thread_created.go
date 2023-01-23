package event

import (
	"github.com/Goscord/goscord/goscord/discord"
	"github.com/Goscord/goscord/goscord/rest"
	"github.com/goccy/go-json"
)

type ThreadCreate struct {
	Data *discord.Channel `json:"d"`
}

func NewThreadCreate(rest *rest.Client, data []byte) (*ThreadCreate, error) {
	pk := new(ThreadCreate)

	err := json.Unmarshal(data, pk)

	if err != nil {
		return nil, err
	}

	return pk, nil
}
