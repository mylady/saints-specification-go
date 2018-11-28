package web

/*
	Rest service
*/

//RestPagingQuery :query parameter for paging rest service
type RestPagingQuery struct {
	Start   int    `form:"start,defalut=0" json:"start"`
	Limit   int    `form:"limit,default=1000" json:"limit"`
	Sort    string `form:"sort" json:"sort"`
	Dir     string `form:"dir,default=asc" json:"dir"`
	Keyword string `form:"keyword" json:"keyword"`
}

//RestResponse :response for rest service
type RestResponse struct {
	Result  bool        `json:"result"`
	Data    interface{} `json:"data"`
	Total   int64       `json:"total"`
	ErrMsg  string      `json:"err_msg"`
	ErrCode int         `json:"err_code"`
}

//SuccessResp :success response without data
func SuccessResp() RestResponse {
	return newResp(true, nil, 0, "", 0)
}

//SuccessDataResp :success response with data
func SuccessDataResp(d interface{}, t int64) RestResponse {
	return newResp(true, d, t, "", 0)
}

//FailureResp :fail response
func FailureResp(m string, c int) RestResponse {
	return newResp(false, nil, 0, m, c)
}

func newResp(r bool, d interface{}, t int64, m string, c int) RestResponse {
	return RestResponse{
		Result:  r,
		Data:    d,
		Total:   t,
		ErrMsg:  m,
		ErrCode: c,
	}
}

/*
	GraphQL
*/

/*
	PrototypeBuf
*/
