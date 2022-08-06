package packet

import "runtime"

type ConnectionProperties struct {
	Os      string `json:"$os,omitempty"`
	Browser string `json:"$browser,omitempty"`
	Device  string `json:"$device,omitempty"`
}

type Identify struct {
	Packet
	Data struct {
		Token          string                `json:"token"`
		Properties     *ConnectionProperties `json:"properties"`
		Compress       bool                  `json:"compress,omitempty"`
		LargeThreshold int                   `json:"large_threshold,omitempty"`
		Shard          []int                 `json:"shard,omitempty"`
		Presence       *UpdateStatus         `json:"presence,omitempty"`
		Intents        int                   `json:"intents,omitempty"`
	} `json:"d,omitempty"`
}

func newConnectionProperties(os, browser, device string) *ConnectionProperties {
	return &ConnectionProperties{
		Os:      os, // windows, linux, mac, ios, android
		Browser: browser,
		Device:  device,
	}
}

func NewIdentify(token string, intents int) *Identify {
	identify := &Identify{}

	identify.Opcode = OpIdentify
	identify.Data.Token = token
	identify.Data.Intents = intents
	identify.Data.Properties = newConnectionProperties(runtime.GOOS, "Goscord", "Goscord")
	identify.Data.Compress = false

	return identify
}
