package networking_test

import (
	"github.com/red-letter-day/pgsf/constants"
	"github.com/red-letter-day/pgsf/messages"
	"github.com/red-letter-day/pgsf/pgsf"
	"github.com/stretchr/testify/assert"
	"testing"
)

type ChatMessage struct {
	Content  string
	Receiver uint64
}

func TestNewNetworkEngine(t *testing.T) {
	configuration := pgsf.Configuration{
		Name:       "EngineTest",
		MaxClients: 10,
		Port:       8000,
		ListenUrl:  "/game",
	}

	server := pgsf.NewPgsfInstance(configuration)

	assert.NotNil(t, server, "expected non nil pgsf instance")
	assert.NotNil(t, server.NetworkEngine, "expected non nil network engine")
	assert.NotNil(t, server.Events, "expected non nil event bus")
}

func TestNetworkEngineHandleMessage(t *testing.T) {

	chatMessage := ChatMessage{
		Content:  "hello from chat message",
		Receiver: 1,
	}

	sampleInboundMessage := messages.NewNetworkMessage(0, chatMessage)

	configuration := pgsf.Configuration{
		Name:       "ChatServer1",
		MaxClients: 10,
		Port:       8000,
		ListenUrl:  "/game",
	}

	server := pgsf.NewPgsfInstance(configuration)
	server.Events.Publish(constants.OnInboundNetworkMessage, sampleInboundMessage)

	go server.StartServer()
}
