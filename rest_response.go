package saints_specification_go

//ServerResponse :beego rest response struct
type ServerResponse struct {
	Result     bool        `json:"result"`
	ErrorMsg   string      `json:"error_msg"`
	Data       interface{} `json:"data"`
	TotalCount int64       `json:"total_count"`
}

//NewServerResponse :create default ServerResponse struct
func NewServerResponse(success bool, errmsg string, data interface{}, totalcnt int64) (response *ServerResponse) {
	return &ServerResponse{
		Result:     success,
		ErrorMsg:   errmsg,
		Data:       data,
		TotalCount: totalcnt,
	}
}

//NewBoolResponse :create bool only ServerResponse struct
func NewBoolResponse(success bool, errmsg string) (response *ServerResponse) {
	return NewServerResponse(success, errmsg, nil, 0)
}

//NewDataResponse :create bool and data ServerResponse struct
func NewDataResponse(data interface{}, totalcnt int64) (response *ServerResponse) {
	return NewServerResponse(true, "", data, totalcnt)
}
