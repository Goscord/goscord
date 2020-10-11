package packet

import "encoding/json"

type Hello struct {
	*Packet
	Data struct {
		HeartbeatInterval int `json:"heartbeat_interval,omitempty"`
	} `json:"d,omitempty"`
}

func NewHello(data []byte) (*Hello, error) {
	var packet Hello

	err := json.Unmarshal(data, &packet)

	if err != nil {
		return nil, err
	}

	return &packet, nil
}