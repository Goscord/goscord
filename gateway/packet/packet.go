package packet

import "encoding/json"

type Packet struct {
	Opcode   int             `json:"op,omitempty"`
	Sequence int64           `json:"s,omitempty"`
	Event    string          `json:"t,omitempty"`
	Data     json.RawMessage `json:"d"`
}

func NewPacket(data []byte) (*Packet, error) {
	var packet Packet

	err := json.Unmarshal(data, &packet)

	if err != nil {
		return nil, err
	}

	return &packet, nil
}
