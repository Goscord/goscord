package discord

import "time"

const (
	EmbedDefault    = 0x000000
	EmbedWhite      = 0xffffff
	EmbedAqua       = 0x1abc9c
	EmbedGreen      = 0x2ecc71
	EmbedBlue       = 0x3498db
	EmbedYellow     = 0xffff00
	EmbedPurple     = 0x9b59b6
	EmbedGold       = 0xf1c40f
	EmbedOrange     = 0xe67e22
	EmbedRed        = 0xe74c3c
	EmbedGrey       = 0x95a5a6
	EmbedDarkAqua   = 0x11806a
	EmbedDarkGreen  = 0x1f8b4c
	EmbedDarkBlue   = 0x206694
	EmbedDarkPurple = 0x71368a
	EmbedDarkGold   = 0xc27c0e
	EmbedDarkOrange = 0xa84300
	EmbedDarkRed    = 0x992d22
	EmbedDarkGrey   = 0x979c9f
	EmbedBlurple    = 0x7289da
	EmbedDark       = 0x2c2f33
)

type EmbedType string

const (
	EmbedTypeRich    EmbedType = "rich"
	EmbedTypeImage   EmbedType = "image"
	EmbedTypeVideo   EmbedType = "video"
	EmbedTypeGifv    EmbedType = "gifv"
	EmbedTypeArticle EmbedType = "article"
	EmbedTypeLink    EmbedType = "link"
)

type Embed struct {
	Title       string          `json:"title,omitempty"`
	Type        EmbedType       `json:"type,omitempty"`
	Description string          `json:"description,omitempty"`
	URL         string          `json:"url,omitempty"`
	Timestamp   *time.Time      `json:"timestamp,omitempty"`
	Color       int             `json:"color,omitempty"`
	Footer      *EmbedFooter    `json:"footer,omitempty"`
	Image       *EmbedImage     `json:"image,omitempty"`
	Thumbnail   *EmbedThumbnail `json:"thumbnail,omitempty"`
	Video       *EmbedVideo     `json:"video,omitempty"`
	Provider    *EmbedProvider  `json:"provider,omitempty"`
	Author      *EmbedAuthor    `json:"author,omitempty"`
	Fields      []*EmbedField   `json:"fields,omitempty"`
}

type EmbedThumbnail struct {
	URL      string `json:"url"`
	ProxyURL string `json:"proxy_url,omitempty"`
	Height   int    `json:"height,omitempty"`
	Width    int    `json:"width,omitempty"`
}

type EmbedVideo struct {
	URL      string `json:"url,omitempty"`
	ProxyUrl string `json:"proxy_url,omitempty"`
	Height   int    `json:"height,omitempty"`
	Width    int    `json:"width,omitempty"`
}

type EmbedImage struct {
	URL      string `json:"url"`
	ProxyURL string `json:"proxy_url,omitempty"`
	Height   int    `json:"height,omitempty"`
	Width    int    `json:"width,omitempty"`
}

type EmbedProvider struct {
	Name string `json:"name,omitempty"`
	URL  string `json:"url,omitempty"`
}

type EmbedAuthor struct {
	Name         string `json:"name"`
	URL          string `json:"url,omitempty"`
	IconURL      string `json:"icon_url,omitempty"`
	ProxyIconURL string `json:"proxy_icon_url,omitempty"`
}

type EmbedFooter struct {
	Text         string `json:"text"`
	IconURL      string `json:"icon_url,omitempty"`
	ProxyIconURL string `json:"proxy_icon_url,omitempty"`
}

type EmbedField struct {
	Name   string `json:"name"`
	Value  string `json:"value"`
	Inline bool   `json:"inline,omitempty"`
}
