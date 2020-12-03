package message

import (
	"encoding/json"
	"fmt"
	"github.com/Seyz123/yalis/channel"
	"github.com/Seyz123/yalis/channel/message/embed"
	"github.com/Seyz123/yalis/guild"
	"github.com/Seyz123/yalis/rest"
	"github.com/Seyz123/yalis/user"
	"time"
)

type Message struct {
	Rest            *rest.Client  `json:"-"`
	Id              string        `json:"id"`
	ChannelId       string        `json:"channel_id"`
	GuildId         string        `json:"guild_id,omitempty"`
	Author          *user.User    `json:"author"`
	Member          *guild.Member `json:"member"`
	Content         string        `json:"content"`
	Timestamp       *time.Time    `json:"timestamp"`
	EditedTimestamp *time.Time    `json:"edited_timestamp"`
	Tts             bool          `json:"tts"`
	MentionEveryone bool          `json:"mention_everyone"`
	//MentionRoles []*guild.Role `json:"mention_roles"`
	MentionChannels []*channel.Channel `json:"mention_channels"`
	Attachments     []*Attachment      `json:"attachments"`
	Embeds          []*embed.Embed     `json:"embeds"`
	//Reactions []*message.Reaction `json:"reactions"`

	// ToDo : Add other properties
}

func (m *Message) Reply(content interface{}) (*Message, error) {
	b, err := FormatMessage(content)

	if err != nil {
		return nil, err
	}

	res, err := m.Rest.Request(fmt.Sprintf(rest.EndpointCreateMessage, m.ChannelId), "POST", b)

	if err != nil {
		return nil, err
	}

	message := new(Message)

	err = json.Unmarshal(res, message)

	if err != nil {
		return nil, err
	}

	return message, nil
}
