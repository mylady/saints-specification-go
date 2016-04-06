package saints_specification_go

import (
	"fmt"
	"time"
)

type JsonTime time.Time

func (t JsonTime) MarshalJSON() ([]byte, error) {
	timestr := fmt.Sprintf("\"%s\"", time.Time(t).Format("2006-01-02T15:04:05.999Z07:00"))
	return []byte(timestr), nil
}

type TransmitMessage struct {
	MessageId       string
	MessageType     uint
	MessageContent  string
	MessageTime     JsonTime
	MessageReceiver []string
}

const TRANSMIT_MESSAGE_REGISTER_EVENT = "hello"
const TRANSMIT_MESSAGE_JOINROOM_EVENT = "join_room"
const TRANSMIT_MESSAGE_LEAVEROOM_EVENT = "leave_room"
const TRANSMIT_MESSAGE_MESSAGEUP_EVENT = "message_up"
const TRANSMIT_MESSAGE_MESSAGEDOWN_EVENT = "message_down"

const TRANSMIT_MESSAGE_ROOM_DEVICE_EVENT = "device_event"
const TRANSMIT_MESSAGE_ROOM_DEVICE_STATUS = "device_status"
const TRANSMIT_MESSAGE_ROOM_DEVICE_CONTROL = "device_control"
const TRANSMIT_MESSAGE_ROOM_DIGITAL_CONTROL = "digital_control"
const TRANSMIT_MESSAGE_ROOM_HOSPITAL = "hospital"
const TRANSMIT_MESSAGE_ROOM_MAINTAIN = "maintain"
