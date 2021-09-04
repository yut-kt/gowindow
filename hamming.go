package gowindow

import "math"

// Hamming is func to apply Hamming Window
// Non-Destructive function
func Hamming(s []float64) []float64 {
	sw := make([]float64, len(s))
	v := 2.0 * math.Pi / float64(len(s)-1)
	for n := range s {
		sw[n] = s[n] * (0.54 - 0.46*math.Cos(v*float64(n)))
	}
	return sw
}

// HammingD is func to apply Hamming Window
// Non-Destructive function
func HammingD(s []float64) {
	v := 2.0 * math.Pi / float64(len(s)-1)
	for n := range s {
		s[n] *= 0.54 - 0.46*math.Cos(v*float64(n))
	}
}
