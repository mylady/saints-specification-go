package code

//ResourceCode
const (
	Camera     = 0
	DoorPoint  = 1
	Door       = 2
	Alarm      = 3
	Call       = 4
	CallHost   = 5
	PowerGrid  = 6
	EnvMonitor = 7
	RFIDTag    = 8
	Relay      = 9
	LED        = 10

	DecoderChannel      = 128
	MatrixChannel       = 129
	TVWallChannel       = 130
	VideoAnalyseChannel = 131

	VM        = 256
	DM        = 257
	Matrix    = 258
	TVWall    = 259
	AlarmHost = 260

	PowerGridHost  = 261
	EnvMonitorHost = 262
	UPS            = 263
	VideoAnalyse   = 264
	RelayHost      = 265
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
