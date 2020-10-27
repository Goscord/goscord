package main

import (
	"fmt"

	"github.com/Seyz123/yalis"
)

var client *yalis.Client

func main() {
	fmt.Println("Testing...")

	client = yalis.NewClient("NzM1NjQyNjE2NDc3MjUzNjg0.XxjOkw.DxpP72dLDdLbJ6IqE2OvV-zX7-k")

	_ = client.On("ready", OnReady)

	if err := client.Login(); err != nil {
		panic(err)
	}

	select {}
}

func OnReady() {
	fmt.Println("Bot online !")
}
