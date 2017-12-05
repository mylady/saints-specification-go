package saints_specification_go

type BioIdentifyMessage struct {
	IdentifyType    int
	IdentifyTime    JsonTime
	IdentifyContent string
}
