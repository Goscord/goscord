package rest

type InviteHandler struct {
	rest *Client
}

func NewInviteHandler(rest *Client) *InviteHandler {
	return &InviteHandler{rest: rest}
}

// TODO
