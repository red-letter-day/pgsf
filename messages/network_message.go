package messages

import (
	"github.com/olahol/melody"
)

// NetworkMessage contains the sender id, along with type information and the payload itself.
type NetworkMessage struct {
	Sender  *melody.Session
	Payload []byte
}

// Creates a new inbound message, by reflecting on the payload type.
func NewNetworkMessage(sender *melody.Session, payload []byte) NetworkMessage {
	return NetworkMessage{
		Sender:  sender,
		Payload: payload,
	}
}
