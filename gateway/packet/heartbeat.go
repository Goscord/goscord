package packet

type Heartbeat struct {
	Packet
	Data int64 `json:"d"`
}

func NewHeartbeat(lastSequence int64) *Heartbeat {
	heartbeat := &Heartbeat{}

	heartbeat.Opcode = OpHeartbeat
	heartbeat.Data = lastSequence

	return heartbeat
}
