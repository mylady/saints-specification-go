package service_type

const (
	//0-1023 service provided by saints
	DeviceDataService      = 0
	DeviceEventService     = 1
	DeviceStatusService    = 2
	MaintainService        = 3
	PrisonerDataService    = 4
	SSOService             = 5
	DigitalService         = 6
	CentralUpdateService   = 7
	TransmitMessageService = 8
	SerialGatewayService   = 9
	SendMessageService     = 10
	VideoService           = 11
	SchoolService          = 12
	PrisonService          = 13

	//1024 - 2048 service provided by third party
	WebMapService           = 1024
	WebFeatureService       = 1025
	WebTiledMapService      = 1026
	ArcgisDynamicMapService = 1027
	ArcgisTiledMapService   = 1028
	ArcgisGeometryService   = 1029
)
