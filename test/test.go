package main

import (
	"fmt"
	"strings"

	"github.com/Goscord/goscord"
	"github.com/Goscord/goscord/discord"
	"github.com/Goscord/goscord/gateway"
)

var client *gateway.Session

func main() {
	fmt.Println("Testing...")

	client = goscord.New(&gateway.Options{
		Token: "ODMxNTgzNTY2ODg2MjA3NTE4.G_OA7n.IYmWDrcgq_6sBv0CPr0rfHYX3ZjNsOmlgMMK_M",
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
	_ = client.SetStatus("idle")
}

func OnMessageCreate(msg *discord.Message) {
	if strings.ToLower(msg.Content) == "ping" {
		_, _ = client.Channel.Send(msg.ChannelId, "Pong!")
	}
}
