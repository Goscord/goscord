package packet

import (
	"github.com/bytedance/sonic"
	"time"
)

type Hello struct {
	*Packet
	Data struct {
		HeartbeatInterval time.Duration `json:"heartbeat_interval"`
	} `json:"d"`
}

func NewHello(data []byte) (*Hello, error) {
	var packet Hello

	err := sonic.Unmarshal(data, &packet)

	if err != nil {
		return nil, err
	}

	return &packet, nil
}
