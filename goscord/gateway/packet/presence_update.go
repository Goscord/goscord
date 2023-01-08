package packet

import (
	"github.com/Goscord/goscord/goscord/discord"
)

type PresenceUpdate struct {
	Packet
	Data struct {
		Since      int                  `json:"since"`
		Activities [1]*discord.Activity `json:"activities"`
		Status     discord.StatusType   `json:"status"`
		AFK        bool                 `json:"afk"`
	} `json:"d"`
}

func NewPresenceUpdate(activity *discord.Activity, status discord.StatusType) *PresenceUpdate {
	update := new(PresenceUpdate)

	update.Opcode = OpPresenceUpdate
	update.Data.Activities[0] = activity
	update.Data.Status = status
	update.Data.AFK = false

	return update
}
