package rest

type UserHandler struct {
	rest *Client
}

func NewUserHandler(rest *Client) *UserHandler {
	return &UserHandler{rest: rest}
}

// TODO
