package emission

import (
	"fmt"
	"math"

	"gonum.org/v1/gonum/integrate"
	"gonum.org/v1/gonum/interp"
)

func SimpleIters(al, eps float64, mits int) FArr64 {
	lt := interpolate(LambdaTbl[0], LambdaTbl[1])
	kt := interpolate(KTbl[0], KTbl[1])
	its := 0

	ts := make(FArr64, int(Params.L/Params.H)+1)
	for i := 0; i < len(ts); i++ {
		ts[i] = Params.T0
	}

	lcs := getLConds(lt, kt, ts)
	rcs := getRConds(lt, kt, ts)

	a, b, c, d := calcCoeffs(kt, lt, ts)

	tsn := runningThrough(a, b, c, d, lcs, rcs)
	for math.Abs((getF1(tsn)-getF2(kt, tsn))/getF1(tsn)) > eps && its < mits {
		ts = tsn
		lcs = getLConds(lt, kt, ts)
		rcs = getRConds(lt, kt, ts)
		a, b, c, d = calcCoeffs(kt, lt, ts)

		tsxr := runningThrough(a, b, c, d, lcs, rcs)
		for i := 0; i < len(ts); i++ {
			tsn[i] = ts[i] + Params.Alpha*(tsxr[i]-ts[i])
		}

		its++
	}

	return tsn
}

func calcCoeffs(kt, lt interp.AkimaSpline, tbl FArr64) (FArr64, FArr64, FArr64, FArr64) {
	n := int(Params.L / Params.H)
	a := make(FArr64, n-1)
	b := make(FArr64, n-1)
	c := make(FArr64, n-1)
	d := make(FArr64, n-1)

	for i := 1; i < n; i++ {
		lnm := lt.Predict(tbl[i-1])
		ln := lt.Predict(tbl[i])
		lnp := lt.Predict(tbl[i+1])

		a[i-1] = getXi(lnm, ln) / Params.H
		c[i-1] = getXi(ln, lnp) / Params.H
		b[i-1] = a[i-1] + c[i-1] + getP(tbl[i], kt)*Params.H
		d[i-1] = getF(tbl[i], kt) * Params.H
	}

	return a, b, c, d
}

func runningThrough(a, b, c, d FArr64, lcs, rcs Conds) FArr64 {
	n := len(a)

	xil := FArr64{-lcs.M / lcs.K}
	etal := FArr64{lcs.P / lcs.K}

	for i := 0; i < n; i++ {
		xil = append(xil, c[i]/(b[i]-a[i]*xil[i]))
		etal = append(etal, (d[i]+a[i]*etal[i])/(b[i]-a[i]*xil[i]))
	}

	y := FArr64{(rcs.P - rcs.K*etal[n]) / (rcs.M + rcs.K*xil[n])}
	for i := n; i > -1; i-- {
		y = append([]float64{xil[i]*y[0] + etal[i]}, y...)
	}

	return y
}

func getLConds(lt, kt interp.AkimaSpline, tbl FArr64) Conds {
	var cs Conds

	hs := Params.H * Params.H
	xif := getXi(lt.Predict(tbl[0]), lt.Predict(tbl[1]))
	pf := getXi(getP(tbl[0], kt), getP(tbl[1], kt))
	ff := getXi(getF(tbl[0], kt), getF(tbl[0], kt))

	cs.K = xif + hs/8*pf + hs/4*getP(tbl[0], kt)
	cs.M = hs/8*pf - xif
	cs.P = Params.H*Params.F0 + hs/4*(getF(tbl[0], kt)+ff)

	return cs
}

func getRConds(lt, kt interp.AkimaSpline, tbl FArr64) Conds {
	var cs Conds

	hs := Params.H * Params.H
	xif := getXi(lt.Predict(tbl[len(tbl)-1]), lt.Predict(tbl[len(tbl)-2]))
	pn := getP(tbl[len(tbl)-1], kt)
	pf := getXi(getP(tbl[len(tbl)-1], kt), getP(tbl[len(tbl)-2], kt))
	fn := getF(tbl[len(tbl)-1], kt)
	ff := getXi(getF(tbl[len(tbl)-1], kt), getF(tbl[len(tbl)-2], kt))

	cs.K = xif - hs/8*pf
	cs.M = -Params.H*Params.Alpha - xif - hs/8*pf - hs/4*pn
	cs.P = -Params.H*Params.Alpha*Params.T0 - hs/4*(fn+ff)

	return cs
}

func interpolate(xs, ys FArr64) interp.AkimaSpline {
	var as interp.AkimaSpline

	err := as.Fit(xs, ys)
	if err != nil {
		fmt.Println("Failed to initialize spline")
	}

	return as
}

func getP(x float64, kt interp.AkimaSpline) float64 {
	return 0
}

func getF(x float64, kt interp.AkimaSpline) float64 {
	return -4 * Params.Np * Params.Np * Params.Sigma * kt.Predict(x) * (math.Pow(x, 4) - math.Pow(Params.T0, 4))
}

func getXi(x1, x2 float64) float64 {
	return (x1 + x2) / 2
}

func getF1(tbl FArr64) float64 {
	return Params.F0 - Params.Alpha*(tbl[len(tbl)-1]-Params.T0)
}

func getF2(kt interp.AkimaSpline, tbl FArr64) float64 {
	x := make(FArr64, int((Params.L+Params.H)/Params.H))
	for i := 1; i < len(x); i++ {
		x[i] = x[i-1] + Params.H
	}
	y := make(FArr64, len(tbl))
	for i := 0; i < len(y); i++ {
		y[i] = kt.Predict(tbl[i]) * (math.Pow(tbl[i], 4) - math.Pow(Params.T0, 4))
	}
	return integrate.Simpsons(x, y)
}
