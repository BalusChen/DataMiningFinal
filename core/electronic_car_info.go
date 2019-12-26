package core

import "strconv"

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
	MoveDepth         float64 // feat8
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



func ParseRecord(record []string) (*ElecCarInfo, error) {
	energy, _ := strconv.ParseUint(record[1], 10, 64)
	hasBluetooth, _ := strconv.ParseBool(record[2])
	instructionSpeed, _ := strconv.ParseFloat(record[3], 64)
	hasDualSIM, _ := strconv.ParseBool(record[4])
	frontCameraPixel, _ := strconv.ParseUint(record[5], 10, 64)
	has4G, _ := strconv.ParseBool(record[6])
	memoryGB, _ := strconv.ParseUint(record[7], 10, 64)
	moveDepth, _ := strconv.ParseFloat(record[8], 64)
	weight, _ := strconv.ParseUint(record[9], 10, 64)
	numCores, _ := strconv.ParseUint(record[10], 10, 64)
	mainCameraPxiel, _ := strconv.ParseUint(record[11], 10, 64)
	cameraPixelHeight, _ := strconv.ParseUint(record[12], 10, 64)
	cameraPixelWidth, _ := strconv.ParseUint(record[13], 10, 64)
	ramMB, _ := strconv.ParseUint(record[14], 10, 64)
	screenHeight, _ := strconv.ParseUint(record[15], 10, 64)
	screenWidth, _ := strconv.ParseUint(record[16], 10, 64)
	chargeTime, _ := strconv.ParseUint(record[17], 10, 64)
	has3G, _ := strconv.ParseBool(record[18])
	hasTouchScreen, _ := strconv.ParseBool(record[19])
	hasWifi, _ := strconv.ParseBool(record[20])
	price, _ := strconv.ParseUint(record[21], 10, 64)

	elecCarInfo := &ElecCarInfo{
		ID:                record[0],
		Energy:            energy,
		HasBlueotooth:     hasBluetooth,
		InstructionSpeed:  instructionSpeed,
		HasDualSIM:        hasDualSIM,
		FontCameraPixel:   frontCameraPixel,
		Has4G:             has4G,
		MemoryGB:          memoryGB,
		MoveDepth:         moveDepth,
		Weight:            weight,
		NumCores:          numCores,
		MainCameraPixel:   mainCameraPxiel,
		CameraPixelHeight: cameraPixelHeight,
		CameraPixelWidth:  cameraPixelWidth,
		RAMMB:             ramMB,
		ScreenHeight:      screenHeight,
		ScreenWidth:       screenWidth,
		ChargeTime:        chargeTime,
		Has3G:             has3G,
		HasTouchScreen:    hasTouchScreen,
		HasWifi:           hasWifi,
		Price:             price,
	}

	return elecCarInfo, nil
}
