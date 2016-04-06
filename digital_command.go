package saints_specification_go

type DigitalCommand struct {
	CommandType uint
	Layout      DigitalLayout
	Marquee     DigitalMarquee
	Speech      string
	Value       string
	ElapseTime  uint
}
