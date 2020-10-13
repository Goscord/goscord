package yalis

import (
    "github.com/Seyz123/yalis/rest"
    "github.com/Seyz123/yalis/ws"
	ev "github.com/asaskevich/EventBus"
)

type Client struct {
	token string
	bus *ev.EventBus
	rest *rest.Client
	ws *ws.Session
}

func NewClient(token string) (*Client) {
	client := new(Client)
	
	client.token = token
	client.bus = ev.New().(*ev.EventBus)
	client.rest = rest.NewClient(token)
	client.ws = ws.NewSession(token, client.Bus())
	
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

func (c *Client) Bus() *ev.EventBus {
    return c.bus
}

func (c *Client) On(ev string, fn interface{}) error {
    return c.bus.Subscribe(ev, fn)
}