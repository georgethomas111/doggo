package plot

import (
	"io"
	"time"

	"github.com/wcharczuk/go-chart"
)

type Line struct {
	X      []time.Time
	Y      []float64
	XName  string
	YName  string
	Width  int
	Height int
}

// DrawLine accepts a writer and writes the graph in PNG format to it..
func (l *Line) DrawLine(w io.Writer) {
	mainSeries := chart.TimeSeries{
		Style: chart.Style{
			Show:        true,
			StrokeColor: chart.ColorBlue,
			FillColor:   chart.ColorBlue.WithAlpha(100),
		},
		XValues: l.X,
		YValues: l.Y,
	}

	graph := chart.Chart{
		Width:  l.Width,
		Height: l.Height,
		YAxis: chart.YAxis{
			Name:      l.YName,
			Style:     chart.StyleShow(),
			NameStyle: chart.StyleShow(),
		},
		XAxis: chart.XAxis{
			Name:           l.XName,
			Style:          chart.StyleShow(),
			NameStyle:      chart.StyleShow(),
			ValueFormatter: chart.TimeMinuteValueFormatter,
		},

		Series: []chart.Series{
			mainSeries,
		},
	}

	graph.Render(chart.PNG, w)
}
