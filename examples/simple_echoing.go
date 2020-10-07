package examples

import (
	"github.com/red-letter-day/pgsf/constants"
	"github.com/red-letter-day/pgsf/messages"
	"github.com/red-letter-day/pgsf/pgsf"
)

type ChatMessage struct {
	Content  string
	Receiver uint64
}

func main() {
	chatMessage := ChatMessage{
		Content:  "hello from chat message",
		Receiver: 1,
	}

	sampleInboundMessage := messages.NewNetworkMessage(0, chatMessage)

	configuration := pgsf.Configuration{
		Name:       "Example Chat Server",
		MaxClients: 10,
		Port:       8000,
		ListenUrl:  "/game",
	}

	server := pgsf.NewPgsfInstance(configuration)
	server.Events.Publish(constants.OnInboundNetworkMessage, sampleInboundMessage)
	go server.StartServer()
}
