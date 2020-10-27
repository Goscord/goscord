package ws

import (
	"fmt"

	"github.com/Seyz123/yalis/ws/event"
)

type GuildCreateHandler struct{}

func (h *GuildCreateHandler) Handle(s *Session, data []byte) {
	ev, err := event.NewGuildCreate(data)

	if err != nil {
		return
	}

	fmt.Println("Got guild create")
	fmt.Println(ev.Data)
}
