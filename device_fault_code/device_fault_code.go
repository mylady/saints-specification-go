package device_fault_code

import (
	"github.com/mylady/saints-specification-go/device_alarm_code"
)

const (
	VideoLost     = device_alarm_code.VideoLost
	AudioLost     = device_alarm_code.AudioLost
	DiskFull      = device_alarm_code.DiskFull
	DiskFault     = device_alarm_code.DiskFault
	DeviceFault   = device_alarm_code.DeviceFault
	Offline       = device_alarm_code.Offline
	Disturb       = device_alarm_code.Disturb
	PowerAbnormal = device_alarm_code.PowerAbnormal
	VideoQuality  = device_alarm_code.VideoQuality
)

var FaultCodeMap = make(map[int]string)

func init() {
	FaultCodeMap[VideoLost] = "视频丢失"
	FaultCodeMap[AudioLost] = "音频丢失"
	FaultCodeMap[DiskFull] = "硬盘满"
	FaultCodeMap[DiskFault] = "硬盘故障"
	FaultCodeMap[DeviceFault] = "设备故障"
	FaultCodeMap[Offline] = "离线"
	FaultCodeMap[Disturb] = "干扰"
	FaultCodeMap[PowerAbnormal] = "电源异常"
	FaultCodeMap[VideoQuality] = "视质异常"
}
