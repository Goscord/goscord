package discord

import (
	"encoding/json"
	"fmt"
	"github.com/bytedance/sonic"
)

type ComponentType int

const (
	ComponentTypeActionRow  ComponentType = iota + 1 // container for other components
	ComponentTypeButton                              // clickable button
	ComponentTypeSelectMenu                          // dropdown
	ComponentTypeTextInput                           // text input
)

type ButtonStyle int

const (
	ButtonStylePrimary ButtonStyle = iota + 1
	ButtonStyleSecondary
	ButtonStyleSuccess
	ButtonStyleDanger
	ButtonStyleLink
)

type TextInputStyle uint

const (
	TextInputShort TextInputStyle = iota + 1
	TextInputParagraph
)

type MessageComponent interface {
	json.Marshaler
}

type unmarshalableMessageComponent struct {
	MessageComponent
	json.Unmarshaler
}

func (u *unmarshalableMessageComponent) UnmarshalJSON(data []byte) error {
	var v struct {
		Type ComponentType `json:"type"`
	}

	err := sonic.Unmarshal(data, &v)
	if err != nil {
		return err
	}

	switch v.Type {
	case ComponentTypeActionRow:
		u.MessageComponent = &ActionRows{}

	case ComponentTypeButton:
		u.MessageComponent = &Button{}

	case ComponentTypeSelectMenu:
		u.MessageComponent = &SelectMenu{}

	case ComponentTypeTextInput:
		u.MessageComponent = &TextInput{}

	default:
		return fmt.Errorf("unknown component type: %d", v.Type)
	}

	return sonic.Unmarshal(data, u.MessageComponent)
}

type ActionRows struct {
	Type       ComponentType      `json:"type"`       // 1 for an action rows
	Components []MessageComponent `json:"components"` // Can contain only one type of component
}

func (r ActionRows) MarshalJSON() ([]byte, error) {
	type actionsRow ActionRows

	return sonic.Marshal(struct {
		Type ComponentType `json:"type"`
		actionsRow
	}{
		ComponentTypeActionRow,
		actionsRow(r),
	})
}

func (r ActionRows) UnmarshalJSON(data []byte) error {
	var v struct {
		Components []unmarshalableMessageComponent `json:"components"`
	}

	err := sonic.Unmarshal(data, &v)
	if err != nil {
		return err
	}

	r.Components = make([]MessageComponent, len(v.Components))

	for i, v := range v.Components {
		r.Components[i] = v.MessageComponent
	}

	return err
}

type Button struct {
	Type     ComponentType `json:"type"` // 2 for a button
	Style    ButtonStyle   `json:"style"`
	Label    string        `json:"label,omitempty"`
	Emoji    *Emoji        `json:"emoji,omitempty"` // name, id, and animated required
	CustomId string        `json:"custom_id,omitempty"`
	Url      string        `json:"url,omitempty"`
	Disabled bool          `json:"disabled,omitempty"`
}

func (b Button) MarshalJSON() ([]byte, error) {
	type button Button

	return sonic.Marshal(struct {
		Type ComponentType `json:"type"`
		button
	}{
		ComponentTypeButton,
		button(b),
	})
}

type SelectMenu struct {
	Type         ComponentType   `json:"type"` // 3 fo a select menu
	CustomId     string          `json:"custom_id"`
	Options      []*SelectOption `json:"options"`
	ChannelTypes []*ChannelType  `json:"channel_types,omitempty"`
	PlaceHolder  string          `json:"placeholder,omitempty"`
	MinValues    int             `json:"min_values,omitempty"`
	MaxValues    int             `json:"max_values,omitempty"`
	Disabled     bool            `json:"disabled,omitempty"`
}

func (sm SelectMenu) MarshalJSON() ([]byte, error) {
	type selectMenu SelectMenu

	return sonic.Marshal(struct {
		Type ComponentType `json:"type"`
		selectMenu
	}{
		ComponentTypeSelectMenu,
		selectMenu(sm),
	})
}

type TextInput struct {
	Type        ComponentType  `json:"type"` // 4 for a text input
	CustomId    string         `json:"custom_id"`
	Style       TextInputStyle `json:"style"`
	Label       string         `json:"label"`
	MinLength   int            `json:"min_length,omitempty"`
	MaxLength   int            `json:"max_length,omitempty"`
	Required    bool           `json:"required,omitempty"`
	Value       string         `json:"value,omitempty"`
	PlaceHolder string         `json:"placeholder,omitempty"`
}

func (ti TextInput) MarshalJSON() ([]byte, error) {
	type textInput TextInput

	return sonic.Marshal(struct {
		Type ComponentType `json:"type"`
		textInput
	}{
		ComponentTypeTextInput,
		textInput(ti),
	})
}

type SelectOption struct {
	Label       string `json:"label"`
	Value       string `json:"value"`
	Description string `json:"description,omitempty"`
	Emoji       *Emoji `json:"emoji,omitempty"` // id, name, and animated required
	Default     bool   `json:"default,omitempty"`
}
