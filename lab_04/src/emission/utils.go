package emission

import "math"

// Arange is used to model numpy.arange behaviour
func Arange(start, stop, step float64) []float64 {
	n := int(math.Ceil((stop - start) / step))
	rnge := make([]float64, n)
	for x := range rnge {
		rnge[x] = start + step*float64(x)
	}
	return rnge
}
