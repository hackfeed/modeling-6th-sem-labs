package cauchy

// FArr64 used
type FArr64 []float64

// FMat64 used
type FMat64 []FArr64

// MakeFMat64 used
func MakeFMat64(n, m int) FMat64 {
	mat := make(FMat64, n)
	for i := 0; i < n; i++ {
		mat[i] = make(FArr64, m)
	}

	return mat
}
