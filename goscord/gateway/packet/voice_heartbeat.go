package packet

type VoiceHeartbeat struct {
	Packet
	Data int64 `json:"d"`
}

func NewVoiceHeartbeat(data int64) *VoiceHeartbeat {
	heartbeat := new(VoiceHeartbeat)

	heartbeat.Opcode = OpVoiceHeartbeat
	heartbeat.Data = data

	return heartbeat
}
