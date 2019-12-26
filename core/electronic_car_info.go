package core

const (
	NumOfFields = 22
)

type ElecCarInfo struct {
	ID                string
	Energy            uint64  // feat1
	HasBlueotooth     bool    // feat2
	InstructionSpeed  float64 // feat3
	HasDualSIM        bool    // feat4
	FontCameraPixel   uint64  // feat5
	Has4G             bool    // feat6
	MemoryGB          uint64  // feat7
	MoveDepth         uint64  // feat8
	Weight            uint64  // feat9
	NumCores          uint64  // feat10
	MainCameraPixel   uint64  // feat11
	CameraPixelHeight uint64  // feat12
	CameraPixelWidth  uint64  // feat13
	RAMMB             uint64  // feat14
	ScreenHeight      uint64  // feat15
	ScreenWidth       uint64  // feat16
	ChargeTime        uint64  // feat17
	Has3G             bool    // feat18
	HasTouchScreen    bool    // feat19
	HasWifi           bool    // feat20
	Price             uint64
}
