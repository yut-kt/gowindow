package gowindow

import "math"

// Group3 is Other polynomial windows
// https://en.wikipedia.org/wiki/Window_function#Other_polynomial_windows

// Welch window consists of a single parabolic section.
// The defining quadratic polynomial reaches a value of zero at the samples just outside the span of the window.
func Welch(s []float64) []float64 {
	sw := append([]float64{}, s...)

	halfN := float64(len(s)-1) / 2
	for n := range s {
		sw[n] *= 1 - math.Pow((float64(n)-halfN)/halfN, 2)
	}
	return sw
}
