package event

import (
	"encoding/json"
	"github.com/Goscord/goscord/user"
)

type Ready struct {
	Data struct {
		Version int        `json:"v"`
		User    *user.User `json:"user"`
		// Guilds []*Guild `json:"guilds"`
		SessionID string `json:"session_id"`
		Shard     []int  `json:"shard,omitempty"`
	} `json:"d"`
}

func NewReady(data []byte) (*Ready, error) {
	pk := new(Ready)

	err := json.Unmarshal(data, pk)

	if err != nil {
		return nil, err
	}

	return pk, nil
}
