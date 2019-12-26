package main

import (
	"flag"
	"github.com/BalusChen/DataMiningFinalProject/classify"
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
	results, kindsOfOptions, name := classify.Wifi(records)
	draw.BarChart(results, kindsOfOptions, name)
}

func drawDiscrete() {

}

func drawContinous() {

}
