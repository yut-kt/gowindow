package gowindow

import "math"

func blackmanHarris(s []float64) {
	v := 2.0 * math.Pi / float64(len(s)-1)
	for n := range s {
		val := v * float64(n)
		s[n] *= 0.35875 - 0.48829*math.Cos(val) + 0.14128*math.Cos(2*val) - 0.01168*math.Cos(3*val)
	}
}

func blackmanHarrisNew(s []float64) []float64 {
	sw := make([]float64, len(s))
	v := 2.0 * math.Pi / float64(len(s)-1)
	for n := range s {
		val := v * float64(n)
		sw[n] = s[n] * (0.35875 - 0.48829*math.Cos(val) + 0.14128*math.Cos(2*val) - 0.01168*math.Cos(3*val))
	}
	return sw
}
