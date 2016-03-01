package saints_go_specification

type DigitalCommand struct {
	CommandType uint
	Layout      DigitalLayout
	Text        DigitalMarquee
	Speech      string
	ElapsedTime uint
}
