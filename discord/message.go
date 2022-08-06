package discord

import (
	"time"

	"github.com/Goscord/goscord/discord/embed"
)

type MessageActivityType int

const (
	_ MessageActivityType = iota
	MessageActivityTypeJoin
	MessageActivityTypeSpectate
	MessageActivityTypeListen
	MessageActivityTypeJoinRequest
)

type MessageFlag int

const (
	_ MessageFlag = 1 << iota
	MessageFlagCrossposted
	MessageFlagIsCrosspost
	MessageFlagSuppressEmbeds
	MessageFlagSourceMessageDeleted
	MessageFlagUrgent
	MessageFlagHasThread
	MessageFlagEphemeral
	MessageFlagLoading
	MessageFlagFailedToMentionSomeRolesInThreads
)

type AllowedMentionsType string

const (
	AllowedMentionsRoleMentions AllowedMentionsType = "roles"
	AllowedMentionsUserMentions AllowedMentionsType = "users"
	AllowedMentionsEveryone     AllowedMentionsType = "everyone"
)

type MessageType int

const (
	Default MessageType = iota
	RecipientAdd
	RecipientRemove
	Call
	ChannelNameChange
	ChannelIconChange
	ChannelPinnedMessage
	UserJoin
	GuildBoost
	GuildBoostTier1
	GuildBoostTier2
	GuildBoostTier3
	ChannelFollowAdd
	GuildDiscoveryAdd
	GuildDiscoveryDisqualified
	GuildDiscoveryRequalified
	GuildDiscoveryGracePeriodInitialWarning
	GuildDiscoveryGracePeriodFinalWarning
	ThreadCreated
	Reply
	ChatInputCommand
	ThreadStarterMessage
	GuildInviteReminder
	ContextMenuCommand
	AutoModerationAction
)

type MessageActivity struct {
	Type    string `json:"type"`
	PartyId string `json:"party_id"`
}

type MessageReference struct {
	MessageId      string `json:"message_id,omitempty"`
	ChannelId      string `json:"channel_id,omitempty"`
	GuildId        string `json:"guild_id,omitempty"`
	FailIfNotExist bool   `json:"fail_if_not_exist,omitempty"`
}

type Emoji struct {
	Id            string   `json:"id"`
	Name          string   `json:"name"`
	Roles         []string `json:"roles,omitempty"`
	User          *User    `json:"user,omitempty"`
	RequireColons bool     `json:"require_colons,omitempty"`
	Managed       bool     `json:"managed,omitempty"`
	Animated      bool     `json:"animated,omitempty"`
	Available     bool     `json:"available,omitempty"`
}

type Reaction struct {
	Count int    `json:"count"`
	Me    bool   `json:"me"`
	Emoji *Emoji `json:"emoji"`
}

type Attachment struct {
	Id       string `json:"id"`
	Filename string `json:"filename"`
	Size     int    `json:"size"`
	URL      string `json:"url"`
	Data     []byte `json:"-"`
	ProxyURL string `json:"proxy_url"`
	Height   int    `json:"height"`
	Width    int    `json:"width"`
}

type ChannelMention struct {
	Id      string      `json:"id"`
	GuildId string      `json:"guild_id"`
	Type    ChannelType `json:"type"`
	Name    string      `json:"name"`
}

type AllowedMentions struct {
	Parse        []*AllowedMentions `json:"parse"`
	Roles        []string           `json:"roles"`
	Users        []string           `json:"users"`
	RepliedUsers bool               `json:"replied_users"`
}

type Message struct {
	Id              string           `json:"id"`
	ChannelId       string           `json:"channel_id"`
	GuildId         string           `json:"guild_id,omitempty"`
	Author          *User            `json:"author"`
	Member          *Member          `json:"member"`
	Content         string           `json:"content"`
	Timestamp       time.Time        `json:"timestamp"`
	EditedTimestamp time.Time        `json:"edited_timestamp"`
	Tts             bool             `json:"tts"`
	MentionEveryone bool             `json:"mention_everyone"`
	Mentions        []*User          `json:"mentions"`
	MentionRoles    []string         `json:"mention_roles"`
	MentionChannels []*Channel       `json:"mention_channels,omitempty"`
	Attachments     []*Attachment    `json:"attachments"`
	Embeds          []*embed.Embed   `json:"embeds"`
	Reactions       []*Reaction      `json:"reactions"`
	Nonce           interface{}      `json:"nonce,omitempty"` // integer or string
	Pinned          bool             `json:"pinned"`
	WebhookId       string           `json:"webhook_id,omitempty"`
	Type            int              `json:"type"`
	Activity        *MessageActivity `json:"activity,omitempty"`
	//Application    *Application     `json:"application,omitempty"`
	ApplicationId     string            `json:"application_id,omitempty"`
	MessageReference  *MessageReference `json:"message_reference,omitempty"`
	Flags             int               `json:"flags,omitempty"`
	ReferencedMessage *Message          `json:"referenced_message,omitempty"`
	//Interaction       *Interaction      `json:"interaction,omitempty"`
	Thread *Channel `json:"thread,omitempty"`
	//Components []*MessageComponent `json:"components,omitempty"`
	StickerItems []*StickerItem `json:"sticker_items,omitempty"`
	Stickers     []*Sticker     `json:"stickers,omitempty"`
	Position     int            `json:"position,omitempty"`
}
