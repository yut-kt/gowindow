package gowindow

import "math"

func blackmanNuttall(s []float64) {
	v := 2.0 * math.Pi / float64(len(s)-1)
	for n := range s {
		val := v * float64(n)
		s[n] *= 0.3635819 - 0.4891775*math.Cos(val) + 0.1365995*math.Cos(2*val) - 0.0106411*math.Cos(3*val)
	}
}

func blackmanNuttallNew(s []float64) []float64 {
	sw := make([]float64, len(s))
	v := 2.0 * math.Pi / float64(len(s)-1)
	for n := range s {
		val := v * float64(n)
		sw[n] = s[n] * (0.3635819 - 0.4891775*math.Cos(val) + 0.1365995*math.Cos(2*val) - 0.0106411*math.Cos(3*val))
	}
	return sw
}
