package yalis

import (
	"github.com/Seyz123/yalis/user"
	"github.com/Seyz123/yalis/ws"
	ev "github.com/asaskevich/EventBus"
)

type Client struct {
	token string
	bus *ev.EventBus
	session *ws.Session
}

func NewClient(token string) *Client {
	client := new(Client)
	
	client.token = token
	client.bus = ev.New().(*ev.EventBus)
	client.session = ws.NewSession(token, client.Bus())
	
	return client
}

func (c *Client) Login() error {
	return c.session.Login()
}

func (c *Client) Token() string {
	return c.token
}

func (c *Client) Session() *ws.Session {
	return c.session
}

func (c *Client) Bus() *ev.EventBus {
    return c.bus
}

func (c *Client) User() *user.User {
	return c.session.User()
}

func (c *Client) On(ev string, fn interface{}) error {
    return c.bus.Subscribe(ev, fn)
}