package main

import (
	"fmt"
	"github.com/olahol/melody"
	"github.com/red-letter-day/pgsf/pgsf"
)

type ChatMessage struct {
	Content  string `json:"content"`
	Receiver string `json:"receiver"`
}

func main() {

	configuration := pgsf.Configuration{
		Name:       "Example Chat Server",
		MaxClients: 10,
		Port:       8081,
		ListenUrl:  "/game",
	}

	server := pgsf.NewPgsfInstance(configuration)

	r := server.Router

	// TODO: Add chan support so we can run these async?
	r.On("opc", func(sender *melody.Session, data *ChatMessage) {
		chatMessage := *data
		fmt.Println("Receive chatmessage:", chatMessage.Content, "and", chatMessage.Receiver)
	})

	r.OnConnect(func(sender *melody.Session) {
		fmt.Println("User connected: ", )
		_ = sender.Write([]byte("welcome"))
	})

	r.OnDisconnect(func(sender *melody.Session, message string) {
		disconnectMessage, _ := sender.Get("dcmsg")
		fmt.Println("User disconnected: ", disconnectMessage)
		//_ = sender.Write([]byte("bye!"))
	})

	server.StartServer()
}
