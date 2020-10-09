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
// Adds handlers for connection, disconnection and general messages.
func (server *SocketServer) initialize() {

	server.Melody.HandleConnect(func(sender *melody.Session) {
		server.router.ProcessConnectMessage(sender)
	})

	server.Melody.HandleDisconnect(func(sender *melody.Session) {
		server.router.ProcessDisconnectMessage(sender)
	})

	server.Melody.HandleMessage(func(sender *melody.Session, message []byte) {
		err := server.router.ProcessByteMessage(sender, message)
		if err != nil {
			fmt.Println("error processing message:", err)
		}
	})
}

// Websocket handler
func (server *SocketServer) WebsocketHandler(w http.ResponseWriter, r *http.Request) error {
	return server.Melody.HandleRequest(w, r)
}
