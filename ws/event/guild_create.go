package event

import (
	"encoding/json"

	"github.com/Seyz123/yalis/guild"
)

type GuildCreate struct {
	Data *guild.Guild `json:"d"`
}

func NewGuildCreate(data []byte) (*GuildCreate, error) {
	pk := new(GuildCreate)

	err := json.Unmarshal(data, pk)

	if err != nil {
		return nil, err
	}

	return pk, nil
}
