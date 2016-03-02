package resource_code

const (
	//front-end of security device
	Camera            = 0
	Door              = 1
	ManuAlarm         = 2
	CallTerminal      = 3
	SoundAlarm        = 4
	IrAlarm           = 5
	PowerGrid         = 6
	EnviromentMonitor = 7

	//front-end virtual security device
	DecodeChannel       = 256
	MatrixChannel       = 257
	TvwallChannel       = 258
	VideoAnalyseChannel = 259

	//back-end of security device
	EncoderHost           = 512
	DecoderHost           = 513
	MatrixController      = 514
	TvwallController      = 515
	AlarmHost             = 516
	DoorController        = 517
	CallHost              = 518
	ListenerHost          = 519
	PowerGridHost         = 520
	LedController         = 521
	EnviromentMonitorHost = 522
	UpsHost               = 523
	VideoAnalyseHost      = 524
	Detector              = 525

	//other device
	PersonalComputer = 1024
	Server           = 1025
	Laptop           = 1026
	Cellphone        = 1027
	Pad              = 1028
	Switcher         = 1029
	Router           = 1030
	Hub              = 1031
	InfopubTerminal  = 1032
	IptvTerminal     = 1033

	//non-device
	RealVideoSource = 2048
	VideoFileSource = 2049
	AudioFileSource = 2050
	WebpageSource   = 2051
	ImageFileSource = 2052
)
