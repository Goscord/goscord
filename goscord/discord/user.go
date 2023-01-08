package discord

import (
	"fmt"
	"strings"
)

type UserFlag int

const (
	UserFlagStaff UserFlag = 1 << iota
	UserFlagPartner
	UserFlagHypesquad
	UserFlagBugHunterLevel1
	UserFlagHypesquadOnlineHouse1 // Bravery member
	UserFlagHypesquadOnlineHouse2 // Brilliance member
	UserFlagHypesquadOnlineHouse3 // Balance member
	UserFlagEarlySupporter
	UserFlagPseudoUser // User in team
	UserFlagBugHunterLevel2
	UserFlagVerifiedBot
	UserFlagVerifiedDeveloper
	UserFlagCertifiedModerator
	UserFlagBotHTTPInteractions // 	Bot uses only HTTP interactions and is shown in the online member list
)

type PremiumType int

const (
	PremiumTypeNone PremiumType = iota
	PremiumTypeNitroClassic
	PremiumTypeNitro
)

type Service string

const (
	ServiceBattleNet       Service = "battlenet"
	ServiceEpicGames       Service = "epicgames"
	ServiceFacebook        Service = "facebook"
	ServiceGithub          Service = "github"
	ServiceLeagueofLegends Service = "leagueoflegends"
	ServicePlaystation     Service = "playstation"
	ServiceReddit          Service = "reddit"
	ServiceSamsungGalaxy   Service = "samsunggalaxy"
	ServiceSpotify         Service = "spotify"
	ServiceSkype           Service = "skype"
	ServiceSteam           Service = "steam"
	ServiceTwitch          Service = "twitch"
	ServiceTwitter         Service = "twitter"
	ServiceXbox            Service = "xbox"
	ServiceYoutube         Service = "youtube"
)

type VisibilityType int

const (
	VisibilityTypeNone VisibilityType = iota
	VisibilityTypeEveryone
)

type Connection struct {
	Id           string         `json:"id"`
	Name         string         `json:"name"`
	Type         Service        `json:"type"`
	Revoked      bool           `json:"revoked,omitempty"`
	Integrations []*Integration `json:"integrations,omitempty"`
	Verified     bool           `json:"verified"`
	FriendSync   bool           `json:"friend_sync"`
	ShowActivity bool           `json:"show_activity"`
	Visibility   VisibilityType `json:"visibility"`
}

type User struct {
	Id            string      `json:"id"`
	Username      string      `json:"username"`
	Discriminator string      `json:"discriminator"`
	Avatar        string      `json:"avatar"`
	Bot           bool        `json:"bot,omitempty"`
	System        bool        `json:"system,omitempty"`
	MfaEnabled    bool        `json:"mfa_enabled,omitempty"`
	Banner        string      `json:"banner,omitempty"`
	AccentColor   int         `json:"accent_color,omitempty"`
	Locale        Locale      `json:"locale,omitempty"`
	Verified      bool        `json:"verified,omitempty"`
	Email         string      `json:"email,omitempty"`
	Flags         UserFlag    `json:"flags,omitempty"`
	PremiumType   PremiumType `json:"premium_type,omitempty"`
	PublicFlags   UserFlag    `json:"public_flags,omitempty"`
}

func (u *User) Tag() string {
	return fmt.Sprintf("%s#%s", u.Username, u.Discriminator)
}

func (u *User) AvatarURL() string {
	if strings.HasPrefix(u.Avatar, "a_") {
		return fmt.Sprintf("https://cdn.discordapp.com/avatars/%s/%s.gif", u.Id, u.Avatar)
	}

	return fmt.Sprintf("https://cdn.discordapp.com/avatars/%s/%s.png", u.Id, u.Avatar)
}
