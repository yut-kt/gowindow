package gowindow

import "math"

func hanning(s []float64) {
	v := 2.0 * math.Pi / float64(len(s)-1)
	for n := range s {
		s[n] *= 0.5 - 0.5*math.Cos(v*float64(n))
	}
}

func hanningNew(s []float64) []float64 {
	sw := make([]float64, len(s))
	v := 2.0 * math.Pi / float64(len(s)-1)
	for n := range s {
		sw[n] = s[n] * (0.5 - 0.5*math.Cos(v*float64(n)))
	}
	return sw
}

func hann(s []float64) {
	hanning(s)
}

func hannNew(s []float64) []float64 {
	return hanningNew(s)
}
