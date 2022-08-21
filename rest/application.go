package rest

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/Goscord/goscord/discord"
)

type ApplicationHandler struct {
	rest *Client
}

func NewApplicationHandler(rest *Client) *ApplicationHandler {
	return &ApplicationHandler{rest: rest}
}

func (ch *ApplicationHandler) GetCommands(applicationId, guildId string) ([]*discord.ApplicationCommand, error) {
	var endpoint string
	var commands []*discord.ApplicationCommand

	if guildId == "" {
		endpoint = fmt.Sprintf(EndpointGetGlobalApplicationCommands, applicationId)
	} else {
		endpoint = fmt.Sprintf(EndpointGetGuildApplicationCommands, applicationId, guildId)
	}

	res, err := ch.rest.Request(endpoint, "GET", bytes.NewBufferString(""), "application/json")

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(res, &commands)

	if err != nil {
		return nil, err
	}

	return commands, nil
}

func (ch *ApplicationHandler) RegisterCommand(applicationId, guildId string, application *discord.ApplicationCommand) error {
	var endpoint string

	if guildId == "" {
		endpoint = fmt.Sprintf(EndpointCreateGlobalApplicationCommand, applicationId)
	} else {
		endpoint = fmt.Sprintf(EndpointCreateGuildApplicationCommand, applicationId, guildId)
	}

	data, err := json.Marshal(application)

	if err != nil {
		return err
	}

	_, err = ch.rest.Request(endpoint, "POST", bytes.NewBuffer(data), "application/json")

	return err
}

func (ch *ApplicationHandler) GetCommand(applicationId, guildId, commandId string) (*discord.ApplicationCommand, error) {
	var endpoint string
	var command *discord.ApplicationCommand

	if guildId == "" {
		endpoint = fmt.Sprintf(EndpointGetGlobalApplicationCommand, applicationId, commandId)
	} else {
		endpoint = fmt.Sprintf(EndpointGetGuildApplicationCommand, applicationId, guildId, commandId)
	}

	res, err := ch.rest.Request(endpoint, "GET", bytes.NewBufferString(""), "application/json")

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(res, &command)

	if err != nil {
		return nil, err
	}

	return command, nil
}

// ToDo : UnregisterCommand
// ToDo : UpdateCommand
