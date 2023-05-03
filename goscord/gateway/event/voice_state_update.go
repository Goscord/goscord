package event

import (
	"github.com/Goscord/goscord/goscord/discord"
	"github.com/Goscord/goscord/goscord/rest"
	"github.com/bytedance/sonic"
)

type VoiceStateUpdate struct {
	Data *discord.VoiceState `json:"d"`
}

func NewVoiceStateUpdate(rest *rest.Client, data []byte) (*VoiceStateUpdate, error) {
	pk := new(VoiceStateUpdate)

	err := sonic.Unmarshal(data, pk)

	if err != nil {
		return nil, err
	}

	return pk, nil
}
