package gowindow

import "math"

// Group2 is B-spline windows
// https://en.wikipedia.org/wiki/Window_function#B-spline_windows
// B-spline windows can be obtained as k-fold convolutions of the rectangular window.
// They include the rectangular window itself (K = 1),
// the § Triangular window (K = 2) and
// the § Parzen window (K = 4).
// Alternative definitions sample the appropriate normalized B-spline basis functions
// instead of convolving discrete-time windows.
// A Kth-order B-spline basis function is a piece-wise polynomial function of degree K−1
// that is obtained by k-fold self-convolution of the rectangular function.

// Triangular windows is the 2nd order B-spline window.
// The L = N form can be seen as the convolution of two N/2-width rectangular windows.
// The Fourier transform of the result is the squared values of the transform of the half-width rectangular window.
//https://en.wikipedia.org/wiki/Window_function#Triangular_window
func Triangular(s []float64) []float64 {
	sw := append([]float64{}, s...)

	halfN := float64(len(s)-1) / 2
	for n := range s {
		sw[n] *= 1 - math.Abs((float64(n)-halfN)/halfN)
	}
	return sw
}

// Bartlett window is the same window function as Triangular.
func Bartlett(s []float64) []float64 {
	return Triangular(s)
}

// Fejer window is the same window function as Triangular.
func Fejer(s []float64) []float64 {
	return Triangular(s)
}

// Parzen window is the 4th order B-spline window.
// Not to be confused with Kernel density estimation.
func Parzen(s []float64) []float64 {
	sw := append([]float64{}, s...)

	N := float64(len(s) - 1)
	halfN := N / 2
	for n := range s {
		nFloat := float64(n)
		v1 := math.Abs(nFloat - halfN)
		v2 := v1 / N
		if v1 <= halfN {
			sw[n] *= 1 - 6*math.Pow(v2, 2) + 6*math.Pow(v2, 3)
		} else {
			sw[n] *= 2 * math.Pow(1-v2, 3)
		}
	}
	return sw
}

// DeLaValleePoussin window is the same window function as Parzen.
func DeLaValleePoussin(s []float64) []float64 {
	return Parzen(s)
}
