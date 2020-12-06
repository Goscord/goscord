package embed

import "time"

type MessageEmbed struct {
	Content string `json:"content,omitempty"`
	Embed *Embed   `json:"embed"`
}

type Embed struct {
	Title string         `json:"title,omitempty"`
	Type string          `json:"type,omitempty"`
	Description string   `json:"description,omitempty"`
	URL string           `json:"url,omitempty"`
	Timestamp *time.Time `json:"timestamp,omitempty"`
	Color int            `json:"color,omitempty"`
	Footer *Footer       `json:"footer,omitempty"`
	Image *Image         `json:"image,omitempty"`
	Thumbnail *Thumbnail `json:"thumbnail,omitempty"`
	// TODO : Video
	Author *Author  `json:"author,omitempty"`
	Fields []*Field `json:"fields,omitempty"`
}

type Footer struct {
	Text string `json:"text,omitempty"`
	IconURL string `json:"icon_url,omitempty"`
	ProxyIconURL string `json:"proxy_icon_url,omitempty"`
}

type Image struct {
	URL string `json:"url,omitempty"`
	ProxyURL string `json:"proxy_url,omitempty"`
	Height int `json:"height,omitempty"`
	Width int `json:"width,omitempty"`
}

type Author struct {
	Name string `json:"name,omitempty"`
	URL string `json:"url,omitempty"`
	IconURL string `json:"icon_url,omitempty"`
	ProxyIconURL string `json:"proxy_icon_url,omitempty"`
}

type Field struct {
	Name string `json:"name,omitempty"`
	Value string `json:"value,omitempty"`
	Inline bool `json:"inline,omitempty"`
}

type Thumbnail struct {
	URL string `json:"url,omitempty"`
	ProxyURL string `json:"proxy_url,omitempty"`
	Height int `json:"height,omitempty"`
	Width int `json:"width,omitempty"`
}