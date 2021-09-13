package gowindow

import "math"

func powerOfSine(s []float64, alpha float64) {
	v := math.Pi / float64(len(s)-1)
	for n := range s {
		s[n] *= math.Pow(math.Sin(float64(n)*v), alpha)
	}
}

func powerOfSineNew(s []float64, alpha float64) []float64 {
	v := math.Pi / float64(len(s)-1)
	sw := make([]float64, len(s))
	for n := range s {
		sw[n] = s[n] * math.Pow(math.Sin(float64(n)*v), alpha)
	}
	return sw
}
