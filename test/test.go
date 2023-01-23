package main

import (
	"fmt"
	"github.com/Goscord/goscord/goscord"
	"github.com/Goscord/goscord/goscord/discord"
	"github.com/Goscord/goscord/goscord/gateway"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// Create client instance
	client := goscord.New(&gateway.Options{
		Token:   "ODMxNTgzNTY2ODg2MjA3NTE4.G__-m6.WExLly-pqc6Bj0-Y5dRT4ASn0PMVgmFGxVCSBs",
		Intents: gateway.IntentsNonPrivileged,
	})

	// Load events
	_ = client.On("ready", OnReady(client))
	_ = client.On("interactionCreate", CommandHandler(client))

	// login client
	if err := client.Login(); err != nil {
		panic(err)
	}

	// Wait here until term signal
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Cleanly close down the Discord session
	client.Close()
}

func OnReady(client *gateway.Session) func() {
	return func() {
		fmt.Println("Logged in as ", client.Me().Tag())

		// Register slash commands
		appCmd := &discord.ApplicationCommand{
			Name:        "test",
			Type:        discord.ApplicationCommandChat,
			Description: "test command",
			Options:     make([]*discord.ApplicationCommandOption, 0),
		}
		_, _ = client.Application.RegisterCommand(client.Me().Id, "", appCmd)
	}
}

func CommandHandler(client *gateway.Session) func(*discord.Interaction) {
	return func(interaction *discord.Interaction) {
		if interaction.Member == nil {
			return
		}

		// Check if the command is "test"
		if interaction.Data.(discord.ApplicationCommandData).Name != "test" {
			return
		}

		_, err := client.JoinVoiceChannel("1001943780766797885", "1062789982617608313", false, false)
		if err != nil {
			fmt.Println(err)
		}

		client.Interaction.CreateResponse(interaction.Id, interaction.Token, ":+1:")
	}
}
