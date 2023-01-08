package rest

import (
	"bytes"
	"fmt"
	discord2 "github.com/Goscord/goscord/goscord/discord"
	"github.com/Goscord/goscord/goscord/discord/embed"

	"github.com/goccy/go-json"
)

type InteractionHandler struct {
	rest *Client
}

func NewInteractionHandler(rest *Client) *InteractionHandler {
	return &InteractionHandler{rest: rest}
}

// CreateResponse creates a response to an interaction
func (ch *InteractionHandler) CreateResponse(interactionId, interactionToken string, content interface{}) error {
	b, err := ch.formatMessage(content)

	if err != nil {
		return err
	}

	_, err = ch.rest.Request(fmt.Sprintf(EndpointCreateInteractionResponse, interactionId, interactionToken), "POST", b, "application/json")

	if err != nil {
		return err
	}

	return nil

}

// GetOriginalResponse GetResponse gets the initial response of an interaction
func (ch *InteractionHandler) GetOriginalResponse(applicationId, interactionToken string) (*discord2.Message, error) {
	res, err := ch.rest.Request(fmt.Sprintf(EndpointGetInteractionResponse, applicationId, interactionToken), "GET", nil, "application/json")

	if err != nil {
		return nil, err
	}

	var response *discord2.Message
	err = json.Unmarshal(res, &response)

	if err != nil {
		return nil, err
	}

	return response, nil
}

// EditOriginalResponse EditResponse edits the response of an interaction
func (ch *InteractionHandler) EditOriginalResponse(applicationId, interactionToken string, content interface{}) (*discord2.Message, error) {
	b, err := ch.formatMessage(content)

	if err != nil {
		return nil, err
	}

	res, err := ch.rest.Request(fmt.Sprintf(EndpointEditInteractionResponse, applicationId, interactionToken), "PATCH", b, "application/json")

	if err != nil {
		return nil, err
	}

	var response *discord2.Message
	err = json.Unmarshal(res, &response)

	if err != nil {
		return nil, err
	}

	return response, nil
}

// DeleteOriginalResponse DeleteResponse deletes the response of an interaction
func (ch *InteractionHandler) DeleteOriginalResponse(applicationId, interactionToken string) error {
	_, err := ch.rest.Request(fmt.Sprintf(EndpointDeleteInteractionResponse, applicationId, interactionToken), "DELETE", nil, "application/json")

	if err != nil {
		return err
	}

	return nil
}

// CreateFollowupMessage creates a followup message for an Interaction
func (ch *InteractionHandler) CreateFollowupMessage(applicationId, interactionToken string, content interface{}) (*discord2.Message, error) {
	b, err := ch.formatMessage(content)

	if err != nil {
		return nil, err
	}

	res, err := ch.rest.Request(fmt.Sprintf(EndpointCreateFollowupMessage, applicationId, interactionToken), "POST", b, "application/json")

	if err != nil {
		return nil, err
	}

	var message *discord2.Message
	err = json.Unmarshal(res, &message)

	if err != nil {
		return nil, err
	}

	return message, nil
}

// GetFollowupMessage gets the followup message of an interaction
func (ch *InteractionHandler) GetFollowupMessage(applicationId, interactionToken, messageId string) (*discord2.Message, error) {
	res, err := ch.rest.Request(fmt.Sprintf(EndpointGetFollowupMessage, applicationId, interactionToken, messageId), "GET", nil, "application/json")

	if err != nil {
		return nil, err
	}

	var message *discord2.Message
	err = json.Unmarshal(res, &message)

	if err != nil {
		return nil, err
	}

	return message, nil
}

// EditFollowupMessage edits the followup message of an interaction
func (ch *InteractionHandler) EditFollowupMessage(applicationId, interactionToken, messageId string, content interface{}) (*discord2.Message, error) {
	b, err := ch.formatMessage(content)

	if err != nil {
		return nil, err
	}

	res, err := ch.rest.Request(fmt.Sprintf(EndpointEditFollowupMessage, applicationId, interactionToken, messageId), "PATCH", b, "application/json")

	if err != nil {
		return nil, err
	}

	var message *discord2.Message
	err = json.Unmarshal(res, &message)

	if err != nil {
		return nil, err
	}

	return message, nil
}

// DeleteFollowupMessage deletes the followup message of an interaction
func (ch *InteractionHandler) DeleteFollowupMessage(applicationId, interactionToken, messageId string) error {
	_, err := ch.rest.Request(fmt.Sprintf(EndpointDeleteFollowupMessage, applicationId, interactionToken, messageId), "DELETE", nil, "application/json")

	if err != nil {
		return err
	}

	return nil
}

// formatMessage formats the message to be sent to the API it avoids code duplication. ToDo : Create a custom type for it
func (ch *InteractionHandler) formatMessage(content interface{}) (*bytes.Buffer, error) {
	b := new(bytes.Buffer)

	switch ccontent := content.(type) {
	case string:
		content = &discord2.InteractionResponse{
			Type: discord2.InteractionCallbackTypeChannelWithSource,
			Data: &discord2.InteractionCallbackMessage{Content: ccontent},
		}
		jsonb, err := json.Marshal(content)

		if err != nil {
			return nil, err
		}

		b = bytes.NewBuffer(jsonb)

	case *embed.Embed:
		content = &discord2.InteractionResponse{
			Type: discord2.InteractionCallbackTypeChannelWithSource,
			Data: &discord2.InteractionCallbackMessage{Embeds: []*embed.Embed{ccontent}},
		}
		jsonb, err := json.Marshal(content)

		if err != nil {
			return nil, err
		}

		b = bytes.NewBuffer(jsonb)

	case *discord2.InteractionCallbackMessage:
		content = &discord2.InteractionResponse{
			Type: discord2.InteractionCallbackTypeChannelWithSource,
			Data: ccontent,
		}
		jsonb, err := json.Marshal(content)

		if err != nil {
			return nil, err
		}

		b = bytes.NewBuffer(jsonb)

	case *discord2.InteractionCallbackAutocomplete:
		content = &discord2.InteractionResponse{
			Type: discord2.InteractionCallbackTypeApplicationCommandAutocompleteResult,
			Data: ccontent,
		}

		jsonb, err := json.Marshal(content)

		if err != nil {
			return nil, err
		}

		b = bytes.NewBuffer(jsonb)

	case *discord2.InteractionCallbackModal:
		content = &discord2.InteractionResponse{
			Type: discord2.InteractionCallbackTypeModal,
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
