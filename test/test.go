package main

import (
	"fmt"
	"github.com/Goscord/goscord"
	"github.com/Goscord/goscord/discord"
	"github.com/Goscord/goscord/gateway"
	"strings"
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
	fmt.Println("Logged in as " + client.Me().Tag())
	_ = client.SetStatus("idle")
}

func OnMessage(msg *discord.Message) {
	if strings.ToLower(msg.Content) == "ping" {
		_, _ = client.Channel.Send(msg.ChannelId, "Pong!")
	}
}
