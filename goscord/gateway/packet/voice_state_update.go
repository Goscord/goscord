package packet

type VoiceStateUpdate struct {
	Packet
	Data struct {
		GuildId   string `json:"guild_id"`
		ChannelId string `json:"channel_id"`
		SelfMute  bool   `json:"self_mute"`
		SelfDeaf  bool   `json:"self_deaf"`
	} `json:"d,omitempty"`
}

func NewVoiceStateUpdate(guildId, channelId string, selfMuted, selfDeaf bool) *VoiceStateUpdate {
	voice := &VoiceStateUpdate{}

	voice.Data.GuildId = guildId
	voice.Data.ChannelId = channelId
	voice.Data.SelfMute = selfMuted
	voice.Data.SelfDeaf = selfDeaf

	return voice
}
