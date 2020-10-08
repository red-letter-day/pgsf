package main

import (
	"fmt"
	"github.com/olahol/melody"
	"github.com/red-letter-day/pgsf/pgsf"
)

type ChatMessage struct {
	Content  string `json:"content"`
	Receiver int `json:"receiver"`
}

func main() {

	configuration := pgsf.Configuration{
		Name:       "Example Chat Server",
		MaxClients: 10,
		Port:       8081,
		ListenUrl:  "/game",
	}

	server := pgsf.NewPgsfInstance(configuration)

	server.Router.On("opc", func(sender *melody.Session, data *ChatMessage) {
		chatMessage := *data
		fmt.Println("Receive chatmessage:", chatMessage.Content, "and", chatMessage.Receiver)
	})

	server.StartServer()
}
