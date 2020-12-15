package gateway

type ResumedHandler struct{}

func (h *ResumedHandler) Handle(s *Session, _ []byte) {
	s.bus.Publish("resumed")
}
