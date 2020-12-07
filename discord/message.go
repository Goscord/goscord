package discord

import (
	"encoding/json"
	"fmt"
	"github.com/Goscord/goscord/discord/embed"
	"github.com/Goscord/goscord/rest"
	"github.com/Goscord/goscord/user"
	"time"
)

type Message struct {
	Rest            *rest.Client `json:"-"`
	Id              string       `json:"id"`
	ChannelId       string       `json:"channel_id"`
	GuildId         string       `json:"guild_id,omitempty"`
	Author          *user.User   `json:"author"`
	Member          *Member      `json:"member"`
	Content         string       `json:"content"`
	Timestamp       *time.Time   `json:"timestamp"`
	EditedTimestamp *time.Time   `json:"edited_timestamp"`
	Tts             bool         `json:"tts"`
	MentionEveryone bool         `json:"mention_everyone"`
	//MentionRoles []*guild.Role `json:"mention_roles"`
	MentionChannels []*Channel     `json:"mention_channels"`
	Attachments     []*Attachment  `json:"attachments"`
	Embeds          []*embed.Embed `json:"embeds"`
	//Reactions []*message.Reaction `json:"reactions"`

	// ToDo : Add other properties
}

func (m *Message) Reply(content interface{}) (*Message, error) {
	switch content.(type) {
	case string:
		content = map[string]string{"content": fmt.Sprintf("<@%s>, %s", m.Author.Id, content.(string))}

	case *embed.Embed:
		content = &embed.MessageEmbed{Content: fmt.Sprintf("<@%s>, ", m.Author.Id), Embed: content.(*embed.Embed)}
	}

	b, err := json.Marshal(content)

	if err != nil {
		return nil, err
	}

	res, err := m.Rest.Request(fmt.Sprintf(rest.EndpointCreateMessage, m.ChannelId), "POST", b)

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

func (m *Message) Channel() (*Channel, error) {
	b, err := m.Rest.Request(fmt.Sprintf(rest.EndpointGetChannel, m.ChannelId), "GET", nil)

	if err != nil {
		return nil, err
	}

	ch, err := NewChannel(m.Rest, b)

	if err != nil {
		return nil, err
	}

	return ch, nil
}

func NewMessage(rest *rest.Client, data []byte) (*Message, error) {
	msg := new(Message)

	err := json.Unmarshal(data, msg)

	if err != nil {
		return nil, err
	}

	msg.Rest = rest

	return msg, nil
}