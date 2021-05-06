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
	A1, B1, C1, M1, A2, B2, C2, M2, Alpha0, AlphaN, L, T0, R, F0, H, T, Eps float64
}
