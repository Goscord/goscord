package discord

import "time"

type GuildTemplate struct {
	Code                    string    `json:"code"`
	Name                    string    `json:"name"`
	Description             string    `json:"description"`
	UsageCount              int       `json:"usage_count"`
	CreatorId               string    `json:"creator_id"`
	Creator                 *User     `json:"creator"`
	CreatedAt               time.Time `json:"created_at"`
	UpdatedAt               time.Time `json:"updated_at"`
	SourceGuildId           string    `json:"source_guild_id"`
	SerializedSourceGuildId string    `json:"serialized_source_guild"`
	IsDirty                 bool      `json:"is_dirty"`
}
