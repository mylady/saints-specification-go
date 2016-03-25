package saints_go_specification

import (
	"time"
)

type JsonTime time.Time

func (t JsonTime) MarshalJSON() ([]byte, error) {
	timestr := fmt.Sprintf("\"%s\"", time.Time(t).Format("2006-01-02T15:04:05.999Z07:00"))
	return []byte(timestr), nil
}

//通信消息
type TransmitMessage struct {
	MessageId       string   //消息唯一编号
	MessageType     int      //消息类型代码
	MessageContent  string   //具体消息结构体,根据消息类型不同结构不同
	MessageTime     JsonTime //消息发送事件
	MessageReceiver []string //消息接受者
}

type TransmitMessage struct {
	MessageId       string
	MessageType     uint
	MessageContent  string
	MessageTime     time.Time
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
const TRANSMIT_MESSAGE_ROOM_MAINTAIN = "maintain"
