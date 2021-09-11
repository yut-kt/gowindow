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
	// BlackmanNuttall https://en.wikipedia.org/wiki/Window_function#Blackman%E2%80%93Nuttall_window
	BlackmanNuttall
	// BlackmanHarris https://en.wikipedia.org/wiki/Window_function#Blackman%E2%80%93Harris_window
	BlackmanHarris
	// None is test window for when missed in switch implementation
	None
)

func (w window) applyWindow(s []float64) {
	switch w {
	case Rectangular:
		rectangular(s)
	case Triangular:
		triangular(s)
	case Bartlett:
		bartlett(s)
	case Fejer:
		fejer(s)
	case Parzen:
		parzen(s)
	case DeLaValleePoussin:
		deLaValleePoussin(s)
	case Welch:
		welch(s)
	case Sine:
		sine(s)
	case Hanning:
		hanning(s)
	case Hann:
		hann(s)
	case Hamming:
		hamming(s)
	case Blackman:
		blackman(s)
	case Nuttall:
		nuttall(s)
	case BlackmanNuttall:
		blackmanNuttall(s)
	case BlackmanHarris:
		blackmanHarris(s)
	}
}

func (w window) applyNewWindow(s []float64) []float64 {
	switch w {
	case Rectangular:
		return rectangularNew(s)
	case Triangular:
		return triangularNew(s)
	case Bartlett:
		return bartlettNew(s)
	case Fejer:
		return fejerNew(s)
	case Parzen:
		return parzenNew(s)
	case DeLaValleePoussin:
		return deLaValleePoussinNew(s)
	case Welch:
		return welchNew(s)
	case Sine:
		return sineNew(s)
	case Hanning:
		return hanningNew(s)
	case Hann:
		return hannNew(s)
	case Hamming:
		return hammingNew(s)
	case Blackman:
		return blackmanNew(s)
	case Nuttall:
		return nuttallNew(s)
	case BlackmanNuttall:
		return blackmanNuttallNew(s)
	case BlackmanHarris:
		return blackmanHarrisNew(s)
	}
	// missed in switch implementation
	return []float64{}
}

// Apply func to apply window func Destructively
func Apply(s []float64, w window) {
	w.applyWindow(s)
}

// ApplyNew func to apply window func Non-Destructively
func ApplyNew(s []float64, w window) []float64 {
	return w.applyNewWindow(s)
}
