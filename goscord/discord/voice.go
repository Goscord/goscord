package discord

import "time"

type VoiceState struct {
	GuildId                 string       `json:"guild_id,omitempty"`
	ChannelId               string       `json:"channel_id"`
	UserId                  string       `json:"user_id"`
	Member                  *GuildMember `json:"member,omitempty"`
	SessionId               string       `json:"session_id"`
	Deaf                    bool         `json:"deaf"`
	Mute                    bool         `json:"mute"`
	SelfDeaf                bool         `json:"self_deaf"`
	SelfMute                bool         `json:"self_mute"`
	SelfStream              bool         `json:"self_stream,omitempty"`
	SelfVideo               bool         `json:"self_video"`
	Suppress                bool         `json:"suppress"`
	RequestToSpeakTimestamp time.Time    `json:"request_to_speak_timestamp"`
}

type VoiceRegion struct {
	Id         string `json:"id"`
	Name       string `json:"name"`
	Optimal    bool   `json:"optimal"`
	Deprecated bool   `json:"deprecated"`
	Custom     bool   `json:"custom"`
}
