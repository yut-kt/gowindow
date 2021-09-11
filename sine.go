package gowindow

import "math"

func sine(s []float64) {
	N := float64(len(s) - 1)
	for n := range s {
		s[n] *= math.Sin(math.Pi * float64(n) / N)
	}
}

func sineNew(s []float64) []float64 {
	sw := make([]float64, len(s))
	N := float64(len(s) - 1)
	for n := range s {
		sw[n] = s[n] * math.Sin(math.Pi*float64(n)/N)
	}
	return sw
}
