package packet

type VoiceHello struct {
	Packet
	Data struct {
		HeartbeatInterval float64 `json:"heartbeat_interval"`
	} `json:"d,omitempty"`
}
