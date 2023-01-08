package event

import (
	"github.com/Goscord/goscord/goscord/discord"
	"github.com/Goscord/goscord/goscord/rest"
	"github.com/goccy/go-json"
)

type MessageCreate struct {
	Data *discord.Message `json:"d"`
}

func NewMessageCreate(rest *rest.Client, data []byte) (*MessageCreate, error) {
	pk := new(MessageCreate)

	err := json.Unmarshal(data, pk)

	if err != nil {
		return nil, err
	}

	return pk, nil
}
