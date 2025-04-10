package spec

// InterfaceCode
const (
	InterfaceCodeUnknown                           = -1
	InterfaceCodeHaiKangVideoSDK                   = 0
	InterfaceCodeDaHuaVideoSDK                     = 1
	InterfaceCodeDaHuaDSSVideoPlatformSDK          = 2
	InterfaceCodeDaHua7016VideoPlatformSDK         = 3
	InterfaceCodeWangLiPVGSDK                      = 4
	InterfaceCodeYuShiVideoSDK                     = 5
	InterfaceCodeJinSanLiVideoPlatformSDK          = 6
	InterfaceCodeBoShiMTWSoftwareAlarmUDP          = 7
	InterfaceCodeBoShiIP7400AlarmSDK               = 8
	InterfaceCodeHuoNiWeiErIP2000AlarmUDP          = 9
	InterfaceCodeHuoNiWeiErIPMAlarmOCX             = 10
	InterfaceCodeLaiBangAnalogCallSERIAL           = 11
	InterfaceCodeLaiBangDigitalCallTCP             = 12
	InterfaceCodeJiaoDaJingGongDoorTCP             = 13
	InterfaceCodeMaiTeAnAlarmPlatformSDK           = 14
	InterfaceCodeHaoWeiAlarmHTTP                   = 15
	InterfaceCodePorisAlarmOPC                     = 16
	InterfaceCodeSaiBangTeCallUDP                  = 17
	InterfaceCodeHaoWeiDVRSDK                      = 18
	InterfaceCodeHaoWeiNVRSDK                      = 19
	InterfaceCodeJinKaiDoorOPC                     = 20
	InterfaceCodeNiuBeiErDoorTCP                   = 21
	InterfaceCodePeakDoorOCX                       = 22
	InterfaceCodeWeiGenDoorSDK                     = 23
	InterfaceCodeYiTongDoorUDP                     = 24
	InterfaceCodeYangBangLedSDKv4                  = 25
	InterfaceCodeYiKuoLedSDKv2008                  = 26
	InterfaceCodeLinXinLedSDKv3_3                  = 27
	InterfaceCodeLinXinLedSDKv5_2013               = 28
	InterfaceCodeShiZhanLedSDK                     = 29
	InterfaceCodeXiangYunLedSDK                    = 30
	InterfaceCodeKuaiYuSoundAlarmv1SERIAL          = 31
	InterfaceCodeBoKangBVGTCP                      = 32
	InterfaceCodeHaiKangB20PlatformSDK             = 33
	InterfaceCodeHaoWeiDecoderSDK                  = 34
	InterfaceCodeYingFeiTuoMatrixUDP               = 35
	InterfaceCodeINTLMatrixSERIAL                  = 36
	InterfaceCodeJinSanLiMatrixUDP                 = 37
	InterfaceCodePuTaiKeAlarmOCX                   = 38
	InterfaceCodeDaHuaTvWallHTTPJSONRPC            = 39
	InterfaceCodeChuangWeiTvWallSERIAL             = 40
	InterfaceCodeTCLTvWallTCP                      = 41
	InterfaceCodeHuiGuUpsAlarmUDP                  = 42
	InterfaceCodeShanTeC6UpsSERIAL                 = 43
	InterfaceCodeShanTe3C3EXUpsSERIAL              = 44
	InterfaceCodeYiShiTeUpsSERIAL                  = 45
	InterfaceCodeHongShiVideoAnalyseTCP            = 46
	InterfaceCodeHongShenVideoAnalyseSDK           = 47
	InterfaceCodeLinXinLedSDKv5_2014               = 48
	InterfaceCodeHaiKang7200VideoPlatformSDK       = 49
	InterfaceCodeHaiKang8300VideoPlatformSDK       = 50
	InterfaceCodeHaiKang8700VideoPlatformSDK       = 51
	InterfaceCodeKeDaVideoPlatformSDK              = 52
	InterfaceCodeGuangTuoVideoDeviceSDK            = 53
	InterfaceCodeMeiZanMeiVideoDeviceSDK           = 54
	InterfaceCodeYingFeiTuoVideoDeviceSDK          = 55
	InterfaceCodeZhongXinLiWeiVideoDeviceSDK       = 56
	InterfaceCodeYuShiNvrSDK                       = 57
	InterfaceCodeNetworkCheck                      = 58
	InterfaceCodeKuaiYuSoundAlarmv2SERIAL          = 59
	InterfaceCodeTCLTvWallSERIAL                   = 60
	InterfaceCodeHuiGuUpsStatusUDP                 = 61
	InterfaceCodeYiShiTeUpsModbusSERIAL            = 62
	InterfaceCodeKuaiYuDigitalCallSDK              = 63
	InterfaceCodeAoSongEnvMonitorModbusTCP         = 64
	InterfaceCodeMiaoGuanEnvMonitorModbusTCP       = 65
	InterfaceCodeHaiKangDecoderSDK                 = 66
	InterfaceCodeDaHuaDecoderSDK                   = 67
	InterfaceCodeBoShiAlarmSERIAL                  = 68
	InterfaceCodeYiTuFaceRecognitionREST           = 69
	InterfaceCodeXiLingVideoAnalyseREST            = 70
	InterfaceCodeSiHuaRFIDWebSocket                = 71
	InterfaceCodeUnicomIOTREST                     = 72
	InterfaceCodeDaHuaVideoAnalyseSDK              = 73
	InterfaceCodeSiHuaRFIDLocatingREST             = 74
	InterfaceCodeJuYingRelayTCP                    = 75
	InterfaceCodeWanGuDoorUDP                      = 76
	InterfaceCodeShengKeAlarmTCP                   = 77
	InterfaceCodeLingXinLedSDKv6                   = 78
	InterfaceCodeYiTuFaceRecognitionV18REST        = 79
	InterfaceCodeRenWeiRFIDUDP                     = 80
	InterfaceCodeYuShiFaceRecognitionREST          = 81
	InterfaceCodeJinZhiVideoPlatformSDK            = 82
	InterfaceCodeMiaoGuanEnvMonitorModbusTCPServer = 83
	InterfaceCodeMeiZanMeiIPCEnvMontorUDP          = 84
	InterfaceCodeRenShuoModuleEnvMonitorSDK        = 85
	InterfaceCodeZhongKeYunKongEnvMonitorREST      = 86
	InterfaceCodeZhongXinWeiRFIDSDK                = 87
	InterfaceCodeXuFeiPowerGridTCP                 = 88
	InterfaceCodeXuFeiIrAlarmTCP                   = 89
	InterfaceCodeJieAnPerimeterTCP                 = 90
	InterfaceCodeSiNiuTeEnvMonitorREST             = 91
	InterfaceCodeKeDaVideoAnalyseWebService        = 92
	InterfaceCodeKeDaVideoDiagnosticREST           = 93
	InterfaceCodeKuaiYuDigitalCallSDKV2            = 94
	InterfaceCodeKuaiYuDigitalCallSDKV3            = 95
	InterfaceCodeGuangTuoPerimeterSDK              = 96
	InterfaceCodeShangTangFaceRecognitionREST      = 97
	InterfaceCodeQuShiVideoAnalyseTCPServer        = 98
	InterfaceCodeQuShiFaceAnalyseTCPServer         = 99
	InterfaceCodeHaiKangISCHTTPServer              = 100
	InterfaceCodeHaiKang8200ArtemisPlatformSDK     = 101
	InterfaceCodeKuaiYuDigitalBroadcastSDK         = 102
	InterfaceCodeHanWeiAirMonitorUDP               = 103
	InterfaceCodeJunLingUPSHTTPServer              = 104
	InterfaceCodeMaoTeRFIDREST                     = 105
	InterfaceCodeMeiDiACModbusSERIAL               = 106
	InterfaceCodeMeiDiACModbusTCPV4                = 107
	InterfaceCodeYongLinRelaySERIAL                = 108
	InterfaceCodeYongLinRelayTCP                   = 109
	InterfaceCodeBiYiTeRelaySERIAL                 = 110
	InterfaceCodeBiYiTeRelayTCP                    = 111
	InterfaceCodeFeiFanAlarmHostTCPServer          = 112
	InterfaceCodeYiTuVideoAnalyseHTTPServer        = 113
	InterfaceCodeHanWeiAirMonitorUDPv2             = 114
	InterfaceCodeXuFeiElecDoorTCP                  = 115
	InterfaceCodeHanWeiElecDoorTCP                 = 116
	InterfaceCodeSiFangBoRuiRelayPlatformREST      = 117
	InterfaceCodeMeiZanMeiFaceAnalyseHTTPServer    = 118
	InterfaceCodeDuYiRFIDRest                      = 119
	InterfaceCodeHeShiVideoAnalyseHTTPServer       = 120
	InterfaceCodeTengDaEnvMonistorMQTT             = 121
	InterfaceCodeTengDaRFIDREST                    = 122
	InterfaceCodeMeiDiACModbusTCPV5                = 123
	InterfaceCodeHaiKangVideoSDKV2                 = 124
	InterfaceCodeHeShiVideoAnalyseActiveMQV2       = 125
	InterfaceCodeJiaoDaJingGongDoorHTTPV2          = 126
	InterfaceCodeZhongGuoDianZiHTTPServer          = 127
	InterfaceCodeTengDaRelayMQTT                   = 128
	InterfaceCodeTengDaAcMQTT                      = 129

	InterfaceCodeSaintsDoor = 10001
)

// DeviceEventData :device event data
type DeviceEventData struct {
	ResourceCode int    `json:"resource_code"`
	Address      string `json:"address"`
	EventType    int    `json:"event_type"`
	EventCode    int    `json:"event_code"`
	AlarmCode    int    `json:"alarm_code"`
	StatusCode   int    `json:"status_code"`
	FaultCode    int    `json:"fault_code"`
	EventTime    string `json:"event_time"`
	EventDesp    string `json:"event_desp"`
	EventExtra   string `json:"event_extra"`
}
