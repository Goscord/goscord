package main

import (
	"fmt"
	"github.com/Goscord/goscord"
	"github.com/Goscord/goscord/gateway"
)

var client *gateway.Session

func main() {
	fmt.Println("Testing...")

	client = goscord.New(&gateway.Options{
		Token: "",
	})

	_ = client.On("ready", OnReady)

	if err := client.Login(); err != nil {
		panic(err)
	}

	select {}
}

func OnReady() {
	fmt.Println("Logged in as " + client.User().Tag())
	_ = client.SetStatus("idle")
}