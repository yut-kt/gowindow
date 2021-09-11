package gowindow

import "math"

func nuttall(s []float64) {
	v := 2.0 * math.Pi / float64(len(s)-1)
	for n := range s {
		val := v * float64(n)
		s[n] *= 0.355768 - 0.487396*math.Cos(val) + 0.144232*math.Cos(2*val) - 0.012604*math.Cos(3*val)
	}
}

func nuttallNew(s []float64) []float64 {
	sw := make([]float64, len(s))
	v := 2.0 * math.Pi / float64(len(s)-1)
	for n := range s {
		val := v * float64(n)
		sw[n] = s[n] * (0.355768 - 0.487396*math.Cos(val) + 0.144232*math.Cos(2*val) - 0.012604*math.Cos(3*val))
	}
	return sw
}
