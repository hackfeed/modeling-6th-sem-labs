package circuit

import (
	"fmt"
	"image/color"
	"os"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

// DrawPlot is used to draw plot with given coordinates and meta info
func DrawPlot(xs, ys FArr64, title, xl, yl, file string) {
	p := plot.New()

	p.Title.Text = title
	p.X.Label.Text = xl
	p.Y.Label.Text = yl
	p.Add(plotter.NewGrid())

	dots := convertDots(xs, ys)

	l, err := plotter.NewLine(dots)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	l.LineStyle.Width = vg.Points(1)
	l.LineStyle.Color = color.RGBA{B: 255, A: 255}

	p.Add(l)

	if err := p.Save(10*vg.Inch, 4*vg.Inch, file); err != nil {
		panic(err)
	}
}

func convertDots(xs, ys FArr64) plotter.XYs {
	var conv plotter.XYs

	for i := 0; i < len(xs); i++ {
		d := plotter.XY{
			X: xs[i],
			Y: ys[i],
		}
		conv = append(conv, d)
	}

	return conv
}
