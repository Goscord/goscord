package discord

type ApplicationFlag int

const (
	ApplicationFlagGatewayPresence              ApplicationFlag = 1 << 12
	ApplicationFlagPresenceLimited              ApplicationFlag = 1 << 13
	ApplicationFlagGatewayGuildMembers          ApplicationFlag = 1 << 14
	ApplicationFlagGatewayGuildMembersLimited   ApplicationFlag = 1 << 15
	ApplicationFlagVerificationGuildLimit       ApplicationFlag = 1 << 16
	ApplicationFlagEmbedded                     ApplicationFlag = 1 << 17
	ApplicationFlagGatewayMessageContent        ApplicationFlag = 1 << 18
	ApplicationFlagGatewayMessageContentLimited ApplicationFlag = 1 << 19
)

type InstallParams struct {
	Scopes      []string              `json:"scopes"`
	Permissions BitwisePermissionFlag `json:"permissions,string"`
}

type Application struct {
	Id                  string          `json:"id"`
	Name                string          `json:"name"`
	Icon                string          `json:"icon"`
	Description         string          `json:"description"`
	RpcOrigins          []string        `json:"rpc_origins,omitempty"`
	BotPublic           bool            `json:"bot_public"`
	BotRequireGrantCode bool            `json:"bot_require_grant_code"`
	TermsOfServiceUrl   string          `json:"terms_of_service_url,omitempty"`
	PrivacyPolicyUrl    string          `json:"privacy_policy_url,omitempty"`
	Owner               *User           `json:"owner,omitempty"`
	Summary             string          `json:"summary"` // deprecated
	VerifyKey           string          `json:"verify_key"`
	Team                *Team           `json:"team"`
	GuildId             string          `json:"guild_id,omitempty"`
	PrimarySkuId        string          `json:"primary_sku_id,omitempty"`
	Slug                string          `json:"slug,omitempty"`
	CoverImage          string          `json:"cover_image,omitempty"`
	Flags               ApplicationFlag `json:"flags,omitempty"`
	Tags                []string        `json:"tags,omitempty"`
	InstallParams       *InstallParams  `json:"install_params,omitempty"`
	CustomInstallUrl    string          `json:"custom_install_url,omitempty"`
}
