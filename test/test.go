package main

import (
	"fmt"
	"github.com/Seyz123/yalis"
)

type Message struct {
	Content string `json:"content,omitempty"`
}

func main() {
	fmt.Println("Testing...")
	client := yalis.NewClient("NzM1NjQyNjE2NDc3MjUzNjg0.XxjOkw.DxpP72dLDdLbJ6IqE2OvV-zX7-k")

	/*
	for i := 0; i < 10; i++ {
		data, err := client.RestClient().Request(fmt.Sprintf("/channels/%s/messages", "756640025743065139"), "POST", []byte(`{"content":"Test"}`))

		if err != nil {
			panic(err)
		}

		fmt.Println(string(data))
	}
	*/

	if err := client.Login(); err != nil {
		panic(err)
	}

	<-make(chan bool)
}