package gowindow

import (
	"math"
)

func triangular(s []float64) {
	for k := range s {
		x := float64(k) / float64(len(s)-1)
		s[k] *= 1 - 2*math.Abs(x-0.5)
	}
}

func triangularNew(s []float64) []float64 {
	sw := make([]float64, len(s))
	for k := range s {
		x := float64(k) / float64(len(s)-1)
		sw[k] = s[k] * (1 - 2*math.Abs(x-0.5))
	}
	return sw
}

func bartlett(s []float64) {
	triangular(s)
}

func bartlettNew(s []float64) []float64 {
	return triangularNew(s)
}

func fejer(s []float64) {
	triangular(s)
}

func fejerNew(s []float64) []float64 {
	return triangularNew(s)
}
