package gowindow

import "math"

func blackman(s []float64) {
	v := 2.0 * math.Pi / float64(len(s)-1)
	for n := range s {
		val := v * float64(n)
		s[n] *= 0.42 - 0.5*math.Cos(val) + 0.08*math.Cos(2*val)
	}
}

func blackmanNew(s []float64) []float64 {
	sw := make([]float64, len(s))
	v := 2.0 * math.Pi / float64(len(s)-1)
	for n := range s {
		val := v * float64(n)
		sw[n] = s[n] * (0.42 - 0.5*math.Cos(val) + 0.08*math.Cos(2*val))
	}
	return sw
}
