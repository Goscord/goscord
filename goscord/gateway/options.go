package gateway

// Options is the options for the gateway, it contains the token, intents and some other options.
type Options struct {
	Token   string
	Intents Intents
	//NumShards int
}
