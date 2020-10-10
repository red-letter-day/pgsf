package router

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/olahol/melody"
	"github.com/red-letter-day/pgsf/messages"
	"reflect"
	"sync"
)

type Router struct {
	callbacks             map[string]func(*melody.Session, interface{})
	onConnectCallbacks    []func(session *melody.Session)
	onDisconnectCallbacks []func(session *melody.Session)
	lock                  *sync.Mutex
}

func NewRouter() *Router {
	return &Router{
		callbacks: make(map[string]func(*melody.Session, interface{})),
		lock:      new(sync.Mutex),
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

	// Stores value, which is going to be
	callbackValue := reflect.ValueOf(callback)
	callbackType := reflect.TypeOf(callback)

	fmt.Printf("Adding <%s> callback for data type <%s>\n", name, callbackType)

	callbackDataElement := callbackType.In(1).Elem()

	// Add the actual callback.
	r.callbacks[name] = func(sender *melody.Session, data interface{}) {
		// Note: This method of converting json into a strongly typed implementation on the callback side
		// is shamelessly stolen from github.com/trevex/golem, which solves this pretty decently in the router.go file.

		// Becomes the type of the struct passed to .On()
		// calling .Interface() on this will allow us to unmarshal into this type.
		result := reflect.New(callbackDataElement)

		// Takes the data interface, turns it into []byte and attempts to unmarshal it into the reflected struct type.
		err := json.Unmarshal(data.([]byte), result.Interface())

		if err != nil {
			// TODO: Better error handling in .On()
			fmt.Printf("Unable to unmarshal data for type %s (%s)\n", callbackType, err)
			return
		}

		// Build our arguments for the callback itself.
		// Fills the first argument with the sender, and the second with the final unmarshaled struct.
		arguments := []reflect.Value{reflect.ValueOf(sender), result}

		// Calls the callback, delivering the data back to the .On function.
		callbackValue.Call(arguments)
	}
}

// OnDisconnect adds a callback to the onDisconnectCallbacks list, which gets called on any client disconnections.
func (r *Router) OnDisconnect(callback interface{}) {
	// Store the types of the callback.
	callbackValue := reflect.ValueOf(callback)
	callbackType := reflect.TypeOf(callback)

	fmt.Printf("Adding OnDisconnect callback for data type <%s>\n", callbackType)

	// Add the actual callback.
	r.onDisconnectCallbacks = append(r.onDisconnectCallbacks, func(sender *melody.Session) {

		// Build our arguments for the callback itself.
		// Fills the first argument with the sender, and the second with the final unmarshaled struct.
		arguments := []reflect.Value{reflect.ValueOf(sender)}

		callbackValue.Call(arguments)
	})
}

// OnDisconnect adds a callback to the onConnectCallbacks list, which gets called on any client connections.
func (r *Router) OnConnect(callback interface{}) {
	// Store the types of the callback.
	callbackValue := reflect.ValueOf(callback)
	callbackType := reflect.TypeOf(callback)

	fmt.Printf("Adding OnConnect callback for with type <%s>\n", callbackType)

	// Add the actual callback.
	r.onConnectCallbacks = append(r.onConnectCallbacks, func(sender *melody.Session) {

		// Build our arguments for the callback itself.
		// Fills the first argument with the sender, and the second with the final unmarshaled struct.
		arguments := []reflect.Value{reflect.ValueOf(sender)}

		callbackValue.Call(arguments)
	})
}

// ProcessByteMessage takes the incoming data from the websocket server and calls the callback registered using .On()
func (r *Router) ProcessByteMessage(connection *melody.Session, data []byte) error {
	inboundMessage, err := messages.NewInboundMessage(data)

	if err != nil {
		return err
	}

	// This little trick essentially forces the compiler to cast the json.RawMessage (alias of []byte) into []byte.
	// It is needed because casting directly causes a compiler error.
	var bytedMessageData []byte
	bytedMessageData = inboundMessage.Data

	// Checks that the callback exists, and if it does, attempts to
	if callback, ok := r.callbacks[inboundMessage.Name]; ok {
		callback(connection, bytedMessageData)
	} else {
		return errors.New("attempted to call non-existant callback (" + inboundMessage.Name + ")")
	}

	return nil
}

// ProcessDisconnectMessage calls all the registered connection callbacks.
func (r *Router) ProcessDisconnectMessage(connection *melody.Session) {
	r.processArrayCallbacks(connection, r.onDisconnectCallbacks)
}

// ProcessConnectionMessage calls all the registered connection callbacks.
func (r *Router) ProcessConnectMessage(connection *melody.Session) {
	r.processArrayCallbacks(connection, r.onConnectCallbacks)
}

// ProcessArrayCallbacks runs all the callbacks for the specified list of callbacks.
func (r *Router) processArrayCallbacks(connection *melody.Session, callbacks []func(session *melody.Session)) {
	for _, callback := range callbacks {
		callback(connection)
	}
}
