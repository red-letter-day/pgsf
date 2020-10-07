package pgsf

import (
	"fmt"
	"github.com/asaskevich/EventBus"
	"github.com/common-nighthawk/go-figure"
	"github.com/red-letter-day/pgsf/constants"
	"github.com/red-letter-day/pgsf/networking"
	"log"
	"net/http"
)

// PgsfInstance is the central point from which the library runs.
// Among other things, this struct contains the websocket server itself, and routes messages to the respective handlers.
type ServerInstance struct {
	configuration Configuration

	// NetworkEngine
	NetworkEngine networking.NetworkEngine

	// Events contains the Eventbus we use to deliver messages throughout the serverside codebase.
	Events EventBus.Bus

	// The initialized bool informs the application whether the server has been started.
	started bool
}

// StartServer initiates the actual websocket server.
func (instance *ServerInstance) StartServer() {
	// Publish a start message so any listeners can run code as needed.
	instance.Events.Publish(constants.OnStartSocketServer)

	router := http.NewServeMux()
	//router.Handle(instance.configuration.ListenUrl, nil)
	log.Fatal(http.ListenAndServe(":8080", router))
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

	fmt.Println(configuration.Name, "starting on port:", configuration.Port)
	fmt.Println("---------------------------------------------")

	instance := ServerInstance{
		configuration: configuration,
		Events:        EventBus.New(),
	}

	// Create the network engine and pass in the EventBus reference.
	instance.NetworkEngine = networking.NewNetworkEngine(&instance.Events)

	return instance
}
