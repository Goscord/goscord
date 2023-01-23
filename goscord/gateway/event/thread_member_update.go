package event

import (
	"github.com/Goscord/goscord/goscord/discord"
	"github.com/Goscord/goscord/goscord/rest"
	"github.com/goccy/go-json"
)

type ThreadMemberUpdate struct {
	Data *discord.ThreadMember `json:"d"`
}

func NewThreadMemberUpdate(rest *rest.Client, data []byte) (*ThreadMemberUpdate, error) {
	pk := new(ThreadMemberUpdate)

	err := json.Unmarshal(data, pk)

	if err != nil {
		return nil, err
	}

	return pk, nil
}
