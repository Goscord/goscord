package main

import (
	"fmt"
	"github.com/Goscord/goscord/goscord"
	"github.com/Goscord/goscord/goscord/discord"
	"github.com/Goscord/goscord/goscord/gateway"
	"strings"
)

var client *gateway.Session

func main() {
	fmt.Println("Testing...")

	client = goscord.New(&gateway.Options{
		Token:   "",
		Intents: gateway.IntentGuilds,
	})

	client.On("ready", OnReady)
	client.On("guildMemberAdd", OnGuildMemberAdd)
	client.On("interactionCreate", OnInteractionCreate)

	if err := client.Login(); err != nil {
		panic(err)
	}

	select {}
}

func OnReady() {
	fmt.Println("Logged in as " + client.Me().Tag())

	appCmd := &discord.ApplicationCommand{
		Name:        "ping",
		Type:        discord.ApplicationCommandChat,
		Description: "Pong pong pong!",
	}

	_, _ = client.Application.RegisterCommand(client.Me().Id, "", appCmd)

	_ = client.SetActivity(&discord.Activity{Name: "Goscord's devs working on the lib rn", Type: discord.ActivityWatching})
	_ = client.SetStatus(discord.StatusTypeIdle)
}

func OnInteractionCreate(i *discord.Interaction) {
	if i.Type == discord.InteractionTypeMessageComponent {
		data := i.MessageComponentData()

		if data.CustomId == "test_btn_pong" {
			client.Interaction.CreateResponse(i.Id, i.Token, &discord.InteractionCallbackMessage{Content: "Pong!"})
		}

		if data.CustomId == "test_text" {
			client.Interaction.CreateResponse(i.Id, i.Token, &discord.InteractionCallbackMessage{Content: fmt.Sprintf("You choose : %s", strings.Join(data.Values, ", "))})
		}
	}

	if i.Type == discord.InteractionTypeApplicationCommand {
		if i.ApplicationCommandData().Name == "ping" {
			if err := client.Interaction.CreateResponse(i.Id, i.Token, &discord.InteractionCallbackMessage{
				Content: "Pong!",
				Components: []discord.MessageComponent{
					&discord.ActionRows{
						Components: []discord.MessageComponent{
							discord.Button{
								CustomId: "test_btn_pong",
								Style:    discord.ButtonStyleSuccess,
								Label:    "Pong",
							},
						},
					},
					&discord.ActionRows{
						Components: []discord.MessageComponent{
							discord.SelectMenu{
								CustomId:    "test_text",
								PlaceHolder: "feur",
								MinValues:   1,
								MaxValues:   1,
								Options: []*discord.SelectOption{
									{
										Description: "lmao u good",
										Label:       "Yes",
										Value:       "yes",
									},
									{
										Description: "lmao u bad",
										Label:       "No",
										Value:       "No",
									},
								},
							},
						},
					},
				},
			}); err != nil {
				fmt.Println(err)
			}
		}
	}
}

func OnGuildMemberAdd(a *discord.GuildMember) {
	if c, ok := client.State().Channel("1001943782016688292"); ok == nil {
		client.Channel.SendMessage(c.Id, "Welcome to the server <@"+a.User.Id+"> !")
	} else {
		fmt.Println("Could not find channel")
	}
}
