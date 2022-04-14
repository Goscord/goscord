package rest

import (
	"encoding/json"
	"fmt"
	"github.com/Goscord/goscord/discord"
)

type GuildHandler struct {
	rest *Client
}

func NewGuildHandler(rest *Client) *GuildHandler {
	return &GuildHandler{rest: rest}
}

func (gh *GuildHandler) GetMember(guildID, userID string) (*discord.Member, error) {
	data, err := gh.rest.Request(fmt.Sprintf(EndpointGetGuildMember, userID), "GET", nil)

	if err != nil {
		return nil, err
	}

	var member discord.Member
	err = json.Unmarshal(data, &member)

	if err != nil {
		return nil, err
	}

	return &member, nil
}

// TODO
