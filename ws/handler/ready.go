package handler

import (
    _ "fmt"
    "github.com/Seyz123/yalis/ws/event"
    ev "github.com/asaskevich/EventBus"
)

type ReadyHandler struct {
    bus *ev.EventBus
}

func NewReady(bus *ev.EventBus) ReadyHandler {
    h := ReadyHandler{}
    
    h.bus = bus
    
    return h
}

func (h ReadyHandler) Handle(data []byte) {
    _, err := event.NewReady(data)
    
    if err != nil {
        return
    }
    
    // ToDo : Handle properties
    
    h.bus.Publish("ready")
}