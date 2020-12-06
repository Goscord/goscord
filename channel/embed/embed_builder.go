package embed

import "time"

type Builder struct {
	embed *Embed
}

func NewEmbedBuilder() *Builder {
	b := &Builder{}

	b.embed = &Embed{}
	b.embed.Type = "rich"

	return b
}

func (b *Builder) SetTitle(title string) *Builder {
	b.embed.Title = title
	return b
}

func (b *Builder) SetDescription(description string) *Builder {
	b.embed.Description = description
	return b
}

func (b *Builder) SetURL(url string) *Builder {
	b.embed.URL = url
	return b
}

func (b *Builder) SetTimestamp(time *time.Time) *Builder {
	b.embed.Timestamp = time
	return b
}

func (b *Builder) SetColor(color int) *Builder {
	b.embed.Color = color
	return b
}

func (b *Builder) SetFooter(text string, icon string) *Builder {
	b.embed.Footer = &Footer{Text: text, IconURL: icon}
	return b
}

func (b *Builder) SetThumbnail(url string) *Builder {
	b.embed.Thumbnail = &Thumbnail{URL: url}
	return b
}

func (b *Builder) SetImage(url string) *Builder {
	b.embed.Image = &Image{URL: url}
	return b
}

func (b *Builder) SetAuthor(name string, icon string) *Builder {
	b.embed.Author = &Author{Name: name, IconURL: icon}
	return b
}

func (b *Builder) AddField(name, value string, inline bool) *Builder {
	b.embed.Fields = append(b.embed.Fields, &Field{Name: name, Value: value, Inline: inline})
	return b
}

func (b *Builder) Embed() *Embed {
	return b.embed
}