<p align="center">
  <a href="https://goscord.dev">
    <img src="./resource/logo.png" height="128">
    <h1 align="center">Goscord</h1>
  </a>
</p>

<p align="center">Goscord is a package for <a href="https://golang.org">Golang</a> that provides high level bindings to the <a href="https://discord.com">Discord</a> API.</p>

<p align="center">
  <a href="https://discord.gg/6Np8sbyHXt">
    <img src="https://badgen.net/badge/icon/Discord?icon=discord&label">
  </a>
  <a href="https://goscord.dev">
    <img src="https://badgen.net/badge/icon/Website?icon=chrome&label">
  </a>
</p>

## Getting Started
### Installing
```sh
# Init the module:
go mod init <url> 

# Install Goscord:
go get -u github.com/Goscord/goscord
```

### Usage
Construct a new Discord client which can be used to access the variety of 
Discord API functions and to set callback functions for Discord events.

```go
package main

import (
    "fmt"


    "github.com/Goscord/goscord/goscord"
    "github.com/Goscord/goscord/goscord/discord"
    "github.com/Goscord/goscord/goscord/gateway"
    "github.com/Goscord/goscord/goscord/gateway/event"
)

var client *gateway.Session

func main() {
    fmt.Println("Starting...")

    client := goscord.New(&gateway.Options{
        Token:   "token",
        Intents: gateway.IntentGuildMessages,
    })

	client.On(event.EventReady, func() {
		fmt.Println("Logged in as " + client.Me().Tag())
	})

	client.On(event.EventMessageCreate, func(msg *discord.Message) {
		if msg.Content == "ping" {
			client.Channel.SendMessage(msg.ChannelId, "Pong ! 🏓")
		}
	})


    client.Login()

    select {}
}
```

See [documentation](https://goscord.dev/documentation) for more detailed information.

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change. 
