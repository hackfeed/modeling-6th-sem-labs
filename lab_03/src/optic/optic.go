package optic

import (
	"fmt"
	"math"

	"gonum.org/v1/gonum/interp"
)

func GetRConds(lt, kt interp.AkimaSpline, tbl FArr64) Conds {
	var cs Conds

	cs.K = getXRight(lt, tbl, 0) +
		math.Pow(Params.H, 2)/8*getPRight(kt, tbl, 0) +
		math.Pow(Params.H, 2)/4*getP(kt, tbl, 0)
	cs.M = math.Pow(Params.H, 2)/8*getPRight(kt, tbl, 0) - getXRight(lt, tbl, 0)
	cs.P = Params.H*Params.F0 + math.Pow(Params.H, 2)/4*(getFRight(kt, tbl, 0)+getFLeft(kt, tbl, 0))

	return cs
}

func GetLConds(lt, kt interp.AkimaSpline, tbl FArr64, n int) Conds {
	var cs Conds

	cs.K = getXLeft(lt, tbl, n)/Params.H -
		Params.Alpha - Params.H*getP(kt, tbl, n)/4 -
		Params.H*getPLeft(kt, tbl, n)/8
	cs.M = getXLeft(lt, tbl, n)/Params.H - Params.H*getPLeft(kt, tbl, n)/8
	cs.P = -(Params.Alpha*Params.T0 + (getFRight(kt, tbl, n)+getFLeft(kt, tbl, n))/4*Params.H)

	return cs
}

func A(lt interp.AkimaSpline, tbl FArr64, n int) float64 {
	return (lt.Predict(tbl[n]) + lt.Predict(tbl[n-1])) / 2 / Params.H
}

func B(lt, kt interp.AkimaSpline, tbl FArr64, n int) float64 {
	return A(lt, tbl, n) + C(lt, tbl, n) + getP(kt, tbl, n)*Params.H
}

func C(lt interp.AkimaSpline, tbl FArr64, n int) float64 {
	return (lt.Predict(tbl[n]) + lt.Predict(tbl[n+1])) / 2 / Params.H
}

func D(kt interp.AkimaSpline, tbl FArr64, n int) float64 {
	return getF(kt, tbl, n) * Params.H
}

func Interpolate(xs, ys FArr64) interp.AkimaSpline {
	var as interp.AkimaSpline

	err := as.Fit(xs, ys)
	if err != nil {
		fmt.Println("Failed to initialize spline")
	}

	return as
}

func getP(kt interp.AkimaSpline, tbl FArr64, n int) float64 {
	return 4 * Params.Np * Params.Np * Params.Sigma * kt.Predict(tbl[n]) * math.Pow(tbl[n], 3)
}

func getF(kt interp.AkimaSpline, tbl FArr64, n int) float64 {
	return 4 * Params.Np * Params.Np * Params.Sigma * kt.Predict(tbl[n]) * math.Pow(Params.T0, 4)
}

func getXRight(lt interp.AkimaSpline, tbl FArr64, n int) float64 {
	return (lt.Predict(tbl[n]) + lt.Predict(tbl[n+1])) / 2
}

func getXLeft(lt interp.AkimaSpline, tbl FArr64, n int) float64 {
	return (lt.Predict(tbl[n]) + lt.Predict(tbl[n-1])) / 2
}

func getPRight(kt interp.AkimaSpline, tbl FArr64, n int) float64 {
	return (getP(kt, tbl, n) + getP(kt, tbl, n+1)) / 2
}

func getPLeft(kt interp.AkimaSpline, tbl FArr64, n int) float64 {
	return (getP(kt, tbl, n) + getP(kt, tbl, n-1)) / 2
}

func getFRight(kt interp.AkimaSpline, tbl FArr64, n int) float64 {
	return (getF(kt, tbl, n) + getF(kt, tbl, n+1)) / 2
}

func getFLeft(kt interp.AkimaSpline, tbl FArr64, n int) float64 {
	return (getF(kt, tbl, n) + getF(kt, tbl, n-1)) / 2
}
