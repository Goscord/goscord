package packet

type VoiceSelectProtocol struct {
	Packet
	Data struct {
		Protocol string `json:"protocol"`
		Data     struct {
			Address string `json:"address"`
			Port    uint16 `json:"port"`
			Mode    string `json:"mode"` // always "xsalsa20_poly1305"
		} `json:"data"`
	} `json:"d,omitempty"`
}

func NewVoiceSelectProtocol(address string, port uint16) *VoiceSelectProtocol {
	voiceSelect := new(VoiceSelectProtocol)

	voiceSelect.Opcode = OpVoiceSelectProtocol

	voiceSelect.Data.Protocol = "udp"
	voiceSelect.Data.Data.Address = address
	voiceSelect.Data.Data.Port = port
	voiceSelect.Data.Data.Mode = "xsalsa20_poly1305"

	return voiceSelect
}
