package draw

import (
	"fmt"
	"github.com/BalusChen/DataMiningFinalProject/classify"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
)

const (
	defaultChartWidth = 20
)

// DrawHistogram draws histogram
func BarChart(infos [classify.KindsOfPriceOptions][]uint64, kindOfOptions int, chartName string) {
	groups := make([]plotter.Values, kindOfOptions)
	for i := 0; i < kindOfOptions; i++ {
		groups[i] = plotter.Values{}
		for j := 0; j < classify.KindsOfPriceOptions; j++ {
			groups[i] = append(groups[i], float64(infos[j][i]))
		}
	}

	// debug
	fmt.Printf("%v", groups)

	p, err := plot.New()
	if err != nil {
		panic(err)
	}
	p.Title.Text = chartName
	p.X.Label.Text = "Price"
	p.Y.Label.Text = "Count"

	offsets := -vg.Points(float64(defaultChartWidth * kindOfOptions / 2))
	incr := vg.Points(float64(defaultChartWidth))
	for i := 0; i < kindOfOptions; i++ {
		bars, err := plotter.NewBarChart(groups[i], defaultChartWidth)
		if err != nil {
			panic(err)
		}

		bars.LineStyle.Width = vg.Length(0)
		bars.Color = plotutil.Color(i)
		bars.Offset = offsets
		offsets += incr

		p.Add(bars)
		p.Legend.Add("Group-", bars)
	}

	p.Legend.Top = true
	p.NominalX("0", "1", "2", "3")
	if err := p.Save(5*vg.Inch, 5*vg.Inch, chartName + ".png"); err != nil {
		panic(err)
	}
}
