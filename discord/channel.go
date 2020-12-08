package discord

import (
	"encoding/json"
	"fmt"
	"github.com/Goscord/goscord/discord/embed"
	"github.com/Goscord/goscord/rest"
	"time"
)

type Channel struct {
	Rest     *rest.Client `json:"-"`
	Id       string       `json:"id"`
	Type     int          `json:"type"`
	GuildId  string       `json:"guild_id"`
	Position int          `json:"position"`
	//PermissionOverwrites []PermissionOverwrite `json:"permission_overwrites"`
	Name             string     `json:"name"`
	Topic            string     `json:"topic"`
	Nsfw             bool       `json:"nsfw"`
	LastMessageId    string     `json:"last_message_id"`
	Bitrate          int        `json:"bitrate"`
	UserLimit        int        `json:"user_limit"`
	RateLimitPerUser int        `json:"rate_limit_per_user"`
	Recipients       []User     `json:"recipients"`
	Icon             string     `json:"icon"`
	OwnerId          string     `json:"owner_id"`
	ApplicationId    string     `json:"application_id"`
	ParentId         string     `json:"parent_id"`
	LastPinTimestamp *time.Time `json:"last_pin_timestamp"`
}

func (ch *Channel) Send(content interface{}) (*Message, error) {
	switch content.(type) {
	case string:
		content = map[string]string{"content": content.(string)}

	case *embed.Embed:
		content = &embed.MessageEmbed{Embed: content.(*embed.Embed)}
	}

	b, err := json.Marshal(content)

	if err != nil {
		return nil, err
	}

	res, err := ch.Rest.Request(fmt.Sprintf(rest.EndpointCreateMessage, ch.Id), "POST", b)

	if err != nil {
		return nil, err
	}

	msg := new(Message)

	err = json.Unmarshal(res, msg)

	if err != nil {
		return nil, err
	}

	return msg, nil
}

func NewChannel(rest *rest.Client, data []byte) (*Channel, error) {
	channel := new(Channel)

	err := json.Unmarshal(data, channel)

	if err != nil {
		return nil, err
	}

	channel.Rest = rest

	return channel, nil
}
