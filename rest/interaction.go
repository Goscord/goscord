package rest

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/Goscord/goscord/discord"
	"github.com/Goscord/goscord/discord/embed"
)

type InteractionHandler struct {
	rest *Client
}

func NewInteractionHandler(rest *Client) *InteractionHandler {
	return &InteractionHandler{rest: rest}
}

func (ch *InteractionHandler) RegisterCommand(applicationId, guildId string, application *discord.ApplicationCommand) error {
	var endpoint string

	if guildId == "" {
		endpoint = fmt.Sprintf(EndpointRegisterGlobalCommand, applicationId)
	} else {
		endpoint = fmt.Sprintf(EndpointRegisterGuildCommand, applicationId, guildId)
	}

	data, err := json.Marshal(application)

	if err != nil {
		return err
	}

	_, err = ch.rest.Request(endpoint, "POST", bytes.NewBuffer(data), "application/json")

	return err
}

// ToDo : UnregisterCommand
// ToDo : UpdateCommand

func (ch *InteractionHandler) CreateResponse(interactionId, interactionToken string, content interface{}) (*discord.InteractionResponse, error) {
	var b *bytes.Buffer

	switch ccontent := content.(type) {
	case string:
		content = &discord.InteractionResponse{
			Type: discord.InteractionCallbackTypeChannelWithSource,
			Data: &discord.InteractionCallbackMessage{Content: ccontent, Flags: discord.MessageFlagEphemeral},
		}
		jsonb, err := json.Marshal(content)

		if err != nil {
			return nil, err
		}

		b = bytes.NewBuffer(jsonb)

	case *embed.Builder:
		content = &discord.InteractionResponse{
			Type: discord.InteractionCallbackTypeChannelWithSource,
			Data: &discord.InteractionCallbackMessage{Content: ccontent.Content(), Embeds: []*embed.Embed{ccontent.Embed()}},
		}
		jsonb, err := json.Marshal(content)

		if err != nil {
			return nil, err
		}

		b = bytes.NewBuffer(jsonb)

	case *embed.Embed:
		content = &discord.InteractionResponse{
			Type: discord.InteractionCallbackTypeChannelWithSource,
			Data: &discord.InteractionCallbackMessage{Embeds: []*embed.Embed{ccontent}},
		}
		jsonb, err := json.Marshal(content)

		if err != nil {
			return nil, err
		}

		b = bytes.NewBuffer(jsonb)

	case *discord.InteractionCallbackMessage:
		content = &discord.InteractionResponse{
			Type: discord.InteractionCallbackTypeChannelWithSource,
			Data: ccontent,
		}
		jsonb, err := json.Marshal(content)

		if err != nil {
			return nil, err
		}

		b = bytes.NewBuffer(jsonb)

	case *discord.InteractionCallbackAutocomplete:
		content = &discord.InteractionResponse{
			Type: discord.InteractionCallbackTypeApplicationCommandAutocompleteResult,
			Data: ccontent,
		}

		jsonb, err := json.Marshal(content)

		if err != nil {
			return nil, err
		}

		b = bytes.NewBuffer(jsonb)

	case *discord.InteractionCallbackModal:
		content = &discord.InteractionResponse{
			Type: discord.InteractionCallbackTypeModal,
			Data: ccontent,
		}
		jsonb, err := json.Marshal(content)

		if err != nil {
			return nil, err
		}

		b = bytes.NewBuffer(jsonb)
	}

	res, err := ch.rest.Request(fmt.Sprintf(EndpointCreateInteractionResponse, interactionId, interactionToken), "POST", b, "application/json")

	if err != nil {
		return nil, err
	}

	var response discord.InteractionResponse
	err = json.Unmarshal(res, &response)

	if err != nil {
		return nil, err
	}

	return &response, nil
}
