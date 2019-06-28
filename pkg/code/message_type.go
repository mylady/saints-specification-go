package code

//InternalMessageType
const (
	InternalDeviceStatus = 0
	InternalDeviceEvent  = 1
)

//TransmitMessageType
const (
	TransmitDeviceStatus  = 0
	TransmitDeviceEvent   = 1
	TransmitDeviceControl = 2
	TransmitCustom        = 99
	TransmitOthers        = 100
)
