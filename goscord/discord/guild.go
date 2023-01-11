package discord

import (
	"time"

	"github.com/goccy/go-json"
)

type RoleTag struct {
	BotId             string          `json:"bot_id,omitempty"`
	IntegrationId     string          `json:"integration_id,omitempty"`
	PremiumSubscriber json.RawMessage `json:"premium_subscriber,omitempty"` // null, whether this is the guild's premium subscriber role
}

type Role struct {
	Id           string                `json:"id"`
	Name         string                `json:"name"`
	Color        int                   `json:"color"`
	Hoist        bool                  `json:"hoist"`
	Icon         string                `json:"icon,omitempty"`
	UnicodeEmoji string                `json:"unicode_emoji,omitempty"`
	Position     int                   `json:"position"`
	Permissions  BitwisePermissionFlag `json:"permissions,string"`
	Managed      bool                  `json:"managed"`
	Mentionable  bool                  `json:"mentionable"`
	Tags         *RoleTag              `json:"tags,omitempty"`
}

// A function to turn a role into a string. The Role ID is the contents of the string.
func (role *Role) String() string {
	return role.Id
}

type MessageNotificationLevel int

const (
	MessageNotificationLevelAllMessages MessageNotificationLevel = iota
	MessageNotificationLevelOnlyMentions
)

type ExplicitContentFilterLevel int

const (
	ExplicitContentFilterLevelDisabled ExplicitContentFilterLevel = iota
	ExplicitContentFilterLevelMembersWithoutRoles
	ExplicitContentFilterLevelAllMembers
)

type MfaLevel int

const (
	MfaLevelNone MfaLevel = iota
	MfaLevelElevated
)

type VerificationLevel int

const (
	VerificationLevelNone VerificationLevel = iota
	VerificationLevelLow
	VerificationLevelMedium
	VerificationLevelHigh
	VerificationLevelVeryHigh
)

type GuildNsfwLevel int

const (
	GuildNsfwLevelDefault GuildNsfwLevel = iota
	GuildNsfwLevelExplicit
	GuildNsfwLevelSafe
	GuildNsfwLevelAgeRestricted
)

type PremiumTier int

const (
	PremiumTierNone PremiumTier = iota
	PremiumTierTier1
	PremiumTierTier2
	PremiumTierTier3
)

type SystemChannelFlag int

const (
	SystemChannelFlagSupressJoinNotifications SystemChannelFlag = 1 << iota
	SystemChannelFlagSupressPremiumSubscriptions
	SystemChannelFlagSupressGuildReminderNotifications
	SystemChannelFlagSupressJoinNotificationReplies
)

type GuildFeature string

const (
	GuildFeatureAnimatedBanner                GuildFeature = "ANIMATED_BANNER"
	GuildFeatureAnimatedIcon                  GuildFeature = "ANIMATED_ICON"
	GuildFeatureAutoModeration                GuildFeature = "AUTO_MODERATION"
	GuildFeatureBanner                        GuildFeature = "BANNER"
	GuildFeatureCommunity                     GuildFeature = "COMMUNITY"
	GuildFeatureDiscoverable                  GuildFeature = "DISCOVERABLE"
	GuildFeatureFeaturable                    GuildFeature = "FEATURABLE"
	GuildFeatureInviteSplash                  GuildFeature = "INVITE_SPLASH"
	GuildFeatureMemberVerificationGateEnabled GuildFeature = "MEMBER_VERIFICATION_GATE_ENABLED"
	GuildFeatureMonetizationEnabled           GuildFeature = "MONETIZATION_ENABLED"
	GuildFeatureMoreStickers                  GuildFeature = "MORE_STICKERS"
	GuildFeatureNews                          GuildFeature = "NEWS"
	GuildFeaturePartnered                     GuildFeature = "PARTNERED"
	GuildFeaturePreviewEnabled                GuildFeature = "PREVIEW_ENABLED"
	GuildFeaturePrivateThreads                GuildFeature = "PRIVATE_THREADS"
	GuildFeatureRoleIcons                     GuildFeature = "ROLE_ICONS"
	GuildFeatureTicketedEventsEnabled         GuildFeature = "TICKETED_EVENTS_ENABLED"
	GuildFeatureVanityURL                     GuildFeature = "VANITY_URL"
	GuildFeatureVerified                      GuildFeature = "VERIFIED"
	GuildFeatureVipRegions                    GuildFeature = "VIP_REGIONS"
	GuildFeatureWelcomeScreenEnabled          GuildFeature = "WELCOME_SCREEN_ENABLED"
)

type IntegrationExpireBehavior int

const (
	IntegrationExpireBehaviorRemoveRole IntegrationExpireBehavior = iota
	IntegrationExpireBehaviorKick
)

type UnavailableGuild Guild

type GuildPreview Guild

type GuildWidgetSettings struct {
	Enabled   bool   `json:"enabled"`
	ChannelId string `json:"channel_id"`
}

type GuildWidget struct {
	Id            string         `json:"id"`
	Name          string         `json:"name"`
	InstantInvite string         `json:"instant_invite"`
	Channels      []*Channel     `json:"channels"`
	Members       []*GuildMember `json:"members"`
	PresenceCount int            `json:"presence_count"`
}

type GuildMember struct {
	User                       *User                 `json:"user"`
	Nick                       string                `json:"nick,omitempty"`
	Roles                      []string              `json:"roles"`
	JoinedAt                   *time.Time            `json:"joined_at"`
	PremiumSince               *time.Time            `json:"premium_since,omitempty"`
	Deaf                       bool                  `json:"deaf"`
	Mute                       bool                  `json:"mute"`
	Pending                    bool                  `json:"pending"`
	Permissions                BitwisePermissionFlag `json:"permissions,string,omitempty"`
	CommunicationDisabledUntil *time.Time            `json:"communication_disabled_until"`
	GuildId                    string                `json:"guild_id"`
}

type Integration struct {
	Id                string                    `json:"id"`
	Name              string                    `json:"name"`
	Type              string                    `json:"type"` // twitch, youtube or discord
	Enabled           bool                      `json:"enabled,omitempty"`
	Syncing           bool                      `json:"syncing,omitempty"`
	RoleId            string                    `json:"role_id,omitempty"`
	EnableEmoticons   bool                      `json:"enable_emoticons,omitempty"`
	ExpireBehavior    IntegrationExpireBehavior `json:"expire_behavior,omitempty"`
	ExpireGracePeriod int                       `json:"expire_grace_period,omitempty"`
	User              *User                     `json:"user,omitempty"`
	Account           *IntegrationAccount       `json:"account"`
	SyncedAt          *time.Time                `json:"synced_at,omitempty"`
	SubscriberCount   int                       `json:"subscriber_count,omitempty"`
	Revoked           bool                      `json:"revoked,omitempty"`
	Applications      []*IntegrationApplication `json:"applications,omitempty"`
}

type IntegrationAccount struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type IntegrationApplication struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Icon        string `json:"icon"`
	Description string `json:"description"`
	Bot         *User  `json:"bot,omitempty"`
}

type Ban struct {
	Reason string `json:"reason"`
	User   *User  `json:"user"`
}

type WelcomeScreen struct {
	Description     string                  `json:"description"`
	WelcomeChannels []*WelcomeScreenChannel `json:"welcome_channels"`
}

type WelcomeScreenChannel struct {
	ChannelId   string `json:"channel_id"`
	Description string `json:"description"`
	EmojiId     string `json:"emoji_id"`
	EmojiName   string `json:"emoji_name"`
}

type PresenceUpdate struct {
	User         *User         `json:"user"`
	GuildId      string        `json:"guild_id"`
	Status       string        `json:"status"`
	Activities   []*Activity   `json:"activities"`
	ClientStatus *ClientStatus `json:"client_status"`
}

type ClientStatus struct {
	Desktop string `json:"deskop,omitempty"` // windows, linux, mac
	Mobile  string `json:"mobile,omitempty"` // ios, android
	Web     string `json:"web,omitempty"`    // browser, bot_account
}

type Guild struct {
	Id                          string                     `json:"id"`
	Name                        string                     `json:"name"`
	Icon                        string                     `json:"icon,omitempty"`
	IconHash                    string                     `json:"icon_hash,omitempty"`
	Splash                      string                     `json:"splash,omitempty"`
	DiscoverySplash             string                     `json:"discovery_splash"`
	Owner                       bool                       `json:"owner,omitempty"`
	OwnerId                     string                     `json:"owner_id"`
	Permissions                 BitwisePermissionFlag      `json:"permissions,string,omitempty"`
	Region                      string                     `json:"region,omitempty"`
	AfkChannelId                string                     `json:"afk_channel_id"`
	AfkTimeout                  int                        `json:"afk_timeout"`
	WidgetEnabled               bool                       `json:"widget_enabled,omitempty"`
	WidgetChannelId             string                     `json:"widget_channel_id,omitempty"`
	VerificationLevel           VerificationLevel          `json:"verification_level"`
	DefaultMessageNotifications MessageNotificationLevel   `json:"default_message_notifications"`
	ExplicitContentFilter       ExplicitContentFilterLevel `json:"explicit_content_filter"`
	Roles                       []*Role                    `json:"roles"`
	Emojis                      []*Emoji                   `json:"emojis"`
	Features                    []GuildFeature             `json:"features"`
	MfaLevel                    MfaLevel                   `json:"mfa_level"`
	ApplicationId               string                     `json:"application_id"`
	SystemChannelId             string                     `json:"system_channel_id"`
	SystemChannelFlags          SystemChannelFlag          `json:"system_channel_flags"`
	RuleChannelId               string                     `json:"rules_channel_id"`
	MaxPresences                int                        `json:"max_presences,omitempty"`
	MaxMembers                  int                        `json:"max_members,omitempty"`
	VanityUrlCode               string                     `json:"vanity_url_code"`
	Description                 string                     `json:"description"`
	Banner                      string                     `json:"banner"`
	PremiumTier                 PremiumTier                `json:"premium_tier"`
	PremiumSubscriptionCount    int                        `json:"premium_subscription_count,omitempty"`
	PreferredLocale             Locale                     `json:"preferred_locale"` // Reference: https://discordapp.com/developers/docs/resources/guild#guild-object-preferred-locale
	PublicUpdatesChannelId      string                     `json:"public_updates_channel_id"`
	MaxVideoChannelUsers        int                        `json:"max_video_channel_users,omitempty"`
	ApproximateMemberCount      int                        `json:"approximate_member_count,omitempty"`
	ApproximatePresenceCount    int                        `json:"approximate_presence_count,omitempty"`
	WelcomeScreen               *WelcomeScreen             `json:"welcome_screen,omitempty"`
	NsfwLevel                   GuildNsfwLevel             `json:"nsfw_level"`
	Stickers                    []*Sticker                 `json:"stickers,omitempty"`
	PremiumProgressBarEnabled   bool                       `json:"premium_progress_bar_enabled"`

	// GUILD_CREATE event specific fields
	JoinedAt       *time.Time        `json:"joined_at"`
	Large          bool              `json:"large"`
	Unavailable    bool              `json:"unavailable,omitempty"`
	MemberCount    int               `json:"member_count"`
	VoiceStates    []*VoiceState     `json:"voice_states"`
	Members        []*GuildMember    `json:"members"`
	Channels       []*Channel        `json:"channels"`
	Threads        []*Channel        `json:"threads"`
	Presences      []*PresenceUpdate `json:"presences"`
	StageInstances []*Channel        `json:"stage_instances"` // ToDo : Change to StageInstance
	//GuildScheduledEvents []*GuildScheduledEvent  `json:"guild_scheduled_events"`
}

// GuildEmojisUpdateEventFields is the fields for the GUILD_EMOJIS_UPDATE event
type GuildEmojisUpdateEventFields struct {
	GuildId string   `json:"guild_id"`
	Emojis  []*Emoji `json:"emojis"`
}

type GuildMemberRemoveEventFields struct {
	GuildId string `json:"guild_id"`
	User    *User  `json:"user"`
}

// GuildStickersUpdateEventFields is the fields for the GUILD_STICKERS_UPDATE event
type GuildStickersUpdateEventFields struct {
	GuildId  string     `json:"guild_id"`
	Stickers []*Sticker `json:"stickers"`
}

// GuildIntegrationsUpdateEventFields is the fields for the GUILD_INTEGRATIONS_UPDATE event
type GuildIntegrationsUpdateEventFields struct {
	GuildId string `json:"guild_id"`
}

// GuildMembersChunkEventFields is the fields for the GUILD_MEMBERS_CHUNK event
type GuildMembersChunkEventFields struct {
	GuildId    string            `json:"guild_id"`
	Members    []*GuildMember    `json:"members"`
	ChunkIndex int               `json:"chunk_index"`
	ChunkCount int               `json:"chunk_count"`
	NotFound   []string          `json:"not_found,omitempty"`
	Presences  []*PresenceUpdate `json:"presences,omitempty"`
	Nonce      string            `json:"nonce,omitempty"`
}

// GuildRoleCreateEventFields is the fields for the GUILD_ROLE_CREATE event
type GuildRoleCreateEventFields struct {
	GuildId string `json:"guild_id"`
	Role    *Role  `json:"role"`
}

// GuildRoleUpdateEventFields is the fields for the GUILD_ROLE_UPDATE event
type GuildRoleUpdateEventFields struct {
	GuildId string `json:"guild_id"`
	Role    *Role  `json:"role"`
}

// GuildRoleDeleteEventFields is the fields for the GUILD_ROLE_DELETE event
type GuildRoleDeleteEventFields struct {
	GuildId string `json:"guild_id"`
	RoleId  string `json:"role_id"`
}
