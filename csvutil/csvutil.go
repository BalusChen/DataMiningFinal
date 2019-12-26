package csvutil

import (
	"encoding/csv"
	"fmt"
	"github.com/BalusChen/DataMiningFinalProject/core"
	"os"
	"strconv"
)

const (
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

func ReadRecord(filename string) []*core.ElecCarInfo {
	csvReader := newCSVReader(filename)
	csvReader.Read() // skip title
	records, err := csvReader.ReadAll()
	if err != nil {
		panic(err)
	}

	elecCarInfos := make([]*core.ElecCarInfo, 0, InitialArraySize)
	for _, record := range records {
		if len(record) != core.NumOfFields {
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

func parseRecord(record []string) (*core.ElecCarInfo, error) {
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

	elecCarInfo := &core.ElecCarInfo{
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
