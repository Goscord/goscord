package rest

const (
	BaseUrl    = "https://discord.com/api/v7"
	GatewayUrl = "wss://gateway.discord.gg/?v=7&encoding=json"

	// Audit Log
	EndpointGetGuildsAuditLog = "/guilds/%s/audit-logs"

	// Auto Moderation
	EndpointListAutoModerationRulesForGuild = "/guilds/%s/auto-moderation/rules"
	EndpointGetAutoModerationRule           = "/guilds/%s/auto-moderation/rules/%s"
	EndpointCreateAutoModerationRule        = "/guilds/%s/auto-moderation/rules"
	EndpointModifyAutoModerationRule        = "/guilds/%s/auto-moderation/rules/%s"
	EndpointDeleteAutoModerationRule        = "/guilds/%s/auto-moderation/rules/%s"

	// Channel
	EndpointGetChannel                       = "/channels/%s"
	EndpointModifyChannel                    = "/channels/%s"
	EndpointDeleteChannel                    = "/channels/%s"
	EndpointGetChannelMessages               = "/channels/%s/messages"
	EndpointGetChannelMessage                = "/channels/%s/messages/%s"
	EndpointCreateMessage                    = "/channels/%s/messages"
	EndpointCrosspostMessage                 = "/channels/%s/messages/%s/crosspost"
	EndpointOwnReaction                      = "/channels/%s/messages/%s/reactions/%s/@me"
	EndpointDeleteUserReaction               = "/channels/%s/messages/%s/reactions/%s/%s"
	EndpointGetReactions                     = "/channels/%s/messages/%s/reactions/%s"
	EndpointDeleteAllReactions               = "/channels/%s/messages/%s/reactions"
	EndpointDeleteAllReactionsForEmoji       = "/channels/%s/messages/%s/reactions/%s"
	EndpointEditMessage                      = "/channels/%s/messages/%s"
	EndpointDeleteMessage                    = "/channels/%s/messages/%s"
	EndpointBulkDeleteMessages               = "/channels/%s/messages/bulk-delete"
	EndpointEditChannelPermissions           = "/channels/%s/permissions/%s"
	EndpointGetChannelInvites                = "/channels/%s/invites"
	EndpointCreateChannelInvite              = "/channels/%s/invites"
	EndpointDeleteChannelPermission          = "/channels/%s/permissions/%s"
	EndpointFollowNewsChannel                = "/channels/%s/followers"
	EndpointTriggerTypingIndicator           = "/channels/%s/typing"
	EndpointGetPinnedMessages                = "/channels/%s/pins"
	EndpointPinMessage                       = "/channels/%s/pins/%s"
	EndpointUnpinMessage                     = "/channels/%s/pins/%s"
	EndpointGroupDMAddRecipient              = "/channels/%s/recipients/%s"
	EndpointGroupDMRemoveRecipient           = "/channels/%s/recipients/%s"
	EndpointStartThreadFromMessage           = "/channels/%s/messages/%s/threads"
	EndpointStartThreadWithoutMessage        = "/channels/%s/threads"
	EndpointStartThreadInForumChannel        = "/channels/%s/threads"
	EndpointJoinThread                       = "/channels/%s/thread-members/@me"
	EndpointAddThreadMember                  = "/channels/%s/thread-members/%s"
	EndpointLeaveThread                      = "/channels/%s/thread-members/@me"
	EndpointRemoveThreadMember               = "/channels/%s/thread-members/%s"
	EndpointGetThreadMember                  = "/channels/%s/thread-members/%s"
	EndpointListThreadMembers                = "/channels/%s/thread-members"
	EndpointListPublicArchivedThreads        = "/channels/%s/threads/archived/public"
	EndpointListPrivateArchivedThreads       = "/channels/%s/threads/archived/private"
	EndpointListJoinedPrivateArchivedThreads = "/channels/%s/users/@me/threads/archived/private"

	// Emoji
	EndpointListGuildEmojis  = "/guilds/%s/emojis"
	EndpointGetGuildEmoji    = "/guilds/%s/emojis/%s"
	EndpointCreateGuildEmoji = "/guilds/%s/emojis"
	EndpointModifyGuildEmoji = "/guilds/%s/emojis/%s"
	EndpointDeleteGuildEmoji = "/guilds/%s/emojis/%s"

	// Guild
	EndpointGetGuildMember = "/guilds/%s/members/%s"
)
