package interface_spec

type RFIDTagExtra struct {
	TagId             string
	PersonId          string
	PersonName        string
	BatteryVoltage    int
	HeartRate         int
	BloodPressureHigh int
	BloodPressureLow  int
	X                 float32
	Y                 float32
	Z                 float32
}
