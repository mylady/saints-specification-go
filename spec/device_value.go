package spec

//DoorValue :value for door
type DoorValue struct {
	PersonName     string `json:"person_name"`
	PersonCard     string `json:"person_card"`
	SceneImage     string `json:"scene_image"`
	SceneImagePath string `jsoon:"scene_image_path"`
}

//EnvValue :value for env_monitor
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
	X                 float32 `json:"x"`
	Y                 float32 `json:"y"`
	Z                 float32 `json:"z"`
	TagID             string  `json:"tag_id"`
	PersonID          string  `json:"person_id"`
	PersonName        string  `json:"person_name"`
	Battery           float32 `json:"battery"`
	HeartRate         float32 `json:"heart_rate"`
	BloodPressureHigh float32 `json:"blood_pressure_high"`
	BloodPressureLow  float32 `json:"blood_pressure_low"`
}

//UPSValue :value for ups
type UPSValue struct {
	BackupTime string  `json:"backup_time"`
	Battery    string  `json:"battery"`
	AVolter    float32 `json:"a_volter"`
	BVolter    float32 `json:"b_volter"`
	CVolter    float32 `json:"c_volter"`
	AElec      float32 `json:"a_elec"`
	BElec      float32 `json:"b_elec"`
	CElec      float32 `json:"c_elec"`
}

//CallValue :value for call
type CallValue struct {
	Caller string `json:"caller"`
	Called string `json:"called"`
}

//Face :value for face
type Face struct {
	PersonID      string  `json:"person_id"`
	PersonName    string  `json:"person_name"`
	Similarity    float32 `json:"similarity"`
	FaceImage     string  `json:"face_image"`
	FaceImagePath string  `json:"face_image_path"`
}

//CarPlate :value for car
type CarPlate struct {
	PlateNumber string `json:"plate_number"`
	PlateColor  string `json:"plate_color"`
	PlateType   string `json:"plate_type"`
	CarColor    string `json:"car_color"`
}

//VideoAnalyseType
const (
	VideoAnalyseTypeBehavior = 0
	VideoAnalyseTypeFace     = 1
	VideoAnalyseTypeCarPlate = 2
)

//VideoAnalyseValue :value for video analyse
type VideoAnalyseValue struct {
	AnalyseType    int      `json:"analyse_type"`
	SceneImage     string   `json:"scene_image"`
	SceneImagePath string   `json:"scene_image_path"`
	Face           Face     `json:"face"`
	Car            CarPlate `json:"car"`
}
