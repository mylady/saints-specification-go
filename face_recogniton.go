package saints_specification_go

type FaceRecognition struct {
	CameraAddress  string
	FaceImagePath  string
	SceneImagePath string
	FaceX          int
	FaceY          int
	FaceWidth      int
	FaceHeight     int
	IsHit          bool
	HitImagePath   string
	Similarity     float64
	PersonId       string
	PersonName     string
}
