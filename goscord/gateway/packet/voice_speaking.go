package packet

type SpeakingFlag uint64

const (
	NotSpeaking SpeakingFlag = 0
	Microphone  SpeakingFlag = 1 << iota
	Soundshare
	Priority
)

type VoiceSpeaking struct {
	Packet
	Data struct {
		Speaking SpeakingFlag `json:"speaking"`
		Delay    int          `json:"delay"`
		SSRC     uint32       `json:"ssrc"`
		UserId   string       `json:"user_id,omitempty"`
	} `json:"d"`
}

func NewVoiceSpeaking(speaking bool, ssrc uint32) *VoiceSpeaking {
	voiceSpeaking := new(VoiceSpeaking)

	voiceSpeaking.Opcode = OpVoiceSpeaking

	if speaking {
		voiceSpeaking.Data.Speaking = Microphone
	} else {
		voiceSpeaking.Data.Speaking = NotSpeaking
	}

	voiceSpeaking.Data.Delay = 0
	voiceSpeaking.Data.SSRC = ssrc

	return voiceSpeaking
}
