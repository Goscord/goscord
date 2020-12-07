package goscord

import (
	"github.com/Goscord/goscord/gateway"
)

func New(token string) *gateway.Session {
	return gateway.NewSession(token)
}