package packet

type ConnectionProperties struct {
    Os string `json:"$os,omitempty"`
    Browser string `json:"$browser,omitempty"`
    Device string `json:"$device,omitempty"`
}

type Identify struct {
	Packet
	Data struct {
	    Token string `json:"token"`
	    Properties *ConnectionProperties `json:"properties,omitempty"`
	    Intents int `json:"intents,omitempty"`
	    Compress bool `json:"compress,omitempty"`
	} `json:"d,omitempty"`
}

func newConnectionProperties(os, browser, device string) *ConnectionProperties {
    return &ConnectionProperties{
        Os: os,
        Browser: browser,
        Device: device,
    }
}

func NewIdentify(token string) *Identify { // ToDo : Compression
    identify := &Identify{}
    
    identify.Opcode = OpIdentify
    identify.Data.Token = token
    identify.Data.Properties = newConnectionProperties("android", "test", "test")
    //identify.Data.Compress = false
    //identify.Data.Intents = 0
    
    return identify
}