package gowindow

import "math"

// Hanning is func to apply Hanning Window
// Non-Destructive function
func Hanning(s []float64) []float64 {
	sw := make([]float64, len(s))
	v := 2.0 * math.Pi / float64(len(s)-1)
	for n := range s {
		sw[n] = s[n] * (0.5 - 0.5*math.Cos(v*float64(n)))
	}
	return sw
}

// Hann is func to apply Hanning Window
// Non-Destructive function
func Hann(s []float64) []float64 {
	return Hanning(s)
}

// HanningD is func to apply Hanning Window
// Destructive function
func HanningD(s []float64) {
	v := 2.0 * math.Pi / float64(len(s)-1)
	for n := range s {
		s[n] *= 0.5 - 0.5*math.Cos(v*float64(n))
	}
}

// HannD is func to apply Hanning Window
// Destructive function
func HannD(s []float64) {
	HanningD(s)
}
