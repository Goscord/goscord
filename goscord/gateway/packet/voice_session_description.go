package packet

type VoiceSessionDescription struct {
	Packet
	Data struct {
		Mode      string   `json:"mode"`
		SecretKey [32]byte `json:"secret_key"`
	} `json:"d,omitempty"`
}
