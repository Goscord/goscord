package builder

import (
	"github.com/Goscord/goscord/goscord/discord"
	"io"
)

// FileData is used to send files to Discord
type FileData struct {
	Name   string
	Reader io.Reader
}

// MessageBuilder is used to build a message
type MessageBuilder struct {
	content string
	embeds  []*discord.Embed
	files   []*FileData
	flags   discord.MessageFlag
}

// NewMessageBuilder creates a new MessageBuilder
func NewMessageBuilder() *MessageBuilder {
	return &MessageBuilder{}
}

// SetContent sets the content of the message
func (b *MessageBuilder) SetContent(content string) *MessageBuilder {
	b.content = content
	return b
}

// SetEmbeds sets the builder of the message
func (b *MessageBuilder) SetEmbeds(embeds []*discord.Embed) *MessageBuilder {
	b.embeds = embeds
	return b
}

// AddEmbed embed adds an embed to the message
func (b *MessageBuilder) AddEmbed(embed *discord.Embed) *MessageBuilder {
	b.embeds = append(b.embeds, embed)
	return b
}

// SetFiles sets the files of the message
func (b *MessageBuilder) SetFiles(files []*FileData) *MessageBuilder {
	b.files = files
	return b
}

// AddFile adds a file to the message
func (b *MessageBuilder) AddFile(name string, reader io.Reader) *MessageBuilder {
	b.files = append(b.files, &FileData{Name: name, Reader: reader})
	return b
}

// SetFlags sets the flags of the message (e.g. ephemeral)
func (b *MessageBuilder) SetFlags(flags discord.MessageFlag) *MessageBuilder {
	b.flags = flags
	return b
}

// AddFlag adds a flag to the message (e.g. ephemeral)
func (b *MessageBuilder) AddFlag(flag discord.MessageFlag) *MessageBuilder {
	b.flags |= flag
	return b
}

// Content returns the content of the message
func (b *MessageBuilder) Content() string {
	return b.content
}

// Embeds returns the embed of the message
func (b *MessageBuilder) Embeds() []*discord.Embed {
	return b.embeds
}

// Files returns the files of the message
func (b *MessageBuilder) Files() []*FileData {
	return b.files
}

// Build builds the message
func (b *MessageBuilder) Build() *discord.Message {
	return &discord.Message{
		Content: b.content,
		Embeds:  b.embeds,
		Flags:   b.flags,
	}
}
