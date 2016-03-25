package saints_go_specification

type DigitalCommand struct {
	CommandType uint
	Layout      DigitalLayout
	Marquee     DigitalMarquee
	Speech      string
	Value       string
	ElapseTime  uint
}
