package gowindow

import "math"

// Group4 is Sine window
// https://en.wikipedia.org/wiki/Window_function#Sine_window

// PowerOfSine window functions have the form:
// w[n]=sin^α(πn/N): The arg alpha is α
// https://en.wikipedia.org/wiki/Window_function#Power-of-sine/cosine_windows
func PowerOfSine(s []float64, alpha float64) []float64 {
	sw := append([]float64{}, s...)

	v := math.Pi / float64(len(s)-1)
	for n := range s {
		sw[n] *= math.Pow(math.Sin(float64(n)*v), alpha)
	}
	return sw
}

// PowerOfCosine window functions have the form:
// w[n]=cos^α(πn/N-π/2): The arg alpha is α
// https://en.wikipedia.org/wiki/Window_function#Power-of-sine/cosine_windows
func PowerOfCosine(s []float64, alpha float64) []float64 {
	sw := append([]float64{}, s...)

	v := math.Pi / float64(len(s)-1)
	halfPi := math.Pi / 2.0
	for n := range s {
		sw[n] *= math.Pow(math.Cos(float64(n)*v-halfPi), alpha)
	}
	return sw
}

// Sine window is sometimes also called cosine window.
// https://en.wikipedia.org/wiki/Window_function#Sine_window
func Sine(s []float64) []float64 {
	return PowerOfSine(s, 1)
}

// Cosine window is the same window function as Sine.
// https://en.wikipedia.org/wiki/Window_function#Sine_window
func Cosine(s []float64) []float64 {
	return PowerOfCosine(s, 1)
}

// Vorbis window used [Vorbis](https://xiph.org/vorbis/doc/Vorbis_I_spec.html)
// https://en.wikipedia.org/wiki/Vorbis
func Vorbis(s []float64) []float64 {
	sw := append([]float64{}, s...)

	halfPi := math.Pi / 2
	v := math.Pi / float64(len(s)-1)
	for n := range s {
		sw[n] *= math.Sin(halfPi*math.Pow(math.Sin(float64(n)*v), 2))
	}
	return sw
}
