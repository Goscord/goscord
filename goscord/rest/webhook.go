package rest

type WebhookHandler struct {
	rest *Client
}

func NewWebhookHandler(rest *Client) *WebhookHandler {
	return &WebhookHandler{rest: rest}
}

// TODO
