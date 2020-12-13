package rest

type GuildHandler struct {
	rest *Client
}

func NewGuildHandler(rest *Client) *GuildHandler {
	return &GuildHandler{rest: rest}
}

// TODO
