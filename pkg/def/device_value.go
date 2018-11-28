package def

//EnvValue :struct for env_monitor
type EnvValue struct {
	Temp float32 `json:"temp"`
	Humi float32 `json:"humi"`
	Hcho float32 `json:"hcho"`
	Co2  float32 `json:"co2"`
	Pm25 float32 `json:"pm25"`
	Tvoc float32 `json:"tvoc"`
	Nh3  float32 `json:"nh3"`
	H2s  float32 `json:"h2s"`
}

//PowerGridValue : value for power grid
type PowerGridValue struct {
	Volter float32 `json:"volter"`
	Elec   float32 `json:"elec"`
}

//RFIDTagValue :value for rfid tag
type RFIDTagValue struct {
	X         float32 `json:"x"`
	Y         float32 `json:"y"`
	Battery   float32 `json:"battery"`
	HeartRate float32 `json:"heart_rate"`
}

//UPSValue :value for ups
type UPSValue struct {
	Volter     float32 `json:"volter"`
	Elec       float32 `json:"elec"`
	BackupTime float32 `json:"backup_time"`
}
