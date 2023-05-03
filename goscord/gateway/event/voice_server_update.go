package event

import (
	"github.com/Goscord/goscord/goscord/discord"
	"github.com/Goscord/goscord/goscord/rest"
	"github.com/bytedance/sonic"
)

type VoiceServerUpdate struct {
	Data *discord.VoiceServerUpdateEventFields `json:"d"`
}

func NewVoiceServerUpdate(rest *rest.Client, data []byte) (*VoiceServerUpdate, error) {
	pk := new(VoiceServerUpdate)

	err := sonic.Unmarshal(data, pk)

	if err != nil {
		return nil, err
	}

	return pk, nil
}
