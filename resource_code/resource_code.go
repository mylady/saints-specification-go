package resource_code

const (
	Unknown = -1
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

var ResourceCodeMap = make(map[int]string)

func init() {
	//unknown
	ResourceCodeMap[-1] = "未知类型"
	// front-end
	// 0,1,2,3,4,5,6,7,
	ResourceCodeMap[0] = "摄像头"
	ResourceCodeMap[1] = "门禁点"
	ResourceCodeMap[2] = "手动报警"
	ResourceCodeMap[3] = "对讲终端"
	ResourceCodeMap[4] = "分贝报警"
	ResourceCodeMap[5] = "红外报警"
	ResourceCodeMap[6] = "电网网段"
	ResourceCodeMap[7] = "环境监测点"
	// font-virtual
	//	256, 257, 258, 259,
	ResourceCodeMap[256] = "解码通道"
	ResourceCodeMap[257] = "矩阵通道"
	ResourceCodeMap[258] = "显示屏输出"
	ResourceCodeMap[259] = "视频分析通道"
	//back-end
	//	512, 513, 514, 515, 516, 517, 518, 519, 520, 521, 522, 523, 524, 525,
	ResourceCodeMap[512] = "编码器主机"
	ResourceCodeMap[513] = "解码器主机"
	ResourceCodeMap[514] = "矩阵控制器"
	ResourceCodeMap[515] = "大屏控制器"
	ResourceCodeMap[516] = "报警主机"
	ResourceCodeMap[517] = "门禁控制器"
	ResourceCodeMap[518] = "对讲主机"
	ResourceCodeMap[519] = "监听主机"
	ResourceCodeMap[520] = "电网主机"
	ResourceCodeMap[521] = "电子屏控制器"
	ResourceCodeMap[522] = "环境监测主机"
	ResourceCodeMap[523] = "不间断电源"
	ResourceCodeMap[524] = "视频分析主机"
	ResourceCodeMap[525] = "探针"
	//	other
	//	1024, 1025, 1026, 1027, 1028, 1029, 1030, 1031, 1032, 1033,
	ResourceCodeMap[1024] = "个人电脑"
	ResourceCodeMap[1025] = "服务器"
	ResourceCodeMap[1026] = "笔记本电脑"
	ResourceCodeMap[1027] = "手机"
	ResourceCodeMap[1028] = "平板电脑"
	ResourceCodeMap[1029] = "交换机"
	ResourceCodeMap[1030] = "路由器"
	ResourceCodeMap[1031] = "集线器"
	ResourceCodeMap[1032] = "信息发布终端"
	ResourceCodeMap[1033] = "数字电视终端"
	//	non-device
	//	2048, 2049, 2050, 2051, 2052,
	ResourceCodeMap[2048] = "直播源"
	ResourceCodeMap[2049] = "视频文件"
	ResourceCodeMap[2050] = "音频文件"
	ResourceCodeMap[2051] = "网页"
	ResourceCodeMap[2052] = "图片文件"
}
