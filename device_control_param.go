package saints_go_specification

type DeviceControlParam struct {
	Cmd        uint
	LibCode    uint
	Address    string
	Parameters []string
}