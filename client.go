package yalis

import (
	"github.com/Seyz123/yalis/rest"
	"github.com/Seyz123/yalis/ws"
)

type Client struct {
	token string
	rest *rest.Client
	ws *ws.Session
}

func NewClient(token string) (*Client) {
	client := &Client{}

	client.token = token
	client.rest = rest.NewClient(token)
	client.ws = ws.NewSession(token)

	return client
}

func (c *Client) Login() error {
	return c.WebSocket().Login()
}

func (c *Client) Token() string {
	return c.token
}

func (c *Client) RestClient() *rest.Client {
	return c.rest
}

func (c *Client) WebSocket() *ws.Session {
	return c.ws
}