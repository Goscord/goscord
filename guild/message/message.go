package message

type Message struct {
	Id        string `json:"id"`
	ChannelId string `json:"channel_id"`
	GuildId   string `json:"guild_id,omitempty"`
	//Author    *user.User `json:"author"`
	//Member *guild.Member `json:"member"`
	Content string `json:"content"`
	// ToDo : Add other properties
}
