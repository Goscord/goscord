package discord

type BitwisePermissionFlag int64

const (
	BitwisePermissionFlagCreateInstantInvite BitwisePermissionFlag = 1 << iota
	BitwisePermissionFlagKickMembers
	BitwisePermissionFlagBanMembers
	BitwisePermissionFlagAdministrator
	BitwisePermissionFlagManageChannels
	BitwisePermissionFlagManageGuild
	BitwisePermissionFlagAddReactions
	BitwisePermissionFlagViewAuditLog
	BitwisePermissionFlagPrioritySpeaker
	BitwisePermissionFlagStream
	BitwisePermissionFlagViewChannel
	BitwisePermissionFlagSendMessages
	BitwisePermissionFlagSendTTSMessages
	BitwisePermissionFlagManageMessages
	BitwisePermissionFlagEmbedLinks
	BitwisePermissionFlagAttachFiles
	BitwisePermissionFlagReadMessageHistory
	BitwisePermissionFlagMentionEveryone
	BitwisePermissionFlagUseExternalEmojis
	BitwisePermissionFlagViewGuildInsights
	BitwisePermissionFlagConnect
	BitwisePermissionFlagSpeak
	BitwisePermissionFlagMuteMembers
	BitwisePermissionFlagDeafenMembers
	BitwisePermissionFlagMoveMembers
	BitwisePermissionFlagUseVAD
	BitwisePermissionFlagChangeNickname
	BitwisePermissionFlagManageNicknames
	BitwisePermissionFlagManageRoles
	BitwisePermissionFlagManageWebhooks
	BitwisePermissionFlagManageEmojisAndStickers
	BitwisePermissionFlagManageUseApplicationCommands
	BitwisePermissionFlagManageRequestToSpeak
	BitwisePermissionFlagManageManageEvents
	BitwisePermissionFlagManageManageThreads
	BitwisePermissionFlagManageCreatePublicThreads
	BitwisePermissionFlagManageCreatePrivateThreads
	BitwisePermissionFlagManageUseExternalStickers
	BitwisePermissionFlagManageSendMessagesInThreads
	BitwisePermissionFlagManageUseEmbeddedActivities
	BitwisePermissionFlagManageModerateMembers
)

func (b BitwisePermissionFlag) Has(flag BitwisePermissionFlag) bool {
	return b&flag == flag
}

type Locale string

const (
	EnglishUS    Locale = "en-US"
	EnglishGB    Locale = "en-GB"
	Bulgarian    Locale = "bg"
	ChineseCN    Locale = "zh-CN"
	ChineseTW    Locale = "zh-TW"
	Croatian     Locale = "hr"
	Czech        Locale = "cs"
	Danish       Locale = "da"
	Dutch        Locale = "nl"
	Finnish      Locale = "fi"
	French       Locale = "fr"
	German       Locale = "de"
	Greek        Locale = "el"
	Hindi        Locale = "hi"
	Hungarian    Locale = "hu"
	Italian      Locale = "it"
	Japanese     Locale = "ja"
	Korean       Locale = "ko"
	Lithuanian   Locale = "lt"
	Norwegian    Locale = "no"
	Polish       Locale = "pl"
	PortugueseBR Locale = "pt-BR"
	Romanian     Locale = "ro"
	Russian      Locale = "ru"
	SpanishES    Locale = "es-ES"
	Swedish      Locale = "sv-SE"
	Thai         Locale = "th"
	Turkish      Locale = "tr"
	Ukrainian    Locale = "uk"
	Vietnamese   Locale = "vi"
)

type StatusType string

const (
	StatusTypeOnline       StatusType = "online"
	StatusTypeIdle         StatusType = "idle"
	StatusTypeDoNotDisturb StatusType = "dnd"
	StatusTypeOffline      StatusType = "offline"
)

type ActivityType int

const (
	ActivityPlaying ActivityType = iota
	ActivityStreaming
	ActivityListening
	ActivityWatching
	ActivityCustom
	ActivityCompeting
)

type ActivityFlag int

const (
	ActivityFlagInstance ActivityFlag = 1 << iota
	ActivityFlagJoin
	ActivityFlagSpectate
	ActivityFlagJoinRequest
	ActivityFlagSync
	ActivityFlagPlay
	ActivityFlagPartyPrivacyFriends
	ActivityFlagPartyPrivacyVoiceChannel
	ActivityFlagEmbedded
)

type ActivityTimestamps struct {
	Start int `json:"start,omitempty"` // Unix timestamp in milliseconds
	End   int `json:"end,omitempty"`   // Unix timestamp in milliseconds
}

type ActivityEmoji struct {
	Name     string `json:"name"`
	Id       string `json:"id,omitempty"`
	Animated bool   `json:"animated,omitempty"`
}

type ActivityParty struct {
	Id   string `json:"id,omitempty"`
	Size [2]int `json:"size,omitempty"` // array of 2 ints
}

type ActivityAssets struct {
	LargeImage string `json:"large_image,omitempty"`
	LargeText  string `json:"large_text,omitempty"`
	SmallImage string `json:"small_image,omitempty"`
	SmallText  string `json:"small_text,omitempty"`
}

// ToDo : Add ActivityAssets funcs

type ActivitySecrets struct {
	Join     string `json:"join,omitempty"`
	Spectate string `json:"spectate,omitempty"`
	Match    string `json:"match,omitempty"`
}

type ActivityButton struct {
	Label string `json:"label"`
	URL   string `json:"url"`
}

type Activity struct {
	Name          string              `json:"name"`
	Type          ActivityType        `json:"type"`
	StreamURL     string              `json:"url,omitempty"`
	CreatedAt     int                 `json:"created_at"`
	Timestamps    *ActivityTimestamps `json:"timestamps,omitempty"`
	ApplicationId string              `json:"application_id,omitempty"`
	Details       string              `json:"details,omitempty"`
	State         string              `json:"state,omitempty"`
	Emojis        []*ActivityEmoji    `json:"emojis,omitempty"`
	Party         *ActivityParty      `json:"party,omitempty"`
	Assets        *ActivityAssets     `json:"assets,omitempty"`
	Secrets       *ActivitySecrets    `json:"secrets,omitempty"`
	Instance      bool                `json:"instance,omitempty"`
	Flags         ActivityFlag        `json:"flags,omitempty"`
	Buttons       []*ActivityButton   `json:"buttons,omitempty"`
}
