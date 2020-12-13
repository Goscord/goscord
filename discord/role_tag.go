package discord

type RoleTag struct {
	BotId             string      `json:"bot_id"`
	IntegrationId     string      `json:"integration_id"`
	PremiumSubscriber interface{} `json:"premium_subscriber"`
}
