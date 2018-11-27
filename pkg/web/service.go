package web

//Service :saints service struct
type Service struct {
	Address  string `json:"address"`
	Type     int    `json:"type"`
	Priority int    `json:"priority"`
}

//ServiceIPHolder :holder for address
const ServiceIPHolder = "ip"

//Service type enum
const (
	ServiceTypeIdentity = iota
	ServiceTypeDevice
	ServiceTypeFileHub
	ServiceTypeGIS
)

//ServiceTypeDict :service type name dict
var ServiceTypeDict = map[int]string{
	ServiceTypeIdentity: "身份认证服务",
	ServiceTypeDevice:   "设备信息服务",
	ServiceTypeFileHub:  "文件管理服务",
	ServiceTypeGIS:      "地理信息服务",
}
