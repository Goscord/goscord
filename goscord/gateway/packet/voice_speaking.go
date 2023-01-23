package packet

type VoiceSpeaking struct {
	Packet
	Data struct {
		Speaking int    `json:"speaking"`
		Delay    int    `json:"delay"`
		SSRC     uint32 `json:"ssrc"`
	} `json:"d,omitempty"`
}

func NewVoiceSpeaking(speaking bool, delay int, ssrc uint32) *VoiceSpeaking {
	voiceSpeaking := new(VoiceSpeaking)

	voiceSpeaking.Opcode = OpVoiceSpeaking

	if speaking {
		voiceSpeaking.Data.Speaking = 1
	} else {
		voiceSpeaking.Data.Speaking = 0
	}

	voiceSpeaking.Data.Delay = delay
	voiceSpeaking.Data.SSRC = ssrc

	return voiceSpeaking
}
