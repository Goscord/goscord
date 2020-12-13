package discord

import (
	"github.com/Goscord/goscord/discord/embed"
	"time"
)

type Message struct {
	Id              string         `json:"id"`
	ChannelId       string         `json:"channel_id"`
	GuildId         string         `json:"guild_id,omitempty"`
	Author          *User          `json:"author"`
	Member          *Member        `json:"member"`
	Content         string         `json:"content"`
	Timestamp       *time.Time     `json:"timestamp"`
	EditedTimestamp *time.Time     `json:"edited_timestamp"`
	Tts             bool           `json:"tts"`
	MentionEveryone bool           `json:"mention_everyone"`
	MentionRoles    []*Role        `json:"mention_roles"`
	MentionChannels []*Channel     `json:"mention_channels"`
	Attachments     []*Attachment  `json:"attachments"`
	Embeds          []*embed.Embed `json:"embeds"`
	//Reactions []*message.Reaction `json:"reactions"`

	// ToDo : Add other properties
}
