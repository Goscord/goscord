package main

import (
	"fmt"
	"github.com/Seyz123/yalis"
)

type Message struct {
	Content string `json:"content,omitempty"`
}

var client *yalis.Client

func main() {
	fmt.Println("Testing...")
	
	client = yalis.NewClient("NzM1NjQyNjE2NDc3MjUzNjg0.XxjOkw.DxpP72dLDdLbJ6IqE2OvV-zX7-k")
	
	client.On("ready", OnReady)

	if err := client.Login(); err != nil {
		panic(err)
	}

	<-make(chan bool)
}

func OnReady() {
    fmt.Println("Bot online !")
}