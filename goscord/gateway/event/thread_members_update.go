package event

import (
	"github.com/Goscord/goscord/goscord/discord"
	"github.com/Goscord/goscord/goscord/rest"
	"github.com/bytedance/sonic"
)

type ThreadMembersUpdate struct {
	Data *discord.ThreadMembersUpdateEventFields `json:"d"`
}

func NewThreadMembersUpdate(rest *rest.Client, data []byte) (*ThreadMembersUpdate, error) {
	pk := new(ThreadMembersUpdate)

	err := sonic.Unmarshal(data, pk)

	if err != nil {
		return nil, err
	}

	return pk, nil
}
