package rest

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/Goscord/goscord/discord"
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
