package gowindow

import "math"

// Group7 is Hybrid windows.
// Window functions have also been constructed as multiplicative or additive combinations of other windows.
// https://en.wikipedia.org/wiki/Window_function#Hybrid_windows

// BartlettHann is hybrid Bartlett and Hann.
// https://en.wikipedia.org/wiki/Window_function#Bartlett%E2%80%93Hann_window
func BartlettHann(s []float64) []float64 {
	sw := append([]float64{}, s...)

	N := float64(len(s))
	v := 2 * math.Pi / (N - 1)
	for n := range s {
		nFloat := float64(n)
		sw[n] *= 0.62 - 0.48*math.Abs(nFloat/(N-1)-0.5) - 0.38*math.Cos(nFloat*v)
	}
	return sw
}

// HannPoisson is Hann window multiplied by a Poisson window, which has no side-lobes,
// in the sense that (for alpha[α] >= 2) its Fourier transform drops off forever away from the main lobe.
// It can thus be used in hill climbing algorithms like Newton's method.
// https://en.wikipedia.org/wiki/Window_function#Hann%E2%80%93Poisson_window
func HannPoisson(s []float64, alpha float64) []float64 {
	sw := append([]float64{}, s...)

	halfN := float64(len(s)-1) / 2
	for k := range s {
		sw[k] *= 0.5 * (1 - math.Cos(math.Pi*float64(k)/halfN)) * math.Exp(-alpha*math.Abs(float64(k)-halfN)/halfN)
	}
	return sw
}
