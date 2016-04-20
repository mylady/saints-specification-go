package digital_terminal_type

const (
	InfopubTerminal = 0
	IptvTerminal    = 1
	LedScreen       = 2
)

var terminalMap = map[uint]uint{
	InfopubTerminal: 1032,
	IptvTerminal:    1033,
	LedScreen:       521,
}

func GetTemrinalMap(terminal_type uint) (resource_code uint) {
	if v, ok := terminalMap[terminal_type]; ok {
		return v
	} else {
		return 65535
	}
}
