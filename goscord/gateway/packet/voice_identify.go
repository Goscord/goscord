package packet

type VoiceIdentify struct {
	Packet
	Data struct {
		ServerId  string `json:"server_id"`
		UserID    string `json:"user_id"`
		SessionId string `json:"session_id"`
		Token     string `json:"token"`
	} `json:"d,omitempty"`
}

func NewVoiceIdentify(serverId, userId, sessionId, token string) *VoiceIdentify {
	identify := new(VoiceIdentify)

	identify.Opcode = OpVoiceIdentify

	identify.Data.ServerId = serverId
	identify.Data.UserID = userId
	identify.Data.SessionId = sessionId
	identify.Data.Token = token

	return identify
}
