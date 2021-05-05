package emission

// FArr64 is used to represent []float64
type FArr64 []float64

// FMat64 is used to represent [][]float64
type FMat64 []FArr64

// Conds is used to represent emission system conditions
type Conds struct {
	K float64
	M float64
	P float64
}

// Emission is used to represent emission system parameters
type Emission struct {
	Np     float64
	L      float64
	T0     float64
	Tconst float64
	Sigma  float64
	F0     float64
	Alpha  float64
	H      float64
}
