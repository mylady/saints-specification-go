package spec

// ResourceCode
const (
	ResourceCodeUnknown        = -1
	ResourceCodeEncoderChannel = 0
	ResourceCodeDoorPoint      = 1
	ResourceCodeDoor           = 2
	ResourceCodeAlarm          = 3
	ResourceCodeCallSlave      = 4
	ResourceCodeCallMaster     = 5
	ResourceCodePowerGrid      = 6
	ResourceCodeEnvMonitor     = 7
	ResourceCodeRFIDTag        = 8
	ResourceCodeRelayModule    = 9
	ResourceCodeLED            = 10

	ResourceCodeDecoderChannel      = 128
	ResourceCodeMatrixChannel       = 129
	ResourceCodeScreenOutput        = 130
	ResourceCodeVideoAnalyseChannel = 131

	ResourceCodeEncoder          = 256
	ResourceCodeDecoder          = 257
	ResourceCodeMatrixController = 258
	ResourceCodeTvWallController = 259
	ResourceCodeAlarmHost        = 260
	ResourceCodeDoorController   = 261
	ResourceCodePowerGridHost    = 262
	ResourceCodeEnvMonitorHost   = 263
	ResourceCodeUPSHost          = 264
	ResourceCodeVideoAnalyseHost = 265
	ResourceCodeRelayHost        = 266
	ResourceCodeVideoGBServer    = 267
)

// DeviceEventType
const (
	EventTypeUnknown = -1
	EventTypeStatus  = 0
	EventTypeNormal  = 1
	EventTypeAlarm   = 2
	EventTypeFault   = 3
)

// DeviceEventCode
const (
	EventCodeUnknown          = -1
	EventCodeFlashOpen        = 0
	EventCodeButtonOpen       = 1
	EventCodeListen           = 2
	EventCodeCallin           = 3
	EventCodeFaceCapture      = 4
	EventCodeCarPlateIdentify = 5
	EventCodeFingerPrintOpen  = 6
	EventCodeFaceOpen         = 7
)

// DeviceAlarmCode
const (
	AlarmCodeUnknown             = -1
	AlarmCodeAreaAlarm           = 0
	AlarmCodeAreaIntrusion       = 1
	AlarmCodeLimitHeigt          = 2
	AlarmCodeWander              = 3
	AlarmCodeRaiseUp             = 4
	AlarmCodeTargetLost          = 5
	AlarmCodeAbnormalBehavior    = 6
	AlarmCodeObjectLeft          = 7
	AlarmCodeStay                = 8
	AlarmCodeTimeout             = 9
	AlarmCodeInvalidOperation    = 10
	AlarmCodeThreadOperation     = 11
	AlarmCodeRemove              = 12
	AlarmCodeDutyCheck           = 13
	AlarmCodeFierceMotion        = 14
	AlarmCodeTermperatureAlarm   = 15
	AlarmCodeHumidityAlarm       = 16
	AlarmCodeDoorPointAlarm      = 17
	AlarmCodeForceOperation      = 18
	AlarmCodeSoundAlam           = 19
	AlarmCodeFaceDefenseAlarm    = 20
	AlarmCodeRFIDBrokenAlarm     = 21
	AlarmCodeHeartRateAlarm      = 22
	AlarmCodeRFIDButtonAlarm     = 23
	AlarmCodeSmokeAlarm          = 24
	AlarmCodeFireAlarm           = 25
	AlarmCodeGathering           = 26
	AlarmCodeDutySleep           = 27
	AlarmCodeFight               = 28
	AlarmCodePowerGridOpenAlarm  = 29
	AlarmCodePowerGridTouchAlarm = 30
	AlarmCodePowerGridShortAlarm = 31
	AlarmCodeVideoAnalyseAlarm   = 32
	AlarmCodeHang                = 33
	AlarmCodeCoverHeadSleep      = 34
)

// DeviceFaultCode
const (
	FaultCodeUnknown          = -1
	FaultCodeMotionDetect     = 0
	FaultCodeVideoLost        = 1
	FaultCodeVideoShelter     = 2
	FaultCodeAudioLost        = 3
	FaultCodeDiskFull         = 4
	FaultCodeDiskFault        = 5
	FaultCodeNetwork          = 6
	FaultCodeDisturb          = 7
	FaultCodePowerAbnormal    = 8
	FaultCodeVideoQuality     = 9
	FaultCodeOffline          = 10
	FaultCodeBright           = 11
	FaultCodeColor            = 12
	FaultCodeContrast         = 13
	FaultCodeFrozen           = 14
	FaultCodeBlur             = 15
	FaultCodeNoise            = 16
	FaultCodeWave             = 17
	FaultCodeRoll             = 18
	FaultCodeChroma           = 19
	FaultCodePTZ              = 20
	FaultCodeMono             = 21
	FaultCodeShake            = 22
	FaultCodeFlash            = 23
	FaultCodeScene            = 24
	FaultCodeDark             = 25
	FaultCodeSignalDelay      = 26
	FaultCodeVideoStreamDelay = 27
	FaultCodeKeyFrameDelay    = 28
	FaultCodeLuma             = 29
	FaultCodeMosaic           = 30
	FaultCodeRealStream       = 31
	FaultCodeRecord           = 32
	FaultCodeFault            = 33
	FaultCodeLowBattery       = 34
)

// DeviceStatusCode
const (
	StatusCodeUnknown    = -1
	StatusCodeNormal     = 0
	StatusCodeFault      = 1
	StatusCodeOpen       = 2
	StatusCodeClose      = 3
	StatusCodeStayOpen   = 4
	StatusCodeStayClose  = 5
	StatusCodeDisArm     = 6
	StatusCodeArm        = 7
	StatusCodeAlarm      = 8
	StatusCodeBypass     = 9
	StatusCodeDataUpdate = 10
)

// DeviceControlCode
const (
	ControlCodeStopInterfacer       = -1
	ControlCodeAwayArm              = 0
	ControlCodeStayArm              = 1
	ControlCodeDisArm               = 2
	ControlCodeConnectRelay         = 3
	ControlCodeDisconnectRelay      = 4
	ControlCodeBypass               = 5
	ControlCodeCancelBypass         = 6
	ControlCodeQueryArm             = 7
	ControlCodeCallSlave            = 8
	ControlCodeListenSlave          = 9
	ControlCodeHangupSlave          = 10
	ControlCodeStartBroadcast       = 11
	ControlCodeStopBroadcast        = 12
	ControlCodeSetContent           = 13
	ControlCodeResumeContent        = 14
	ControlCodeSwithDecoderChannel  = 15
	ControlCodeSwitchMatrixChannel  = 16
	ControlCodeMatrixPTZ            = 17
	ControlCodeMatrixPreset         = 18
	ControlCodeTvWallLayout         = 19
	ControlCodeOpenDoor             = 20
	ControlCodeCloseDoor            = 21
	ControlCodeResetArm             = 22
	ControlCodeLockDoor             = 23
	ControlCodeUnlockDoor           = 24
	ControlCodeVoiceBroadcast       = 25
	ControlCodeOpenWindowProcedure  = 26
	ControlCodeCloseWindowProcedure = 27
	ControlCodeStopProcedure        = 28
	ControlCodeAcON                 = 29
	ControlCodeAcOFF                = 30
	ControlCodeAcMode               = 31
	ControlCodeAcTemperature        = 32
	ControlCodeCmd                  = 33
)

// VideoEncode
const (
	VideoEncodeUnknown = -1
	VideoEncodeH264    = 0
	VideoEncodeH265    = 1
)
