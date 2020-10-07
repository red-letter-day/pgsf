package networking

import (
	"fmt"
	"github.com/asaskevich/EventBus"
	"github.com/red-letter-day/pgsf/constants"
	"github.com/red-letter-day/pgsf/messages"
)

// NetworkEngine handles incoming network messages.
type NetworkEngine struct {
	events *EventBus.Bus
}

// Initializes the Network Engine and subscribes to the inbound message topic.
func NewNetworkEngine(eventBus *EventBus.Bus) NetworkEngine {
	networkEngine := NetworkEngine{
		events: eventBus,
	}
	events := *networkEngine.events

	// We can ignore the error because it only gets triggered if the 2nd argument is not a function.
	_ = events.Subscribe(constants.OnInboundNetworkMessage, networkEngine.handleMessage)

	return networkEngine
}

// handleMessage handles an incoming message.
func (nw *NetworkEngine) handleMessage(message messages.NetworkMessage) {
	fmt.Println("incoming message:", message)
}
