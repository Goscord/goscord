package event

import (
	"encoding/json"
	"github.com/Seyz123/yalis/rest"

	"github.com/Seyz123/yalis/channel/message"
)

type MessageCreate struct {
	Data *message.Message `json:"d"`
}

func NewMessageCreate(rest *rest.RestClient, data []byte) (*MessageCreate, error) {
	pk := new(MessageCreate)

	err := json.Unmarshal(data, pk)

	if err != nil {
		return nil, err
	}

	pk.Data.Rest = rest

	return pk, nil
}
