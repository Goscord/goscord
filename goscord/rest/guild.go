package rest

import (
	"fmt"
	"github.com/Goscord/goscord/goscord/discord"
	"github.com/bytedance/sonic"
)

type GuildHandler struct {
	rest *Client
}

func NewGuildHandler(rest *Client) *GuildHandler {
	return &GuildHandler{rest: rest}
}

func (gh *GuildHandler) GetMember(guildId, userId string) (*discord.GuildMember, error) {
	data, err := gh.rest.Request(fmt.Sprintf(EndpointGetGuildMember, guildId, userId), "GET", nil, "application/json")

	if err != nil {
		return nil, err
	}

	var member *discord.GuildMember
	err = sonic.Unmarshal(data, &member)

	if err != nil {
		return nil, err
	}

	return member, nil
}

func (gh *GuildHandler) AddMemberRole(guildId, userId, roleId string) error {
	_, err := gh.rest.Request(fmt.Sprintf(EndpointAddGuildMemberRole, guildId, userId, roleId), "PUT", nil, "application/json")

	if err != nil {
		return err
	}

	return nil
}

// ToDo
