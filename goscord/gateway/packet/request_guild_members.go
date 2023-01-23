package packet

type GuildRequestMembers struct {
	Packet
	Data struct {
		GuildId   string   `json:"guild_id"`            // ID of the guild to get members for
		Query     string   `json:"query,omitempty"`     // string that username starts with, or an empty string to return all members
		Limit     int      `json:"limit"`               // max number of members to send matching the query; a limit of 0 can be used with an empty string query to return all members
		Presences bool     `json:"presences,omitempty"` // used to specify if we want the presences of the matched members
		UserIds   []string `json:"user_ids,omitempty"`  // used to specify which users you wish to fetch
		Nonce     string   `json:"nonce,omitempty"`     // the nonce to identify the Guild Members Chunk response
	} `json:"d"`
}

// ToDo : Implement NewGuildRequestMembers function
