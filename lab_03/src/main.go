package main

import (
	"lab_03/optic"
)

func main() {
	lt := optic.Interpolate(optic.LambdaTbl[0], optic.LambdaTbl[1])
	kt := optic.Interpolate(optic.KTbl[0], optic.KTbl[1])
	tbl := make(optic.FArr64, int(1./optic.Params.H)+2)

	xil := optic.FArr64{0}
	etal := optic.FArr64{0}
	xl := optic.FArr64{}

	x := 0.
	n := 0

	for x+optic.Params.H < 1 {
		xl = append(xl, x)
		xil = append(xil, optic.C(lt, tbl, n)/(optic.B(lt, kt, tbl, n)-optic.A(lt, tbl, n)*xil[n]))
		etal = append(etal, (optic.D(kt, tbl, n)+optic.A(lt, tbl, n)*xil[n])/
			(optic.B(lt, kt, tbl, n)-optic.A(lt, tbl, n)*xil[n]))

		n++
		x += optic.Params.H
	}

	xl = append(xl, x+optic.Params.H, x+optic.Params.H*2)

	lcs := optic.GetLConds(lt, kt, tbl, n)
	tbl[n] = (lcs.P - lcs.M*xil[n]) / (lcs.K + lcs.M*xil[n])

	for i := n - 1; i > -1; i-- {
		tbl[i] = xil[i+1]*tbl[i+1] + etal[i+1]
	}

	optic.DrawPlot(xl, tbl, "T(x)", "x", "T", "data/tx.png")
}
