# Goscord
Goscord is a [Go](https://golang.org/) package that provides high level bindings to the [Discord](https://discord.com/) API.  

You can join our community on our [Discord server](https://discord.gg/6Np8sbyHXt).  

## Getting Started
### Installing
```sh
go get -u github.com/Goscord/goscord
```

### Usage
Construct a new Discord client which can be used to access the variety of 
Discord API functions and to set callback functions for Discord events.

```go
package main

import (
    "fmt"

    "github.com/Goscord/goscord"
    "github.com/Goscord/goscord/discord"
    "github.com/Goscord/goscord/gateway"
)

var client *gateway.Session

func main() {
    fmt.Println("Starting...")

    client := goscord.New(&gateway.Options{ 
        Token: "token", 
        Intents: gateway.IntentGuildMessages,
    })

    client.On("ready", func() {
        fmt.Println("Logged in as " + client.Me().Tag())
    })

    client.On("messageCreate", func(msg *discord.Message) {
        if msg.Content == "ping" {
            client.Channel.Send(msg.ChannelId, "Pong ! 🏓")
        }
    })

    client.Login()

    select {}
}
```

See [documentation](https://goscord.dev/documentation) for more detailed information.

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change. 
