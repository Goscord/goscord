package packet

type Heartbeat struct {
	Packet
	Data int `json:"d"`
}

func NewHeartbeat(lastSequence int) (*Heartbeat) {
	heartbeat := &Heartbeat{}
	
	heartbeat.Opcode = OpHeartbeat
	heartbeat.Data = lastSequence
	
	return heartbeat
}