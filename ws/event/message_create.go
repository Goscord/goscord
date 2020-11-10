package event

import (
	"encoding/json"

	"github.com/Seyz123/yalis/guild/message"
)

type MessageCreate struct {
	Data *message.Message `json:"d"`
}

func NewMessageCreate(data []byte) (*MessageCreate, error) {
	pk := new(MessageCreate)

	err := json.Unmarshal(data, pk)

	if err != nil {
		return nil, err
	}

	return pk, nil
}
