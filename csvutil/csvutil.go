package csvutil

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

const (
	NumOfFields      = 22
	InitialArraySize = 100
)

var ()

func newCSVReader(filename string) *csv.Reader {
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	return csv.NewReader(f)
}

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

func ReadRecord(filename string) []*ElecCarInfo {
	csvReader := newCSVReader(filename)
	records, err := csvReader.ReadAll()
	if err != nil {
		panic(err)
	}

	elecCarInfos := make([]*ElecCarInfo, strconv.IntSize)
	for _, record := range records {
		if len(record) != NumOfFields {
			continue
		}

		elecCarInfo, err := parseRecord(record)
		if err == nil {
			elecCarInfos = append(elecCarInfos, elecCarInfo)
		}
	}

	for _, elecCarInfo := range elecCarInfos {
		fmt.Printf("%+v\n", elecCarInfo)
	}

	return elecCarInfos
}

func Classify() map[interface{}]*ElecCarInfo {
}

func parseRecord(record []string) (*ElecCarInfo, error) {
	energy, _ := strconv.ParseUint(record[1], 10, 64)
	hasBluetooth, _ := strconv.ParseBool(record[2])
	instructionSpeed, _ := strconv.ParseFloat(record[3], 64)
	hasDualSIM, _ := strconv.ParseBool(record[4])
	frontCameraPixel, _ := strconv.ParseUint(record[5], 10, 64)
	has4G, _ := strconv.ParseBool(record[6])
	memoryGB, _ := strconv.ParseUint(record[7], 10, 64)
	moveDepth, _ := strconv.ParseUint(record[8], 10, 64)
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
		record[0],
		energy,
		hasBluetooth,
		instructionSpeed,
		hasDualSIM,
		frontCameraPixel,
		has4G,
		memoryGB,
		moveDepth,
		weight,
		numCores,
		mainCameraPxiel,
		cameraPixelHeight,
		cameraPixelWidth,
		ramMB,
		screenHeight,
		screenWidth,
		chargeTime,
		has3G,
		hasTouchScreen,
		hasWifi,
		price,
	}

	return elecCarInfo, nil
}
