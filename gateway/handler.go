package gateway

type EventHandler interface {
    Handle(s *Session, data []byte)
}