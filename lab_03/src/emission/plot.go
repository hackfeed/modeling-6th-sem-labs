package emission

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

// DrawMultiplePlot is used to draw multiple plots with given coordinates and meta info on one plot
func DrawMultiplePlot(xs1, ys1, xs2, ys2 FArr64, title, xl, yl, file string) {
	p := plot.New()

	p.Title.Text = title
	p.X.Label.Text = xl
	p.Y.Label.Text = yl
	p.Add(plotter.NewGrid())

	dots1 := convertDots(xs1, ys1)
	dots2 := convertDots(xs2, ys2)

	l1, err := plotter.NewLine(dots1)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	l1.LineStyle.Width = vg.Points(1)
	l1.LineStyle.Color = color.RGBA{B: 255, A: 255}

	l2, err := plotter.NewLine(dots2)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	l2.LineStyle.Width = vg.Points(1)
	l2.LineStyle.Color = color.RGBA{R: 255, A: 255}

	p.Add(l1, l2)

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
