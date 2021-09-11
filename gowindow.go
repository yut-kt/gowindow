package gowindow

type window int

const (
	// Rectangular https://en.wikipedia.org/wiki/Window_function#Rectangular_window
	Rectangular window = iota
	// Triangular https://en.wikipedia.org/wiki/Window_function#Triangular_window
	Triangular
	// Bartlett https://en.wikipedia.org/wiki/Window_function#Triangular_window
	Bartlett
	// Fejer https://en.wikipedia.org/wiki/Window_function#Triangular_window
	Fejer
	// Parzen https://en.wikipedia.org/wiki/Window_function#Parzen_window
	Parzen
	// DeLaValleePoussin https://en.wikipedia.org/wiki/Window_function#Parzen_window
	DeLaValleePoussin
	// Welch https://en.wikipedia.org/wiki/Window_function#Welch_window
	Welch
	// Sine https://en.wikipedia.org/wiki/Window_function#Sine_window
	Sine
	// Hanning https://en.wikipedia.org/wiki/Window_function#Hann_and_Hamming_windows
	Hanning
	// Hann https://en.wikipedia.org/wiki/Window_function#Hann_and_Hamming_windows
	Hann
	// Hamming https://en.wikipedia.org/wiki/Window_function#Hann_and_Hamming_windows
	Hamming
	// Blackman https://en.wikipedia.org/wiki/Window_function#Blackman_window
	Blackman
	// Nuttall https://en.wikipedia.org/wiki/Window_function#Nuttall_window,_continuous_first_derivative
	Nuttall
)

// Apply func to apply window func Destructively
func Apply(s []float64, w window) {
	chooseApplyFunc(w)(s)
}

func chooseApplyFunc(w window) func([]float64) {
	switch w {
	case Rectangular:
		return rectangular
	case Triangular:
		return triangular
	case Bartlett:
		return bartlett
	case Fejer:
		return fejer
	case Parzen:
		return parzen
	case DeLaValleePoussin:
		return deLaValleePoussin
	case Welch:
		return welch
	case Sine:
		return sine
	case Hanning:
		return hanning
	case Hann:
		return hann
	case Hamming:
		return hamming
	case Blackman:
		return blackman
	case Nuttall:
		return nuttall
	}
	// return empty if unknown window
	return func(float64s []float64) {}
}

// ApplyNew func to apply window func Non-Destructively
func ApplyNew(s []float64, w window) []float64 {
	return chooseApplyNewFunc(w)(s)
}

func chooseApplyNewFunc(w window) func([]float64) []float64 {
	switch w {
	case Rectangular:
		return rectangularNew
	case Triangular:
		return triangularNew
	case Bartlett:
		return bartlettNew
	case Fejer:
		return fejerNew
	case Parzen:
		return parzenNew
	case DeLaValleePoussin:
		return deLaValleePoussinNew
	case Welch:
		return welchNew
	case Sine:
		return sineNew
	case Hanning:
		return hanningNew
	case Hann:
		return hannNew
	case Hamming:
		return hammingNew
	case Blackman:
		return blackmanNew
	case Nuttall:
		return nuttallNew
	}
	// return empty if unknown window
	return func(float64s []float64) []float64 { return []float64{} }
}
