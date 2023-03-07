package builder

import (
	"github.com/Goscord/goscord/goscord/discord"
	"time"
)

type EmbedBuilder struct {
	content string
	embed   *discord.Embed
}

func NewEmbedBuilder() *EmbedBuilder {
	b := &EmbedBuilder{}

	b.embed = &discord.Embed{}
	b.embed.Type = discord.EmbedTypeRich

	return b
}

func (b *EmbedBuilder) SetContent(content string) *EmbedBuilder {
	b.content = content
	return b
}

func (b *EmbedBuilder) SetTitle(title string) *EmbedBuilder {
	b.embed.Title = title
	return b
}

func (b *EmbedBuilder) SetDescription(description string) *EmbedBuilder {
	b.embed.Description = description
	return b
}

func (b *EmbedBuilder) SetURL(url string) *EmbedBuilder {
	b.embed.URL = url
	return b
}

func (b *EmbedBuilder) SetTimestamp(time *time.Time) *EmbedBuilder {
	b.embed.Timestamp = time
	return b
}

func (b *EmbedBuilder) SetColor(color int) *EmbedBuilder {
	b.embed.Color = color
	return b
}

func (b *EmbedBuilder) SetFooter(text string, icon string) *EmbedBuilder {
	b.embed.Footer = &discord.EmbedFooter{Text: text, IconURL: icon}
	return b
}

func (b *EmbedBuilder) SetThumbnail(url string) *EmbedBuilder {
	b.embed.Thumbnail = &discord.EmbedThumbnail{URL: url}
	return b
}

func (b *EmbedBuilder) SetImage(url string) *EmbedBuilder {
	b.embed.Image = &discord.EmbedImage{URL: url}
	return b
}

func (b *EmbedBuilder) SetAuthor(name string, icon string) *EmbedBuilder {
	b.embed.Author = &discord.EmbedAuthor{Name: name, IconURL: icon}
	return b
}

func (b *EmbedBuilder) AddField(name, value string, inline bool) *EmbedBuilder {
	b.embed.Fields = append(b.embed.Fields, &discord.EmbedField{Name: name, Value: value, Inline: inline})
	return b
}

func (b *EmbedBuilder) Content() string {
	return b.content
}

func (b *EmbedBuilder) Embed() *discord.Embed {
	return b.embed
}
