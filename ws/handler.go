package ws

type EventHandler interface {
    Handle(s *Session, data []byte)
}