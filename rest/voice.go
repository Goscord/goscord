package rest

type VoiceHandler struct {
	rest *Client
}

func NewVoiceHandler(rest *Client) *VoiceHandler {
	return &VoiceHandler{rest: rest}
}

// TODO
