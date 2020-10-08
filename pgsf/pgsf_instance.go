package pgsf

import (
	"fmt"
	"github.com/asaskevich/EventBus"
	"github.com/common-nighthawk/go-figure"
	"github.com/gin-gonic/gin"
	"github.com/red-letter-day/pgsf/constants"
	"github.com/red-letter-day/pgsf/networking"
	"github.com/red-letter-day/pgsf/router"
)

// PgsfInstance is the central point from which the library runs.
// Among other things, this struct contains the websocket server itself, and routes messages to the respective handlers.
type ServerInstance struct {
	configuration Configuration

	SocketServer networking.SocketServer

	// Events contains the Eventbus we use to deliver messages throughout the serverside codebase.
	Events EventBus.Bus

	Router *router.Router

	// The initialized bool informs the application whether the server has been started.
	started bool
}

// StartServer initiates the actual websocket server.
func (instance *ServerInstance) StartServer() {

	// Initialize the socket server itself, and then bind it to the router.
	instance.SocketServer = networking.NewSocketServer(&instance.Events, instance.Router)

	gin.SetMode(gin.ReleaseMode)
	webRouter := gin.Default()
	webRouter.GET(instance.configuration.ListenUrl, func(c *gin.Context) {
		err := instance.SocketServer.Melody.HandleRequest(c.Writer, c.Request)

		if err != nil {
			fmt.Println("Error handling WS request:", err)
		}
	})

	err := webRouter.Run(fmt.Sprintf(":%d", instance.configuration.Port))

	if err != nil {
		fmt.Println("Error starting pgsf websocket server:", err)
	}

	fmt.Println(instance.configuration.Name, "starting on port:", instance.configuration.Port)

	// Publish a start message so any listeners can run code as needed.
	instance.Events.Publish(constants.OnStartSocketServer)

	// Mark the instance as started.
	instance.started = true
}

// Returns the PgsfInstance's configuration.
func (instance *ServerInstance) GetConfiguration() Configuration {
	return instance.configuration
}

// Initializes a new PgsfInstance with a specified configuration struct.
// Also creates a new EventBus (github.com/asaskevich/Eventbus).
func NewPgsfInstance(configuration Configuration) ServerInstance {

	fig := figure.NewFigure(configuration.Name, "standard", true)
	fig.Print()

	instance := ServerInstance{
		configuration: configuration,
		Events:        EventBus.New(),
		Router:        router.NewRouter(),
	}

	return instance
}
