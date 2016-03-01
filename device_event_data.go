package saints_go_specification

import (
	"time"
)

type DeviceEventData struct {
	ResourceCode     uint
	Address          string
	EventType        uint
	EventLevel       uint
	EventCode        uint
	AlarmCode        uint
	EventTime        time.Time
	EventDescription string
	EventExtra       interface{}
}
