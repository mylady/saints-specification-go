package resource_code

const (
	Unknown = -1
	//front-end of security device
	Camera             = 0
	Door               = 1
	ManuAlarm          = 2
	CallTerminal       = 3
	SoundAlarm         = 4
	IrAlarm            = 5
	PowerGrid          = 6
	TemperatureMonitor = 7
	SmokeMonitor       = 8
	AirMonitor         = 9

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

var ResourceCodeMap = make(map[int]string)

func init() {
	//unknown
	ResourceCodeMap[Unknown] = "未知类型"
	// front-end
	// 0,1,2,3,4,5,6,7,
	ResourceCodeMap[Camera] = "摄像头"
	ResourceCodeMap[Door] = "门禁点"
	ResourceCodeMap[ManuAlarm] = "手动报警"
	ResourceCodeMap[CallTerminal] = "对讲终端"
	ResourceCodeMap[SoundAlarm] = "分贝报警"
	ResourceCodeMap[IrAlarm] = "红外报警"
	ResourceCodeMap[PowerGrid] = "电网网段"
	ResourceCodeMap[EnviromentMonitor] = "环境监测点"
	// font-virtual
	//	256, 257, 258, 259,
	ResourceCodeMap[DecodeChannel] = "解码通道"
	ResourceCodeMap[MatrixChannel] = "矩阵通道"
	ResourceCodeMap[TvwallChannel] = "显示屏输出"
	ResourceCodeMap[VideoAnalyseChannel] = "视频分析通道"
	//back-end
	//	512, 513, 514, 515, 516, 517, 518, 519, 520, 521, 522, 523, 524, 525,
	ResourceCodeMap[EncoderHost] = "编码器主机"
	ResourceCodeMap[DecoderHost] = "解码器主机"
	ResourceCodeMap[MatrixController] = "矩阵控制器"
	ResourceCodeMap[TvwallController] = "大屏控制器"
	ResourceCodeMap[AlarmHost] = "报警主机"
	ResourceCodeMap[DoorController] = "门禁控制器"
	ResourceCodeMap[CallHost] = "对讲主机"
	ResourceCodeMap[ListenerHost] = "监听主机"
	ResourceCodeMap[PowerGridHost] = "电网主机"
	ResourceCodeMap[LedController] = "电子屏控制器"
	ResourceCodeMap[EnviromentMonitorHost] = "环境监测主机"
	ResourceCodeMap[UpsHost] = "不间断电源"
	ResourceCodeMap[VideoAnalyseHost] = "视频分析主机"
	ResourceCodeMap[Detector] = "探针"
	//	other
	//	1024, 1025, 1026, 1027, 1028, 1029, 1030, 1031, 1032, 1033,
	ResourceCodeMap[PersonalComputer] = "个人电脑"
	ResourceCodeMap[Server] = "服务器"
	ResourceCodeMap[Laptop] = "笔记本电脑"
	ResourceCodeMap[Cellphone] = "手机"
	ResourceCodeMap[Pad] = "平板电脑"
	ResourceCodeMap[Switcher] = "交换机"
	ResourceCodeMap[Router] = "路由器"
	ResourceCodeMap[Hub] = "集线器"
	ResourceCodeMap[InfopubTerminal] = "信息发布终端"
	ResourceCodeMap[IptvTerminal] = "数字电视终端"
	//	non-device
	//	2048, 2049, 2050, 2051, 2052,
	ResourceCodeMap[RealVideoSource] = "直播源"
	ResourceCodeMap[VideoFileSource] = "视频文件"
	ResourceCodeMap[AudioFileSource] = "音频文件"
	ResourceCodeMap[WebpageSource] = "网页"
	ResourceCodeMap[ImageFileSource] = "图片文件"
}
