package goscord

import (
	"github.com/Goscord/goscord/goscord/gateway"
)

func New(options *gateway.Options) *gateway.Session {
	return gateway.NewSession(options)
}
