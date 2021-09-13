package gowindow

import "math"

func exponential(s []float64, T float64) {
	N := len(s)
	for n := range s {
		s[n] *= math.Exp(-math.Abs(float64(n)-float64(N-1)/2.0) / T)
	}
}

func exponentialNew(s []float64, T float64) []float64 {
	N := len(s)
	sw := make([]float64, len(s))
	for n := range s {
		sw[n] = s[n] * math.Exp(-math.Abs(float64(n)-float64(N-1)/2.0)/T)
	}
	return sw
}

func poisson(s []float64, T float64) {
	exponential(s, T)
}

func poissonNew(s []float64, T float64) []float64 {
	return exponentialNew(s, T)
}
