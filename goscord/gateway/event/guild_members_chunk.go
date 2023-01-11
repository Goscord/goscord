package event

import (
	"github.com/Goscord/goscord/goscord/discord"
	"github.com/Goscord/goscord/goscord/rest"
	"github.com/goccy/go-json"
)

// GuildMembersChunk Is sent in response to Guild Request Members. You can use the chunk_index and chunk_count to calculate how many chunks are left for your request.
type GuildMembersChunk struct {
	Data *discord.GuildMembersChunkEventFields `json:"d"`
}

func NewGuildMembersChunk(rest *rest.Client, data []byte) (*GuildMembersChunk, error) {
	pk := new(GuildMembersChunk)

	err := json.Unmarshal(data, pk)

	if err != nil {
		return nil, err
	}

	return pk, nil
}
