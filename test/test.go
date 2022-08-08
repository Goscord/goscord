package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/Goscord/goscord"
	"github.com/Goscord/goscord/discord"
	"github.com/Goscord/goscord/discord/embed"
	"github.com/Goscord/goscord/gateway"
)

var client *gateway.Session

func main() {
	fmt.Println("Testing...")

	client = goscord.New(&gateway.Options{
		Token:   "",
		Intents: gateway.IntentGuilds + gateway.IntentGuildMessages,
	})

	_ = client.On("ready", OnReady)
	_ = client.On("messageCreate", OnMessageCreate)

	if err := client.Login(); err != nil {
		panic(err)
	}

	select {}
}

func OnReady() {
	fmt.Println("Logged in as " + client.Me().Tag())
	_ = client.SetActivity(&discord.Activity{Name: fmt.Sprintf("%d servers", len(client.State().Guilds())), Type: discord.ActivityWatching})
	_ = client.SetStatus("idle")
}

func OnMessageCreate(msg *discord.Message) {
	if strings.ToLower(msg.Content) == "ping" {
		_, _ = client.Channel.Send(msg.ChannelId, "Pong!")
	}

	if strings.ToLower(msg.Content) == "embed" {
		embed := embed.NewEmbedBuilder()
		embed.SetAuthor("Testing", "")
		embed.SetDescription("Just testing some things lol")
		_, _ = client.Channel.Send(msg.ChannelId, embed)
	}

	if strings.ToLower(msg.Content) == "dogeimage" {
		dogeImg, err := os.Open("doge.png")
		if err != nil {
			panic(err)
		}

		defer dogeImg.Close()

		_, _ = client.Channel.Send(msg.ChannelId, dogeImg)
	}
}
