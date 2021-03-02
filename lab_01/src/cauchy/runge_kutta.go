package cauchy

// RungeKutta used
func RungeKutta(x0, y0, a, h float64, n int) FArr64 {
	r := make(FArr64, 0)

	for i := 0; i <= n; i++ {
		r = append(r, y0)
		y0 += h * ((1-a)*equation(x0, y0) + a*equation(x0+h/2/a, y0+h*equation(x0, y0)/2/a))
		x0 += h
	}

	return r
}
