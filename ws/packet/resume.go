package packet

type Resume struct {
	Packet
	Data struct {
	    Token string `json:"token"`
	    SessionID string `json:"session_id"`
	    Sequence int `json:"seq"`
	} `json:"d,omitempty"`
}

func NewResume(token, sessionID string, sequence int) *Resume {
    resume := &Resume{}
    
    resume.Opcode = OpResume
    resume.Data.Token = token
    resume.Data.SessionID = sessionID
    resume.Data.Sequence = sequence
    
    fmt.Println(sequence)
    
    return resume
}