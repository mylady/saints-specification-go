package saints_specification_go

import (
	"fmt"
	"time"

	"github.com/nu7hatch/gouuid"
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

func NewMessage(content string, messageType int, receivers []string) (message TransmitMessage, err error) {
	uid, _ := uuid.NewV4()
	message = TransmitMessage{
		MessageId:       uid.String(),
		MessageContent:  content,
		MessageType:     uint(messageType),
		MessageTime:     JsonTime(time.Now()),
		MessageReceiver: receivers,
	}
	return
}
