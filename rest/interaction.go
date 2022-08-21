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

// CreateResponse creates a response to an interaction
func (ch *InteractionHandler) CreateResponse(interactionId, interactionToken string, content interface{}) (*discord.InteractionResponse, error) {
	b, err := formatMessage(content)

	if err != nil {
		return nil, err
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

// GetResponse gets the initial response of an interaction
func (ch *InteractionHandler) GetResponse(applicationId, interactionToken string) (*discord.InteractionResponse, error) {
	res, err := ch.rest.Request(fmt.Sprintf(EndpointGetInteractionResponse, applicationId, interactionToken), "GET", nil, "application/json")

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

// EditResponse edits the response of an interaction
func (ch *InteractionHandler) EditResponse(applicationId, interactionToken string, content interface{}) (*discord.InteractionResponse, error) {
	b, err := formatMessage(content)

	if err != nil {
		return nil, err
	}

	res, err := ch.rest.Request(fmt.Sprintf(EndpointEditInteractionResponse, applicationId, interactionToken), "PATCH", b, "application/json")

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

// DeleteResponse deletes the response of an interaction
func (ch *InteractionHandler) DeleteResponse(applicationId, interactionToken string) error {
	_, err := ch.rest.Request(fmt.Sprintf(EndpointDeleteInteractionResponse, applicationId, interactionToken), "DELETE", nil, "application/json")

	if err != nil {
		return err
	}

	return nil
}

// CreateFollowupMessage creates a followup message for an Interaction
func (ch *InteractionHandler) CreateFollowupMessage(applicationId, interactionToken string, content interface{}) (*discord.Message, error) {
	b, err := formatMessage(content)

	if err != nil {
		return nil, err
	}

	res, err := ch.rest.Request(fmt.Sprintf(EndpointCreateFollowupMessage, applicationId, interactionToken), "POST", b, "application/json")

	if err != nil {
		return nil, err
	}

	var message discord.Message
	err = json.Unmarshal(res, &message)

	if err != nil {
		return nil, err
	}

	return &message, nil
}

// GetFollowupMessage gets the followup message of an interaction
func (ch *InteractionHandler) GetFollowupMessage(applicationId, interactionToken, messageId string) (*discord.Message, error) {
	res, err := ch.rest.Request(fmt.Sprintf(EndpointGetFollowupMessage, applicationId, interactionToken, messageId), "GET", nil, "application/json")

	if err != nil {
		return nil, err
	}

	var message discord.Message
	err = json.Unmarshal(res, &message)

	if err != nil {
		return nil, err
	}

	return &message, nil
}

// EditFollowupMessage edits the followup message of an interaction
func (ch *InteractionHandler) EditFollowupMessage(applicationId, interactionToken, messageId string, content interface{}) (*discord.Message, error) {
	b, err := formatMessage(content)

	if err != nil {
		return nil, err
	}

	res, err := ch.rest.Request(fmt.Sprintf(EndpointEditFollowupMessage, applicationId, interactionToken, messageId), "PATCH", b, "application/json")

	if err != nil {
		return nil, err
	}

	var message discord.Message
	err = json.Unmarshal(res, &message)

	if err != nil {
		return nil, err
	}

	return &message, nil
}

// DeleteFollowupMessage deletes the followup message of an interaction
func (ch *InteractionHandler) DeleteFollowupMessage(applicationId, interactionToken, messageId string) error {
	_, err := ch.rest.Request(fmt.Sprintf(EndpointDeleteFollowupMessage, applicationId, interactionToken, messageId), "DELETE", nil, "application/json")

	if err != nil {
		return err
	}

	return nil
}

// formatMessage formats the message to be sent to the API it avoids code duplication
func formatMessage(content interface{}) (*bytes.Buffer, error) {
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

	return b, nil
}
