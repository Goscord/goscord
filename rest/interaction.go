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
