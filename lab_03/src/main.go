package main

import (
	"lab_03/emission"
)

func main() {
	xl := make(emission.FArr64, int((emission.Params.L+emission.Params.H)/emission.Params.H))
	for i := 1; i < len(xl); i++ {
		xl[i] = xl[i-1] + emission.Params.H
	}
	{
		tbl := emission.SimpleIters(0.25, 1e-5, 100)
		emission.DrawPlot(xl, tbl, "T(x)", "x", "T", "data/tx.png")
	}
	{
		emission.Params = emission.Emission{1.4, 0.2, 300, 400, 5.668e-12, 100, 0.05, 1e-4}
		emission.Params.T0 = 1200
		tbl := emission.SimpleIters(0.25, 1e-5, 100)
		emission.DrawPlot(xl, tbl, "T(x)", "x", "T", "data/tx0.png")
	}
	{
		emission.Params = emission.Emission{1.4, 0.2, 300, 400, 5.668e-12, 100, 0.05, 1e-4}
		emission.Params.F0 = -10
		tbl := emission.SimpleIters(0.25, 1e-5, 100)
		emission.DrawPlot(xl, tbl, "T(x)", "x", "T", "data/tx1.png")
	}
	{
		emission.Params = emission.Emission{1.4, 0.2, 300, 400, 5.668e-12, 100, 0.05, 1e-4}
		tbl1 := emission.SimpleIters(0.25, 1e-5, 100)
		emission.Params.Alpha *= 3
		tbl2 := emission.SimpleIters(0.25, 1e-5, 100)
		emission.DrawMultiplePlot(xl, tbl1, xl, tbl2, "T(x)", "x", "T", "data/tx2.png")
	}
	{
		emission.Params = emission.Emission{1.4, 0.2, 300, 400, 5.668e-12, 100, 0.05, 1e-4}
		emission.Params.F0 = 0
		tbl := emission.SimpleIters(0.25, 1e-5, 100)
		emission.DrawPlot(xl, tbl, "T(x)", "x", "T", "data/tx3.png")
	}
}
