package circuit

// FArr64 is used to represent []float64
type FArr64 []float64

// FMat64 is used to represent [][]float64
type FMat64 []FArr64

func (m FMat64) GetColumn(n int) FArr64 {
	var (
		c  FArr64 = make(FArr64, 0)
		cs int    = len(m)
	)

	for i := 0; i < cs; i++ {
		c = append(c, m[i][n])
	}

	return c
}

// RCoeffs64 is used to represent Runge-Kutta coefficients
type RCoeffs64 struct {
	Kn float64
	Pn float64
}

// Circuit is used to represent circuit parameters
type Circuit struct {
	R   float64
	Le  float64
	Lk  float64
	Ck  float64
	Rk  float64
	Uc0 float64
	I0  float64
	Tw  float64
}
