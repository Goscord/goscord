package gateway

const (
	IntentGuilds int = 1 << iota
	IntentGuildMembers
	IntentGuildBans
	IntentGuildEmojisAndStickers
	IntentGuildIntegrations
	IntentGuildWebhooks
	IntentGuildInvites
	IntentGuildVoiceStates
	IntentGuildPresences
	IntentGuildMessages
	IntentGuildMessageReactions
	IntentGuildMessageTyping
	IntentDirectMessages
	IntentDirectMessageReactions
	IntentDirectMessageTyping
	IntentGuildScheduledEvents
	IntentAutoModerationConfiguration
	IntentAutoModerationExecution
)
