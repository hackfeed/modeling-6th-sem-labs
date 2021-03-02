package cauchy

import "math"

func fapprox(x float64) float64 {
	return x * x * x / 3
}

func sapprox(x float64) float64 {
	return fapprox(x) + math.Pow(x, 7)/63
}

func tapprox(x float64) float64 {
	return sapprox(x) + 2*math.Pow(x, 11)/2079 + math.Pow(x, 15)/59535
}

func foapprox(x float64) float64 {
	return sapprox(x) + 2*math.Pow(x, 11)/2079 + 13*math.Pow(x, 15)/218295 +
		82*math.Pow(x, 19)/37328445 + 662*math.Pow(x, 23)/10438212015 +
		4*math.Pow(x, 27)/3341878155 + math.Pow(x, 31)/109876902975
}

// Picard used to solve Cauchy problem with Picard method
func Picard(x0, h float64, n int) FMat64 {
	r := MakeFMat64(4, 0)

	for i := 0; i <= n; i++ {
		r[0] = append(r[0], fapprox(x0))
		r[1] = append(r[1], sapprox(x0))
		r[2] = append(r[2], tapprox(x0))
		r[3] = append(r[3], foapprox(x0))

		x0 += h
	}

	return r
}
