package main

import (
	"fmt"
	"github.com/Goscord/goscord"
	"github.com/Goscord/goscord/discord"
	"github.com/Goscord/goscord/gateway"
)

var client *gateway.Session

func main() {
	fmt.Println("Testing...")

	client = goscord.New(&gateway.Options{
		Token: "",
	})

	_ = client.On("ready", OnReady)
	_ = client.On("message", OnMessage)

	if err := client.Login(); err != nil {
		panic(err)
	}

	select {}
}

func OnReady() {
	fmt.Println("Logged in as " + client.User().Tag())
	
	_ = client.SetActivity(&discord.Activity{Name: "Spotify", Type: 2})
	_ = client.SetStatus("dnd")
}

func OnMessage(msg *discord.Message) {
	if !msg.Author.Bot {
		_, _ = msg.Reply("coucou mec")

		c, err := msg.Channel()

		if err != nil {
			panic(err)
		}

		_, _ = c.Send("ça va mec ?")
	}
}
