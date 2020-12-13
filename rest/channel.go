package rest

import (
	"encoding/json"
	"fmt"
	"github.com/Goscord/goscord/discord"
	"github.com/Goscord/goscord/discord/embed"
)

type ChannelHandler struct {
	rest *Client
}

func NewChannelHandler(rest *Client) *ChannelHandler {
	return &ChannelHandler{rest: rest}
}

func (ch *ChannelHandler) GetMessage(channelId, id string) (*discord.Message, error) {
	data, err := ch.rest.Request(fmt.Sprintf(EndpointGetMessage, channelId, id), "GET", nil)

	if err != nil {
		return nil, err
	}

	var msg discord.Message
	err = json.Unmarshal(data, &msg)

	if err != nil {
		return nil, err
	}

	return &msg, nil
}

func (ch *ChannelHandler) Send(channelId string, content interface{}) (*discord.Message, error) {
	switch content.(type) {
	case string:
		content = map[string]string{"content": content.(string)}

	case *embed.Embed:
		content = &embed.MessageEmbed{Embed: content.(*embed.Embed)}
	}

	b, err := json.Marshal(content)

	if err != nil {
		return nil, err
	}

	res, err := ch.rest.Request(fmt.Sprintf(EndpointCreateMessage, channelId), "POST", b)

	if err != nil {
		return nil, err
	}

	msg := new(discord.Message)

	err = json.Unmarshal(res, msg)

	if err != nil {
		return nil, err
	}

	return msg, nil
}

// TODO
