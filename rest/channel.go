package rest

type ChannelHandler struct {
	rest *Client
}

func NewChannelHandler(rest *Client) *ChannelHandler {
	return &ChannelHandler{rest: rest}
}

// TODO
