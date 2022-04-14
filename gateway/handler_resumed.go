package gateway

type ResumedHandler struct{}

func (_ *ResumedHandler) Handle(s *Session, _ []byte) {
	s.bus.Publish("resumed")
}
