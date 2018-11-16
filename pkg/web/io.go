package web

/*
	Rest service
*/

//RestPagingQuery :query parameter for paging rest service
type RestPagingQuery struct {
	Start   int    `json:"start"`
	Limit   int    `json:"limit"`
	Sort    string `json:"sort"`
	Dir     string `json:"dir"`
	Keyword string `json:"keyword"`
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
