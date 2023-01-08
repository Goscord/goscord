package rest

type EmojiHandler struct {
	rest *Client
}

func NewEmojiHandler(rest *Client) *EmojiHandler {
	return &EmojiHandler{rest: rest}
}

// TODO
