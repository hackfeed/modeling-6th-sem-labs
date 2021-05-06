package emission

import "math"

func SimpleIters() (FMat64, float64) {
	tbl := make(FArr64, int(Params.L/Params.H)+1)
	for i := 0; i < len(tbl); i++ {
		tbl[i] = Params.T0
	}
	ntbl := make(FArr64, int(Params.L/Params.H)+1)

	res := FMat64{tbl}
	ti := 0.
	fl := true

	for fl {
		ptbl := tbl
		cmax := 1.

		for cmax >= 1 {
			ntbl = getT(ptbl)
			cmax = math.Abs((tbl[0] - ntbl[0]) / ntbl[0])

			for i := 0; i < len(tbl); i++ {
				d := math.Abs((tbl[i] - ntbl[i]) / ntbl[i])
				if d > cmax {
					cmax = d
				}
			}

			ptbl = ntbl
		}

		res = append(res, ntbl)
		ti += Params.T

		fl = false

		for i := 0; i < len(tbl); i++ {
			if math.Abs((tbl[i]-ntbl[i])/ntbl[i]) > Params.Eps {
				fl = true
			}
		}

		tbl = ntbl
	}

	return res, ti
}

func getT(tbl FArr64) FArr64 {
	lcs := getLConds(tbl)
	rcs := getRConds(tbl)

	xil := FArr64{0, -lcs.M / lcs.K}
	etal := FArr64{0, lcs.P / lcs.K}

	x := Params.H
	n := 1

	for x+Params.H < Params.L {
		tn := tbl[n]
		den := getBCf(x, tn) - getACf(tn)*xil[n]

		xil = append(xil, getDCf(tn)/den)
		etal = append(etal, (getFCf(x, tn)+getACf(tn)*etal[n])/den)

		n++
		x += Params.H
	}

	ntbl := make(FArr64, n+1)
	ntbl[n] = (rcs.P - rcs.M*etal[n]) / (rcs.K + rcs.M*xil[n])

	for i := n - 1; i > -1; i-- {
		ntbl[i] = xil[i+1]*ntbl[i+1] + etal[i+1]
	}

	return ntbl
}

func getLConds(tbl FArr64) Conds {
	var lcs Conds

	cp := getApproxPlus(getC, tbl[0], Params.T)
	kp := getApproxPlus(getK, tbl[0], Params.T)

	lcs.K = Params.H/8*cp + Params.H/4*getC(tbl[0]) + Params.T/Params.H*kp +
		Params.T*Params.H/8*getP(Params.H/2) + Params.T*Params.H/4*getP(0)
	lcs.M = Params.H/8*cp - Params.T/Params.H*kp + Params.T*Params.H/8*getP(Params.H/2)
	lcs.P = Params.H/8*cp*(tbl[0]+tbl[1]) + Params.H/4*getC(tbl[0])*tbl[0] +
		Params.F0*Params.T + Params.T*Params.H/8*(3*getF(0)+getF(Params.H))

	return lcs
}

func getRConds(tbl FArr64) Conds {
	var rcs Conds

	cm := getApproxMinus(getC, tbl[len(tbl)-1], Params.T)
	km := getApproxMinus(getK, tbl[len(tbl)-1], Params.T)

	rcs.K = Params.H/8*cm + Params.H/4*getC(tbl[len(tbl)-1]) + Params.T/Params.H*km +
		Params.T*Params.AlphaN + Params.T*Params.H/8*getP(Params.L-Params.H/2) +
		Params.T*Params.H/4*getP(Params.L)
	rcs.M = Params.H/8*cm - Params.T/Params.H*km + Params.T*Params.H/8*getP(Params.L-Params.H/2)
	rcs.P = Params.H/8*cm*(tbl[len(tbl)-1]+tbl[len(tbl)-2]) + Params.H/4*getC(tbl[len(tbl)-1])*tbl[len(tbl)-1] +
		Params.T*Params.AlphaN*Params.T0 + Params.T*Params.H/4*(getF(Params.L)+getF(Params.L-Params.H/2))

	return rcs
}

func getBCf(m, x float64) float64 {
	return getACf(x) + getDCf(x) + Params.H*getC(x) + Params.H*Params.T*getP(m)
}

func getFCf(m, x float64) float64 {
	return Params.H*Params.T*getF(m) + x*Params.H*getC(x)
}

func getACf(x float64) float64 {
	return Params.T / Params.H * getApproxMinus(getK, x, Params.T)
}

func getDCf(x float64) float64 {
	return Params.T / Params.H * getApproxPlus(getK, x, Params.T)
}

func getK(x float64) float64 {
	return Params.A1 * (Params.B1 + Params.C1*math.Pow(x, Params.M1))
}

func getC(x float64) float64 {
	return Params.A2 + Params.B2*math.Pow(x, Params.M2) - Params.C2/x/x
}

func getP(x float64) float64 {
	return getAlpha(x) * 2 / Params.R
}

func getF(x float64) float64 {
	return getAlpha(x) * 2 * Params.T0 / Params.R
}

func getAlpha(x float64) float64 {
	d := (Params.AlphaN * Params.L) / (Params.AlphaN - Params.Alpha0)
	c := -Params.Alpha0 * d
	return c / (x - d)
}

func getApproxPlus(fn func(float64) float64, n, st float64) float64 {
	return (fn(n) + fn(n+st)) / 2
}

func getApproxMinus(fn func(float64) float64, n, st float64) float64 {
	return (fn(n) + fn(n-st)) / 2
}
