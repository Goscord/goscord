package discord

type Emoji struct {
	Id            string    `json:"id"`
	Name          string    `json:"name"`
	Roles         []*string `json:"roles"`
	User          *User     `json:"user"`
	RequireColons bool      `json:"require_colons"`
	Managed       bool      `json:"managed"`
	Animated      bool      `json:"animated"`
	Available     bool      `json:"available"`
}
