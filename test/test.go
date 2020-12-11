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
		Token: "NzM1NjQyNjE2NDc3MjUzNjg0.XxjOkw.dH9jmMRmNarhetycaAtvooq07Vg",
	})

	_ = client.On("ready", OnReady)
	_ = client.On("message", OnMessage)
	_ = client.On("guildCreate", OnGuildCreate)
	_ = client.On("guildDelete", OnGuildDelete)

	if err := client.Login(); err != nil {
		panic(err)
	}

	select {}
}

func OnReady() {
	fmt.Println("Logged in as " + client.User().Tag())

	_ = client.SetActivity(&discord.Activity{Name: "Spotify", Type: 3})
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

func OnGuildCreate(guild *discord.Guild) {
	fmt.Println("Guild create : " + guild.Id)
}

func OnGuildDelete(guild *discord.Guild) {
	fmt.Println("Guild delete : " + guild.Id)
}