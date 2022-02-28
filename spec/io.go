package spec

/*
	Rest service
*/

//Pager :interface for http paging
type Pager interface {
	GetStart() int
	GetLimit() int
	GetSort() string
	GetDir() string
	GetKeyword() string
}

//RestPagingQuery :query parameter for paging rest service
type RestPagingQuery struct {
	Start   int    `form:"start,defalut=0" json:"start"`
	Limit   int    `form:"limit,default=100" json:"limit"`
	Sort    string `form:"sort" json:"sort"`
	Dir     string `form:"dir,default=asc" json:"dir"`
	Keyword string `form:"keyword" json:"keyword"`
}

//GetStart :impl pager
func (r *RestPagingQuery) GetStart() int {
	return r.Start
}

//GetLimit :impl pager
func (r *RestPagingQuery) GetLimit() int {
	return r.Limit
}

//GetSort :impl pager
func (r *RestPagingQuery) GetSort() string {
	return r.Sort
}

//GetDir :impl pager
func (r *RestPagingQuery) GetDir() string {
	return r.Dir
}

//GetKeyword :impl pager
func (r *RestPagingQuery) GetKeyword() string {
	return r.Keyword
}

type ComplexSort struct {
	Sort string `json:"sort"`
	Dir  string `json:"dir"`
}

//ComplexPager :interface for http post paging
type ComplexPager interface {
	GetStart() int
	GetLimit() int
	GetComplexSorts() []ComplexSort
	GetKeyword() string
}

//ComplexRestPagingQuery :query parameter for post paging rest service
type ComplexRestPagingQuery struct {
	Start   int           `json:"start"`
	Limit   int           `json:"limit"`
	Sorts   []ComplexSort `json:"sorts"`
	Keyword string        `json:"keyword"`
}

//GetStart :impl pager
func (r *ComplexRestPagingQuery) GetStart() int {
	return r.Start
}

//GetLimit :impl pager
func (r *ComplexRestPagingQuery) GetLimit() int {
	return r.Limit
}

//GetSort :impl pager
func (r *ComplexRestPagingQuery) GetComplexSorts() []ComplexSort {
	return r.Sorts
}

//GetKeyword :impl pager
func (r *ComplexRestPagingQuery) GetKeyword() string {
	return r.Keyword
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
