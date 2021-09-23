package gowindow

// Group1 is Rectangular window.
// https://en.wikipedia.org/wiki/Window_function#Rectangular_window

// Rectangular window is the simplest window.
// w[n] = 1
// equivalent to replacing all but N values of a data sequence by zeros,
// making it appear as though the waveform suddenly turns on and off.
// The rectangular window is the 1st order B-spline window as well as the 0th power power-of-sine window.
// https://en.wikipedia.org/wiki/Window_function#Rectangular_window
func Rectangular(s []float64) []float64 {
	return append([]float64{}, s...)
}

// Boxcar window is the same window function as Rectangular.
func Boxcar(s []float64) []float64 {
	return Rectangular(s)
}

// Dirichlet window is the same window function as Rectangular.
func Dirichlet(s []float64) []float64 {
	return Rectangular(s)
}
