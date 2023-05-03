package event

import (
	"github.com/Goscord/goscord/goscord/discord"
	"github.com/Goscord/goscord/goscord/rest"
	"github.com/bytedance/sonic"
)

type ThreadListSync struct {
	Data *discord.ThreadListSyncEventFields `json:"d"`
}

func NewThreadListSync(rest *rest.Client, data []byte) (*ThreadListSync, error) {
	pk := new(ThreadListSync)

	err := sonic.Unmarshal(data, pk)

	if err != nil {
		return nil, err
	}

	return pk, nil
}
