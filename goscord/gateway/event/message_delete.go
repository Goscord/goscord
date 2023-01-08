package event

import (
	"github.com/Goscord/goscord/goscord/rest"
	"github.com/goccy/go-json"
)

type MessageDelete struct {
	Data struct {
		Id        string `json:"id"`
		ChannelId string `json:"channel_id"`
		GuildId   string `json:"guild_id"`
	} `json:"d"`
}

func NewMessageDelete(_ *rest.Client, data []byte) (*MessageDelete, error) {
	pk := new(MessageDelete)

	err := json.Unmarshal(data, pk)

	if err != nil {
		return nil, err
	}

	return pk, nil
}
