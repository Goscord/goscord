package event

import (
	"encoding/json"

	"github.com/Goscord/goscord/discord"
	"github.com/Goscord/goscord/rest"
)

type InteractionCreate struct {
	Data *discord.Interaction `json:"d"`
}

func NewInteractionCreate(rest *rest.Client, data []byte) (*InteractionCreate, error) {
	pk := new(InteractionCreate)

	err := json.Unmarshal(data, pk)

	if err != nil {
		return nil, err
	}

	return pk, nil
}
