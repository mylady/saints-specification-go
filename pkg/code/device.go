package code

//ResourceCode
const (
	RCCamera     = 0
	RCDoorPoint  = 1
	RCDoor       = 2
	RCAlarm      = 3
	RCCall       = 4
	RCCallHost   = 5
	RCPowerGrid  = 6
	RCEnvMonitor = 7
	RCRFIDTag    = 8
	RCRelay      = 9
	RCLED        = 10

	RCDecoderChannel      = 128
	RCMatrixChannel       = 129
	RCTVWallChannel       = 130
	RCVideoAnalyseChannel = 131

	RCVM        = 256
	RCDM        = 257
	RCMatrix    = 258
	RCTVWall    = 259
	RCAlarmHost = 260

	RCPowerGridHost  = 261
	RCEnvMonitorHost = 262
	RCUPS            = 263
	RCVideoAnalyse   = 264
	RCRelayHost      = 265
)

//AlarmDeviceType
const (
	AlarmTypeManuAlarm     = 0
	AlarmTypeIrAlarm       = 1
	AlarmTypeElecPerimeter = 2
	AlarmTypeSoundAlarm    = 3
	AlarmTypeSmokeAlarm    = 4
)

//CameraDeviceType
const (
	CameraTypeBall = 0
	CameraTypeGun  = 1
)

//VideoAnalyseChannelType
const (
	AnalyseTypeFace         = 0
	AnalyseTypeBehavior     = 1
	AnalyseTypeVideoQuality = 2
)

//DeviceStatus
const (
	StatusNormal     = 0
	StatusArm        = 1
	StatusDisarm     = 2
	StatusOpen       = 3
	StatusClose      = 4
	StatusDisconnect = 5
	StatusConnect    = 6
	StatusAlarm      = 7
	StatusFault      = 8
	StatusUnknown    = 9
)
