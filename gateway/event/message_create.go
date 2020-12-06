package event

import (
	"encoding/json"
	"github.com/Seyz123/yalis/channel"
	"github.com/Seyz123/yalis/rest"
)

type MessageCreate struct {
	Data *channel.Message `json:"d"`
}

func NewMessageCreate(rest *rest.Client, data []byte) (*MessageCreate, error) {
	pk := new(MessageCreate)

	err := json.Unmarshal(data, pk)

	if err != nil {
		return nil, err
	}

	pk.Data.Rest = rest

	return pk, nil
}
