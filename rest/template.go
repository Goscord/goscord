package rest

type TemplateHandler struct {
	rest *Client
}

func NewTemplateHandler(rest *Client) *TemplateHandler {
	return &TemplateHandler{rest: rest}
}

// TODO
