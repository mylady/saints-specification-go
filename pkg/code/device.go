package code

//ResourceCode
const (
	Camera     = 0
	Door       = 1
	Alarm      = 2
	Call       = 3
	PowerGrid  = 4
	EnvMonitor = 5
	RFIDTag    = 6
	Relay      = 7
	LED        = 8

	DecoderChannel      = 128
	MatrixChannel       = 129
	TVWallChannel       = 130
	VideoAnalyseChannel = 131

	VM             = 256
	DM             = 257
	Matrix         = 258
	TVWall         = 259
	AlarmHost      = 260
	DoorController = 261
	CallHost       = 262
	PowerGridHost  = 263
	EnvMonitorHost = 264
	UPS            = 265
	VideoAnalyse   = 266
	RelayHost      = 267
)

//AlarmDeviceType
const (
	ManuAlarm     = 0
	IrAlarm       = 1
	ElecPerimeter = 2
	SoundAlarm    = 3
	SmokeAlarm    = 4
)

//CameraDeviceType
const (
	Ball = 0
	Gun  = 1
)
