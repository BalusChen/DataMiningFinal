package csvutil

import (
	"encoding/csv"
	"github.com/BalusChen/DataMiningFinalProject/core"
	"os"
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

		elecCarInfo, err := core.ParseRecord(record)
		if err == nil {
			elecCarInfos = append(elecCarInfos, elecCarInfo)
		}
	}

	/*
	for _, elecCarInfo := range elecCarInfos {
		fmt.Printf("%+v\n", elecCarInfo)
	}
	 */

	return elecCarInfos
}
