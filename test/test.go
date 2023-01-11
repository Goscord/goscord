package main

import (
	"fmt"
	"github.com/Goscord/goscord/goscord"
	"github.com/Goscord/goscord/goscord/discord"
	"github.com/Goscord/goscord/goscord/gateway"
)

var client *gateway.Session

func main() {
	fmt.Println("Testing...")

	client = goscord.New(&gateway.Options{
		Token:   "",
		Intents: gateway.IntentsNonPrivileged,
	})

	if err := client.Login(); err != nil {
		panic(err)
	}

	select {}
}

func OnReady() {
	fmt.Println("Logged in as " + client.Me().Tag())

	_ = client.SetActivity(&discord.Activity{Name: "Goscord's devs working on the lib rn", Type: discord.ActivityWatching})
	_ = client.SetStatus(discord.StatusTypeIdle)
}
