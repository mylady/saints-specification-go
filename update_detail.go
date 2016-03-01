package saints_go_specification

type UpdateDetail struct {
	ProductId    string
	OS           uint
	OldVersion   uint
	NewVersion   uint
	ShouldUpdate bool
	UpdateAddres string
}
