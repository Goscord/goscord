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
		Token: "ODMxNTgzNTY2ODg2MjA3NTE4.GNmlGo.lJKp4NywTb0-YqFRG8X3Wsjhhtffw3Ww61eVoM",
		Intents: gateway.IntentGuilds |
			gateway.IntentGuildMembers |
			gateway.IntentDirectMessages |
			gateway.IntentGuildMessages |
			gateway.IntentMessageContent,
	})

	// Load events
	_ = client.On("ready", OnReady(client))
	_ = client.On("interactionCreate", CommandHandler(client))

	// Login client
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
			Options: []*discord.ApplicationCommandOption{
				{
					Name:        "message_id",
					Type:        discord.ApplicationCommandOptionString,
					Description: "Message ID",
					Required:    true,
				},
			},
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

		// Get message by ID
		msg, err := client.Channel.GetMessage(interaction.ChannelId, interaction.Data.(discord.ApplicationCommandData).Options[0].Value.(string))
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println(msg)

		client.Interaction.CreateResponse(interaction.Id, interaction.Token, fmt.Sprintf("Embed: %d", len(msg.Embeds)))
	}
}
