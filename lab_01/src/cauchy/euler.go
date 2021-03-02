package cauchy

// Euler used
func Euler(x0, y0, h float64, n int) FArr64 {
	r := make(FArr64, 0)

	for i := 0; i <= n; i++ {
		r = append(r, y0)
		y0 += h * equation(x0, y0)
		x0 += h
	}

	return r
}
