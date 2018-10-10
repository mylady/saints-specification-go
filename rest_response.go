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

//NewCodeRestResponse :create default ServerResponse struct with err code
func NewCodeRestResponse(success bool, errcode int, errmsg string, data interface{}, totalcnt int64) (response *RestResponse) {
	return &RestResponse{
		Result:     success,
		ErrorCode:  errcode,
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

//NewCodeResponse :create bool with code ServerResponse struct
func NewCodeResponse(success bool, errcode int, errmsg string) (response *RestResponse) {
	return NewCodeRestResponse(success, errcode, errmsg, nil, 0)
}
