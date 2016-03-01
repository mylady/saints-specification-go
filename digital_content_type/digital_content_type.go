package digital_content_type

const (
	RealVideo = 0
	VideoFile = 1
	AudioFile = 2
	WebPage   = 3
	ImageFile = 4
)

var contentMap = map[uint]uint{
	RealVideo: 2048,
	VideoFile: 2049,
	AudioFile: 2050,
	WebPage:   2051,
	ImageFile: 2052,
}

func GetContentMap(content_type uint) (resource_code uint) {
	if v, ok := contentMap[content_type]; ok {
		return v
	} else {
		return -1
	}
}
