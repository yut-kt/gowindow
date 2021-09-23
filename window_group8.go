package gowindow

import "math"

// Group8 is Other windows.
// https://en.wikipedia.org/wiki/Window_function#Other_windows

// Lanczos window used in Lanczos resampling.
// https://en.wikipedia.org/wiki/Window_function#Lanczos_window
func Lanczos(s []float64) []float64 {
	sw := append([]float64{}, s...)

	N := float64(len(s))
	halfN := (N - 1) / 2
	for n := range s {
		floatN := float64(n)
		if floatN != halfN {
			numerator := math.Sin(math.Pi * (2*floatN/(N-1) - 1))
			denominator := math.Pi * (2*floatN/(N-1) - 1)
			sw[n] *= numerator / denominator
		}
	}
	return sw
}
