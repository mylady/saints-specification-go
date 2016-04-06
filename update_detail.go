package saints_specification_go

type UpdateDetail struct {
	ProductId    string
	OS           uint
	OldVersion   uint
	NewVersion   uint
	ShouldUpdate bool
	UpdateAddres string
}
