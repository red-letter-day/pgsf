package messages

import "encoding/json"

type InboundMessage struct {
	Name string      `json:"name"`
	Data interface{} `json:"data"`
}

// NewInboundMessage takes a byte array and attempts to unmarshal it into the InboundMessage struct.
func NewInboundMessage(data []byte) (InboundMessage, error) {
	var message InboundMessage
	err := json.Unmarshal(data, &message)
	return message, err
}
