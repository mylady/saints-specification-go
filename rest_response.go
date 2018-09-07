package saints_specification_go

import (
	"encoding/base64"
	"encoding/json"
)

var enableEncrypt = false

func SetEncrypt(enable bool) {
	enableEncrypt = enable
}

//RestResponse :beego rest response struct
type RestResponse struct {
	Result     bool        `json:"result"`
	ErrorMsg   string      `json:"error_msg"`
	Data       interface{} `json:"data"`
	TotalCount int64       `json:"total_count"`
}

//NewRestResponse :create default ServerResponse struct
func NewRestResponse(success bool, errmsg string, data interface{}, totalcnt int64) (response *RestResponse) {
	if enableEncrypt {
		data = encryptData(data)
	}
	return &RestResponse{
		Result:     success,
		ErrorMsg:   errmsg,
		Data:       data,
		TotalCount: totalcnt,
	}
}

//NewBoolResponse :create bool only ServerResponse struct
func NewBoolResponse(success bool, errmsg string) (response *RestResponse) {
	return NewRestResponse(success, errmsg, nil, 0)
}

//NewDataResponse :create bool and data ServerResponse struct
func NewDataResponse(data interface{}, totalcnt int64) (response *RestResponse) {
	return NewRestResponse(true, "", data, totalcnt)
}

//encrypt resp
func encryptData(data interface{}) string {
	jsonStr, _ := json.Marshal(data)
	return base64.StdEncoding.EncodeToString(jsonStr)
}
