package spec

import "encoding/json"

// MessageType
const (
	MessageTypeUnknown       = -1
	MessageTypeDeviceEvent   = 0
	MessageTypeDeviceStatus  = 1
	MessageTypeDeviceFault   = 2
	MessageTypeDeviceControl = 3
	MessageTypeCustom        = 99
	MessageTypeOthers        = 100
)

// MessageTopic
const (
	TopicDeviceEvent   = "saints-device-event"
	TopicDeviceStatus  = "saints-device-status"
	TopicDeviceFault   = "saints-device-fault"
	TopicDeviceControl = "saints-device-control"
	TopicCustom        = "saints-custom"
	TopicOthers        = "others"
)

// Message :message struct for external data exchange
type Message struct {
	Type    int             `json:"type"`
	Content json.RawMessage `json:"content"`
	Sender  string          `json:"sender"`
}
