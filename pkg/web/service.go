package web

//Service :saints service struct
type Service struct {
	Address  string `json:"address"`
	Type     int    `json:"type"`
	Priority int    `json:"priority"`
}

//Service type enum
const (
	ServiceTypeIdentity = iota
	ServiceTypeDevice
	ServiceTypeFileHub
	ServiceTypeGIS
)
