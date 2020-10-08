package router

import (
	"fmt"
	"github.com/olahol/melody"
	"github.com/red-letter-day/pgsf/messages"
	"reflect"
	"sync"
)

type Router struct {
	callbacks map[string]func(*melody.Session, interface{})
	lock *sync.Mutex
}

func NewRouter() *Router {
	fmt.Println("init router")
	return &Router{
		callbacks: make(map[string]func(*melody.Session, interface{})),
		lock: new(sync.Mutex),
	}
}

// On is used to add a message listener for a network message type.
//
// For instance:
//
// 		router.On("hello", func(connection *melody.Session, data *Hello) {
// 		})
//
func (r *Router) On(name string, callback interface{}) {
	callbackValue := reflect.ValueOf(callback)
	callbackType := reflect.TypeOf(callback)

	callbackDataElement := callbackType.In(1).Elem()
	r.callbacks[name] = func(conn *melody.Session, data interface{}) {
		result := reflect.New(callbackDataElement)
		arguments := []reflect.Value{reflect.ValueOf(conn), result}
		callbackValue.Call(arguments)
	}
}

func (r *Router) ProcessByteMessage(connection *melody.Session, data []byte) error {
	r.lock.Lock()

	inboundMessage, err := messages.NewInboundMessage(data)

	if err != nil {
		return err
	}

	fmt.Println("inbound message string:", inboundMessage)

	if callback, ok := r.callbacks[inboundMessage.Name]; ok {
		fmt.Println("calling callback with inbound name:", inboundMessage.Name)
		callback(connection, inboundMessage.Data)
	}

	r.lock.Unlock()
	return nil
}
