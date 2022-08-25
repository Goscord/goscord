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

// GetCommands fetchs all of the commands for your application
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

// RegisterCommand creates a new application command
func (ch *ApplicationHandler) RegisterCommand(applicationId, guildId string, application *discord.ApplicationCommand) (*discord.ApplicationCommand, error) {
	var endpoint string
	var command *discord.ApplicationCommand

	if guildId == "" {
		endpoint = fmt.Sprintf(EndpointCreateGlobalApplicationCommand, applicationId)
	} else {
		endpoint = fmt.Sprintf(EndpointCreateGuildApplicationCommand, applicationId, guildId)
	}

	data, err := json.Marshal(application)

	if err != nil {
		return nil, err
	}

	res, err := ch.rest.Request(endpoint, "POST", bytes.NewBuffer(data), "application/json")

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(res, &command)

	if err != nil {
		return nil, err
	}

	return command, nil
}

// GetCommand fetchs a command for your application
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

// EditCommand edits a command
func (ch *ApplicationHandler) EditCommand(applicationId, guildId, commandId string, command *discord.ApplicationCommand) (*discord.ApplicationCommand, error) {
	var endpoint string

	if guildId == "" {
		endpoint = fmt.Sprintf(EndpointEditGlobalApplicationCommand, applicationId, commandId)
	} else {
		endpoint = fmt.Sprintf(EndpointEditGuildApplicationCommand, applicationId, guildId, commandId)
	}

	res, err := ch.rest.Request(endpoint, "PUT", bytes.NewBufferString(""), "application/json")

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(res, &command)

	if err != nil {
		return nil, err
	}

	return command, nil
}

// DeleteCommand deletes a command
func (ch *ApplicationHandler) DeleteCommand(applicationId, guildId, commandId string) error {
	var endpoint string

	if guildId == "" {
		endpoint = fmt.Sprintf(EndpointDeleteGlobalApplicationCommand, applicationId, commandId)
	} else {
		endpoint = fmt.Sprintf(EndpointDeleteGuildApplicationCommand, applicationId, guildId, commandId)
	}

	_, err := ch.rest.Request(endpoint, "DELETE", bytes.NewBufferString(""), "application/json")

	return err
}

// ToDo : BulkOverwriteCommands

// GetGuildCommandPermissions fetches permissions for all commands for your application in a guild
func (ch *ApplicationHandler) GetGuildCommandPermissions(applicationId, guildId string) ([]*discord.GuildApplicationCommandPermissions, error) {
	var permissions []*discord.GuildApplicationCommandPermissions

	res, err := ch.rest.Request(fmt.Sprintf(EndpointGetGuildApplicationCommandPermissions, applicationId, guildId), "GET", bytes.NewBufferString(""), "application/json")

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(res, &permissions)

	if err != nil {
		return nil, err
	}

	return permissions, nil
}

// GetCommandPermissions fetches permissions for a specific command for your application in a guild
func (ch *ApplicationHandler) GetCommandPermissions(applicationId, guildId, commandId string) (*discord.GuildApplicationCommandPermissions, error) {
	var permissions *discord.GuildApplicationCommandPermissions

	res, err := ch.rest.Request(fmt.Sprintf(EndpointGetApplicationCommandPermissions, applicationId, guildId, commandId), "GET", bytes.NewBufferString(""), "application/json")

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(res, &permissions)

	if err != nil {
		return nil, err
	}

	return permissions, nil
}

type rawEditCommandPermissions struct {
	Permissions []*discord.ApplicationCommandPermissions `json:"permissions"`
}

// EditCommandPermissions edits command permissions for a specific command for your application in a guild
func (ch *ApplicationHandler) EditCommandPermissions(applicationId, guildId, commandId string, permissions []*discord.ApplicationCommandPermissions) ([]*discord.GuildApplicationCommandPermissions, error) {
	var guildPermissions []*discord.GuildApplicationCommandPermissions
	var err error

	data, err := json.Marshal(&rawEditCommandPermissions{Permissions: permissions})

	if err != nil {
		return nil, err
	}

	res, err := ch.rest.Request(fmt.Sprintf(EndpointEditApplicationCommandPermissions, applicationId, guildId, commandId), "PUT", bytes.NewBuffer(data), "application/json")

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(res, &guildPermissions)

	if err != nil {
		return nil, err
	}

	return guildPermissions, err
}
