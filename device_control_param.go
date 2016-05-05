package saints_specification_go

type DeviceControlParam struct {
	ControlCode   uint
	InterfaceCode uint
	Address       string
	Sender        string
	Parameter     []string
}
