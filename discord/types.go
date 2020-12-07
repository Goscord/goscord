package discord

const (
	TypeGuildText = iota
	TypeDm
	TypeGuildVoice
	TypeGroupDm
	TypeGuildCategory
	TypeGuildNews
	TypeGuildStore
)

const (
	Default = iota
	RecipientAdd
	RecipientRemove
	Call
	ChannelNameChange
	ChannelIconChange
	ChannelPinnedMessage
	GuildMemberJoin
	UserPremiumGuildSubscription
	UserPremiumGuildSubscriptionTier1
	UserPremiumGuildSubscriptionTier2
	UserPremiumGuildSubscriptionTier3
	ChannelFollowAdd
	GuildDiscoveryDisqualified
	GuildDiscoveryRequalified
	Reply
)
