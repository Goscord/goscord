package rest

const (
	BaseUrl                = "https://discord.com/api/v7"
	GatewayUrl             = "wss://gateway.discord.gg/?v=7&encoding=json"
	EndpointGetMessage     = "/channels/%s/messages/%s"
	EndpointCreateMessage  = "/channels/%s/messages"
	EndpointEditMessage    = "/channels/%s/messages/%s"
	EndpointGetChannel     = "/channels/%s"
	EndpointGetGuildMember = "/guilds/%s/members/%s"
)
