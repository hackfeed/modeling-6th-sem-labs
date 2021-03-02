package main

import (
	"fmt"
	"math"
	"src/cauchy"
)

func main() {
	xs := 0.
	xe := 2.
	ys := 0.
	h := 1e-4
	n := int(math.Ceil(math.Abs(xe-xs) / h))

	x := make(cauchy.FArr64, 0)
	xn := xs
	for i := 0; i <= n; i++ {
		x = append(x, xn)
		xn += h
	}
	picardSol := cauchy.Picard(xs, h, n)
	eulerSol := cauchy.Euler(xs, ys, h, n)
	rungeKuttaSol := cauchy.RungeKutta(xs, ys, 0.5, h, n)

	fmt.Println()
	cauchy.Log(picardSol, eulerSol, rungeKuttaSol, x)
}
