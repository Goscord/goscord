package discord

type ApplicationCommandType int

const (
	_                      ApplicationCommandType = iota
	ApplicationCommandChat                        // slash command
	ApplicationCommandUser
	ApplicationCommandMessage
)

type ApplicationCommandOptionType int

const (
	_ ApplicationCommandOptionType = iota
	ApplicationCommandOptionSubCommand
	ApplicationCommandOptionSubCommandGroup
	ApplicationCommandOptionString
	ApplicationCommandOptionInteger
	ApplicationCommandOptionBoolean
	ApplicationCommandOptionUser
	ApplicationCommandOptionChannel
	ApplicationCommandOptionRole
	ApplicationCommandOptionMentionable
	ApplicationCommandOptionNumber
	ApplicationCommandOptionAttachment
)

type ApplicationCommandPermissionType int

const (
	_ ApplicationCommandPermissionType = iota
	ApplicationCommandPermissionTypeRole
	ApplicationCommandPermissionTypeUser
	ApplicationCommandPermissionTypeChannel
)

type ApplicationCommand struct {
	Id                       string                 `json:"id,omitempty"`
	ApplicationId            string                 `json:"application_id,omitempty"`
	GuildId                  string                 `json:"guild_id,omitempty"`
	Version                  string                 `json:"version,omitempty"`
	Type                     ApplicationCommandType `json:"type,omitempty"`
	Name                     string                 `json:"name"`
	NameLocalizations        map[Locale]string      `json:"name_localizations,omitempty"`
	DefaultPermission        bool                   `json:"default_permission,omitempty"`
	DefaultMemberPermissions int64                  `json:"default_member_permissions,string,omitempty"`
	DMPermission             bool                   `json:"dm_permission,omitempty"`

	Description              string                      `json:"description,omitempty"`
	DescriptionLocalizations map[Locale]string           `json:"description_localizations,omitempty"`
	Options                  []*ApplicationCommandOption `json:"options"`
}

type ApplicationCommandOptionChoice struct {
	Name              string            `json:"name"`
	NameLocalizations map[Locale]string `json:"name_localizations,omitempty"`
	Value             interface{}       `json:"value"`
}

type ApplicationCommandPermissions struct {
	Id         string                           `json:"id"`
	Type       ApplicationCommandPermissionType `json:"type"`
	Permission bool                             `json:"permission"`
}

type ApplicationCommandPermissionsList struct {
	Permissions []*ApplicationCommandPermissions `json:"permissions"`
}

type GuildApplicationCommandPermissions struct {
	Id            string                           `json:"id"`
	ApplicationId string                           `json:"application_id"`
	GuilddId      string                           `json:"guild_id"`
	Permissions   []*ApplicationCommandPermissions `json:"permissions"`
}

type ApplicationCommandInteractionDataOption struct {
	Name    string                                     `json:"name"`
	Type    ApplicationCommandOptionType               `json:"type"`
	Value   interface{}                                `json:"value,omitempty"` // string, integer, or double
	Options []*ApplicationCommandInteractionDataOption `json:"options,omitempty"`
	Focused bool                                       `json:"focused,omitempty"`
}

type ApplicationCommandOption struct {
	Type                     ApplicationCommandOptionType      `json:"type"`
	Name                     string                            `json:"name"`
	NameLocalizations        map[Locale]string                 `json:"name_localizations,omitempty"`
	Description              string                            `json:"description,omitempty"`
	DescriptionLocalizations map[Locale]string                 `json:"description_localizations,omitempty"`
	ChannelTypes             []int                             `json:"channel_types"`
	Required                 bool                              `json:"required"`
	Options                  []*ApplicationCommandOption       `json:"options"`
	Autocomplete             bool                              `json:"autocomplete"`
	Choices                  []*ApplicationCommandOptionChoice `json:"choices"`
	MinValue                 float64                           `json:"min_value,omitempty"`
	MaxValue                 float64                           `json:"max_value,omitempty"`
	MinLength                int                               `json:"min_length,omitempty"`
	MaxLength                int                               `json:"max_length,omitempty"`
}
