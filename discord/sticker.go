package discord

type StickerType int

const (
	StickerTypeStandard StickerType = iota + 1
	StickerTypeGuild
)

type StickerFormat int

const (
	StickerFormatPng StickerFormat = iota + 1
	StickerFormatApng
	StickerFormatLottie
)

type Sticker struct {
	Id          string        `json:"id"`
	PackId      string        `json:"pack_id,omitempty"`
	Name        string        `json:"name,omitempty"`
	Description string        `json:"description"`
	Tags        string        `json:"tags,omitempty"`  // autocomplete/suggestion tags for the sticker (max 200 characters)
	Asset       string        `json:"asset,omitempty"` // Deprecated previously the sticker asset hash, now an empty string
	Type        StickerType   `json:"type"`
	FormatType  StickerFormat `json:"format_type"`
	Available   bool          `json:"available,omitempty"`
	GuildId     string        `json:"guild_id,omitempty"`
	User        *User         `json:"user,omitempty"`
	SortValue   int           `json:"sort_value,omitempty"`
}

type StickerItem struct {
	Id         string `json:"id"`
	Name       string `json:"name"`
	FormatType string `json:"format_type"`
}

type StickerPack struct {
	Id             string    `json:"id"`
	Stickers       []Sticker `json:"stickers"`
	Name           string    `json:"name"`
	SkuId          string    `json:"sku_id"`
	CoverStickerId string    `json:"cover_sticker_id,omitempty"`
	Description    string    `json:"description"`
	BannerAssetId  string    `json:"banner_asset_id,omitempty"`
}
