package saints_specification_go

//RestResponse :beego rest response struct
type RestResponse struct {
	Result     bool        `json:"result"`
	ErrorMsg   string      `json:"error_msg"`
	ErrorCode  int         `json:"error_code"`
	Data       interface{} `json:"data"`
	TotalCount int64       `json:"total_count"`
}

//NewRestResponse :create default ServerResponse struct
func NewRestResponse(success bool, errmsg string, data interface{}, totalcnt int64) (response *RestResponse) {
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
