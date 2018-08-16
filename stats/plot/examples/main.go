package main

import (
	"net/http"
	"time"

	"github.com/wcharczuk/go-chart"
)

type Line struct {
	X []time.Time
	Y []float64
}

func (l *Line) DrawAckPercent(res http.ResponseWriter, req *http.Request) {
	mainSeries := chart.TimeSeries{
		Name: "ACK percentage",
		Style: chart.Style{
			Show:        true,
			StrokeColor: chart.ColorBlue,
			FillColor:   chart.ColorBlue.WithAlpha(100),
		},
		XValues: l.X,
		YValues: l.Y,
	}

	graph := chart.Chart{
		Width:  1280,
		Height: 720,
		YAxis: chart.YAxis{
			Name:      "Percentage",
			Style:     chart.StyleShow(),
			NameStyle: chart.StyleShow(),
		},
		XAxis: chart.XAxis{
			Name:           "Time",
			Style:          chart.StyleShow(),
			NameStyle:      chart.StyleShow(),
			ValueFormatter: chart.TimeMinuteValueFormatter,
		},

		Series: []chart.Series{
			mainSeries,
		},
	}

	res.Header().Set("Content-Type", chart.ContentTypePNG)
	graph.Render(chart.PNG, res)
}

func lineGraph(x []time.Time, y []float64) {
	l := &Line{
		X: x,
		Y: y,
	}

	http.HandleFunc("/", l.DrawAckPercent)
	http.ListenAndServe(":8080", nil)
}

func main() {
	var x []time.Time
	var y []float64
	seed := time.Now()
	for i := 100; i > 0; i-- {
		seed = seed.Add(time.Minute)
		x = append(x, seed)
		y = append(y, float64(i%10))
	}

	lineGraph(x, y)
}
