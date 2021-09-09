package gowindow

import "math"

func hamming(s []float64) {
	v := 2.0 * math.Pi / float64(len(s)-1)
	for n := range s {
		s[n] *= 0.54 - 0.46*math.Cos(v*float64(n))
	}
}

func hammingNew(s []float64) []float64 {
	sw := make([]float64, len(s))
	v := 2.0 * math.Pi / float64(len(s)-1)
	for n := range s {
		sw[n] = s[n] * (0.54 - 0.46*math.Cos(v*float64(n)))
	}
	return sw
}
