package saints_specification_go

const (
	//StartParam :paging start index
	StartParam = "start"
	//LimitParam :paging limit
	LimitParam = "limit"
	//SortParam :paging sort
	SortParam = "sort"
	//DirectionParam :paging order
	DirectionParam = "dir"
	//SearchParam :paging search
	SearchParam = "keyword"
)

//RestQuery :struct for query string paging params
type RestQuery struct {
	Start     int
	Limit     int
	Sort      string
	Direction string
	Search    string
}
