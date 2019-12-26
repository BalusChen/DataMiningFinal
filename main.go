package main

import (
	"flag"
	"github.com/BalusChen/DataMiningFinalProject/classify"
	"github.com/BalusChen/DataMiningFinalProject/core"
	"github.com/BalusChen/DataMiningFinalProject/csvutil"
	"github.com/BalusChen/DataMiningFinalProject/draw"
)

const (
	defaultTrainCSVFilename = "data/train.csv"
)

var (
	trainCSVFilename string
)

func init() {
	flag.StringVar(&trainCSVFilename, "t", defaultTrainCSVFilename, "train csv")
}

func main() {
	flag.Parse()

	records := csvutil.ReadRecord(trainCSVFilename)

	drawDiscrete(records)
}

func drawDiscrete(records []*core.ElecCarInfo) {
	var (
		results        [4][]uint64
		kindsOfOptions int
		name           string
	)

	// hasWifi?
	results, kindsOfOptions, name = classify.Wifi(records)
	draw.BarChart(results, kindsOfOptions, name)

	// hasBluetooth?
	results, kindsOfOptions, name = classify.Bluetooth(records)
	draw.BarChart(results, kindsOfOptions, name)

	// has Dual SIM?
	results, kindsOfOptions, name = classify.DualSIM(records)
	draw.BarChart(results, kindsOfOptions, name)
}

func drawContinous() {

}
