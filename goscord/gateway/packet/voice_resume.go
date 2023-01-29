package packet

type VoiceResume struct {
	Packet
	Data struct {
		ServerId  string `json:"server_id"`
		SessionId string `json:"session_id"`
		Token     string `json:"token"`
	} `json:"d,omitempty"`
}

func NewVoiceResume(serverId, sessionId, token string) *VoiceResume {
	resume := new(VoiceResume)

	resume.Opcode = OpVoiceResume

	resume.Data.ServerId = serverId
	resume.Data.SessionId = sessionId
	resume.Data.Token = token

	return resume
}
