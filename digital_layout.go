package saints_go_specification

type DigitalLayout struct {
	Width           uint
	Height          uint
	FullScreen      bool
	BackgroundImage string
	BackgroundAudio string
	Areas           []DigitalLayoutArea
}
