package discord

type Role struct {
	Id          string   `json:"id"`
	Name        string   `json:"name"`
	Color       int      `json:"color"`
	Hoist       bool     `json:"hoist"`
	Position    int      `json:"position"`
	Permissions string   `json:"permissions"`
	Managed     bool     `json:"managed"`
	Mentionable bool     `json:"mentionable"`
	Tags        *RoleTag `json:"tags"`
}
