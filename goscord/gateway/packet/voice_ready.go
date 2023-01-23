package packet

type VoiceReady struct {
	Packet
	Data struct {
		SSRC  int      `json:"ssrc"`
		IP    string   `json:"ip"`
		Port  int      `json:"port"`
		Modes []string `json:"modes"`
	} `json:"d,omitempty"`
}
