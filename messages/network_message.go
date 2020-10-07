package messages

import (
	"github.com/red-letter-day/pgsf/payload"
	"reflect"
)

// NetworkMessage contains the sender id, along with type information and the payload itself.
type NetworkMessage struct {
	Sender  uint64
	Type    reflect.Type
	Payload payload.PayloadInterface
}

// Creates a new inbound message, by reflecting on the payload type.
func NewNetworkMessage(sender uint64, payload payload.PayloadInterface) NetworkMessage {
	return NetworkMessage{
		Sender: sender,
		Type:   reflect.TypeOf(payload),

		Payload: payload,
	}
}
