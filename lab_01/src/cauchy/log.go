package cauchy

import (
	"fmt"
	"strings"

	"github.com/logrusorgru/aurora/v3"
)

// Log used
func Log(p FMat64, e, rk, x FArr64) {
	fmt.Printf("%200v\n", aurora.BgRed("CAUCHY PROBLEM SOLUTION"))
	fmt.Printf("%v%198v%v\n", "+", strings.Repeat("-", 198), "+")
	fmt.Printf(
		"|%27v|%27v|%27v|%27v|%28v|%28v|%28v|\n",
		aurora.Green("X"),
		aurora.Green("Picard, 1st approx."),
		aurora.Green("Picard, 2nd approx."),
		aurora.Green("Picard, 3rd approx."),
		aurora.Green("Picard, 4th approx."),
		aurora.Green("Euler"),
		aurora.Green("Runge-Kutta"),
	)
	fmt.Printf("%v%198v%v\n", "+", strings.Repeat("-", 198), "+")
	for i := 0; i < len(e); i++ {
		fmt.Printf(
			"|%27.5f|%27.5f|%27.5f|%27.5f|%28.5f|%28.5f|%28.5f|\n",
			x[i],
			p[0][i],
			p[1][i],
			p[2][i],
			p[3][i],
			e[i],
			rk[i],
		)
	}
	fmt.Printf("%v%36v%v\n", "+", strings.Repeat("-", 198), "+")
}
