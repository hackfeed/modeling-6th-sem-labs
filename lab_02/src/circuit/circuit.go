package circuit

import (
	"fmt"
	"math"

	"gonum.org/v1/gonum/integrate/quad"
	"gonum.org/v1/gonum/interp"
)

// GetT0 is used to find T0 parameter
func GetT0(I float64) float64 {
	return interpolate(I, CurTbl.GetColumn(0), CurTbl.GetColumn(1))
}

// GetT0 is used to find T0 parameter
func GetM(I float64) float64 {
	return interpolate(I, CurTbl.GetColumn(0), CurTbl.GetColumn(2))
}

// GetT0 is used to find T0 parameter
func GetRp(I, T0, m float64) float64 {
	f := func(x float64) float64 {
		return getSigma(getT(x, T0, m)) * x
	}
	val := quad.Fixed(f, 0, 1, 30, nil, 0)

	return Params.Le / (2 * math.Pi * Params.R * Params.R * val)
}

// GetT0 is used to find parameters with Runge-Kutta method
func GetRungeKutta(x, y, z, h, Rp float64) (float64, float64) {
	cfsArr := make([]RCoeffs64, Order)

	for i := 0; i < Order; i++ {
		v := i
		if i == 0 {
			v = Order - 1
		}
		_, yAdd, zAdd := getCurAdd(h, cfsArr[v], i, Order)
		cfsArr[i] = RCoeffs64{h * getF(y+yAdd, z+zAdd, Rp), h * getPhi(y+yAdd)}
	}

	return getNextMembs(y, z, cfsArr)
}

func getT(z, T0, m float64) float64 {
	return (Params.Tw-T0)*math.Pow(z, m) + T0
}

func getF(y, z, Rp float64) float64 {
	return -((Params.Rk+Rp)*y - z) / Params.Lk
}

func getPhi(y float64) float64 {
	return -y / Params.Ck
}

func getSigma(T float64) float64 {
	return interpolate(T, TmpTbl.GetColumn(0), TmpTbl.GetColumn(1))
}

func interpolate(x float64, xs, ys FArr64) float64 {
	var as interp.AkimaSpline

	err := as.Fit(xs, ys)
	if err != nil {
		fmt.Println("Failed to initialize spline")
	}

	return as.Predict(x)
}

func getCurAdd(h float64, cfs RCoeffs64, i, ord int) (float64, float64, float64) {
	if i == 0 {
		return 0, 0, 0
	}
	if i == ord-1 {
		return h, cfs.Kn, cfs.Pn
	}
	return h / 2, cfs.Kn / 2, cfs.Pn / 2
}

func getNextMembs(y, z float64, cfsArr []RCoeffs64) (float64, float64) {
	var (
		kSum float64 = 0
		pSum float64 = 0
		div  float64 = float64(2*(len(cfsArr)-2) + 2)
	)

	for i := 0; i < len(cfsArr); i++ {
		if i > 0 && i < len(cfsArr)-1 {
			kSum += 2 * cfsArr[i].Kn
			pSum += 2 * cfsArr[i].Pn
		} else {
			kSum += cfsArr[i].Kn
			pSum += cfsArr[i].Pn
		}
	}

	return y + kSum/div, z + pSum/div
}
