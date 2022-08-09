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
	_ = client.On("guildMemberAdd", OnGuildMemberAdd)

	if err := client.Login(); err != nil {
		panic(err)
	}

	select {}
}

func OnReady() {
	fmt.Println("Logged in as " + client.Me().Tag())

	_ = client.SetActivity(&discord.Activity{Name: "Goscord devs working on the lib rn", Type: discord.ActivityWatching})
	_ = client.SetStatus("idle")
}

func OnGuildMemberAdd(*discord.GuildMember) {
	if c, ok := client.State().Channel("1001943782016688292"); ok == nil {
		client.Channel.Send(c.Id, "aram sur lol?")
	} else {
		fmt.Println("Could not find channel")
	}
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
