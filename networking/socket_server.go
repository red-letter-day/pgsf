package networking

import (
	"fmt"
	"github.com/asaskevich/EventBus"
	"github.com/olahol/melody"
	"github.com/red-letter-day/pgsf/router"
	"net/http"
)

type SocketServer struct {
	Melody *melody.Melody
	events *EventBus.Bus
	router *router.Router
}

func NewSocketServer(events *EventBus.Bus, router *router.Router) SocketServer {
	server := SocketServer{
		Melody: melody.New(),
		events: events,
		router: router,
	}

	// Adds the default handlers.
	server.initialize()

	return server
}

// Initialize is called internally to add the default handlers.
func (server *SocketServer) initialize() {
	//events := *server.events
	server.Melody.HandleConnect(func(sender *melody.Session) {
		//events.Publish(constants.OnConnect, sender)
		// TODO: Add .OnConnect
	})

	server.Melody.HandleDisconnect(func(sender *melody.Session) {
		//events.Publish(constants.OnDisonnect, sender)
		// TODO: Add .OnDisconnect
	})

	server.Melody.HandleMessage(func(sender *melody.Session, message []byte) {
		//events.Publish(constants.OnInboundNetworkMessage, messages.NewNetworkMessage(sender, message))
		err := server.router.ProcessByteMessage(sender, message)
		if err != nil {
			fmt.Println("error processing message:", err)
		}
	})
}

func (server *SocketServer) WebsocketHandler(w http.ResponseWriter, r *http.Request) error {
	return server.Melody.HandleRequest(w, r)
}
