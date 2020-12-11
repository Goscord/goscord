package discord

type Activity struct {
	Name string `json:"name"`
	Type int    `json:"type"`
	URL  string `json:"url,omitempty"`

	// TODO : Add others
}
