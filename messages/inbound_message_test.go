package messages_test

import (
	"github.com/red-letter-day/pgsf/messages"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewInboundMessage(t *testing.T) {
	inputData := `{"name":"opc", "data":{"content":"ssss","receiver": 12242}}`

	message, err := messages.NewInboundMessage([]byte(inputData))

	assert.Nil(t, err, "expected successful marshaling")
	assert.NotNil(t, message, "expected non nil message")
}
