package handler

type EventHandler interface {
    Handle(data []byte)
}