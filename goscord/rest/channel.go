package rest

import (
	"bytes"
	"fmt"
	"github.com/Goscord/goscord/goscord/discord"
	"github.com/Goscord/goscord/goscord/discord/embed"
	"io"
	"mime/multipart"
	"os"

	"github.com/goccy/go-json"
)

type ChannelHandler struct {
	rest *Client
}

func NewChannelHandler(rest *Client) *ChannelHandler {
	return &ChannelHandler{rest: rest}
}

func (ch *ChannelHandler) GetChannel(channelId string) (*discord.Channel, error) {
	data, err := ch.rest.Request(fmt.Sprintf(EndpointGetChannel, channelId), "GET", nil, "application/json")

	if err != nil {
		return nil, err
	}

	var channel *discord.Channel
	err = json.Unmarshal(data, &channel)

	if err != nil {
		return nil, err
	}

	return channel, nil
}

// Send / reply to a message
func (ch *ChannelHandler) GetMessage(channelId, messageId string) (*discord.Message, error) {
	data, err := ch.rest.Request(fmt.Sprintf(EndpointGetChannelMessage, channelId, messageId), "GET", nil, "application/json")

	if err != nil {
		return nil, err
	}

	var msg *discord.Message
	err = json.Unmarshal(data, &msg)

	if err != nil {
		return nil, err
	}

	return msg, nil
}

func (ch *ChannelHandler) SendMessage(channelId string, content interface{}) (*discord.Message, error) {
	b, contentType, err := formatMessage(content, "")

	if err != nil {
		return nil, err
	}

	res, err := ch.rest.Request(fmt.Sprintf(EndpointCreateMessage, channelId), "POST", b, contentType)

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

func (ch *ChannelHandler) ReplyMessage(channelId, messageId string, content interface{}) (*discord.Message, error) {
	b, contentType, err := formatMessage(content, messageId)

	if err != nil {
		return nil, err
	}

	res, err := ch.rest.Request(fmt.Sprintf(EndpointCreateMessage, channelId), "POST", b, contentType)

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

func (ch *ChannelHandler) Edit(channelId, messageId string, content interface{}) (*discord.Message, error) {
	b, contentType, err := formatMessage(content, "")

	if err != nil {
		return nil, err
	}

	res, err := ch.rest.Request(fmt.Sprintf(EndpointEditMessage, channelId, messageId), "PATCH", b, contentType)

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

func (ch *ChannelHandler) CrosspostMessage(channelId, messageId string) (*discord.Message, error) {
	data, err := ch.rest.Request(fmt.Sprintf(EndpointCrosspostMessage, channelId, messageId), "POST", nil, "application/json")

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

// formatMessage formats the message to be sent to the API it avoids code duplication. // ToDo : Create a custom type for it
func formatMessage(content interface{}, messageId string) (*bytes.Buffer, string, error) {
	b := new(bytes.Buffer)
	contentType := "application/json"

	switch ccontent := content.(type) {
	case string:
		if messageId != "" {
			content = &discord.Message{Content: ccontent, MessageReference: &discord.MessageReference{MessageId: messageId}}
		} else {
			content = &discord.Message{Content: ccontent}
		}

		jsonb, err := json.Marshal(content)

		if err != nil {
			return nil, "", err
		}

		b = bytes.NewBuffer(jsonb)

	case *embed.Embed:
		if messageId != "" {
			content = &discord.Message{Embeds: []*embed.Embed{ccontent}, MessageReference: &discord.MessageReference{MessageId: messageId}}
		} else {
			content = &discord.Message{Embeds: []*embed.Embed{ccontent}}
		}

		jsonb, err := json.Marshal(content)

		if err != nil {
			return nil, "", err
		}

		b = bytes.NewBuffer(jsonb)

	case *discord.Message:
		jsonb, err := json.Marshal(content)

		if err != nil {
			return nil, "", err
		}

		b = bytes.NewBuffer(jsonb)

	case []*os.File:
		w := multipart.NewWriter(b)

		for i, file := range ccontent {
			fw, err := w.CreateFormFile(fmt.Sprintf("attachment-%d", i), file.Name())
			if err != nil {
				return nil, "", err
			}

			_, err = io.Copy(fw, file)
			if err != nil {
				return nil, "", err
			}
		}

		w.Close()
		contentType = w.FormDataContentType()
		fmt.Println(contentType)
	}

	return b, contentType, nil
}

// TODO
