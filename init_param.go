package saints_specification_go

type InitParam struct {
	IP               string
	RemotePort       uint
	LocalPort        uint
	UserName         string
	UserPwd          string
	Address          string
	ComPort          string
	BaudRate         uint
	DataBits         byte
	StopBits         byte
	Parity           byte
	ConnectionString string
	Extra            string
	ControlCmd       string
}
