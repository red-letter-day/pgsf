package router_test

import (
	"fmt"
	"github.com/olahol/melody"
	"github.com/red-letter-day/pgsf/router"
	"github.com/stretchr/testify/assert"
	"testing"
)

type TestContainer struct {
	Message string
}

func TestRouter_ProcessByteMessage(t *testing.T) {
	processRouter := router.NewRouter()

	hasTriggeredTest1 := false
	processRouter.On("test1", func(session *melody.Session, data *TestContainer) {
		hasTriggeredTest1 = true
		fmt.Println("triggered test 1 with message", data.Message)
	})

	messageBytes := []byte(`{"name": "test1", "data": { "message": "1234" }}`)

	err := processRouter.ProcessByteMessage(nil, messageBytes)

	assert.Nil(t, err, "expected nil error from ProcessByteMessage")
	assert.True(t, hasTriggeredTest1, "expected test case to be triggered using message")
}