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
	// FlatTop https://en.wikipedia.org/wiki/Window_function#Flat_top_window
	FlatTop
	// RifeVincentClass1Order1 Class I, Order 1 (K = 1) https://en.wikipedia.org/wiki/Window_function#Rife%E2%80%93Vincent_windows
	// Class I is defined by minimizing the high-order side-lobe amplitude
	// Functionally equivalent to the Hann window
	RifeVincentClass1Order1
	// RifeVincentClass1Order2 Class I, Order 2 (K = 2) https://en.wikipedia.org/wiki/Window_function#Rife%E2%80%93Vincent_windows
	RifeVincentClass1Order2
	// RifeVincentClass1Order3 Class I, Order 3 (K = 3) https://en.wikipedia.org/wiki/Window_function#Rife%E2%80%93Vincent_windows
	RifeVincentClass1Order3
	// RifeVincentClass1Order4 Class I, Order 4 (K = 4) https://en.wikipedia.org/wiki/Window_function#Rife%E2%80%93Vincent_windows
	RifeVincentClass1Order4
	// RifeVincentClass2Decibel36 https://en.wikipedia.org/wiki/Window_function#Rife%E2%80%93Vincent_windows
	// Class II minimizes the main-lobe width for a given maximum side-lobe
	RifeVincentClass2Decibel36
	// RifeVincentClass2Decibel42 https://en.wikipedia.org/wiki/Window_function#Rife%E2%80%93Vincent_windows
	RifeVincentClass2Decibel42
	// RifeVincentClass2Decibel48 https://en.wikipedia.org/wiki/Window_function#Rife%E2%80%93Vincent_windows
	RifeVincentClass2Decibel48
	// RifeVincentClass2Decibel54 https://en.wikipedia.org/wiki/Window_function#Rife%E2%80%93Vincent_windows
	RifeVincentClass2Decibel54
	// RifeVincentClass2Decibel60 https://en.wikipedia.org/wiki/Window_function#Rife%E2%80%93Vincent_windows
	RifeVincentClass2Decibel60
	// RifeVincentClass2Decibel66 https://en.wikipedia.org/wiki/Window_function#Rife%E2%80%93Vincent_windows
	RifeVincentClass2Decibel66
	// RifeVincentClass3Order2 Class III, Order 2 (k = 2) https://en.wikipedia.org/wiki/Window_function#Rife%E2%80%93Vincent_windows
	// Class III is a compromise for which order K = 2 resembles the § Blackman window
	RifeVincentClass3Order2
	// RifeVincentClass3Order3 Class III, Order 3 (k = 3) https://en.wikipedia.org/wiki/Window_function#Rife%E2%80%93Vincent_windows
	RifeVincentClass3Order3
	// RifeVincentClass3Order4 Class III, Order 4 (k = 4) https://en.wikipedia.org/wiki/Window_function#Rife%E2%80%93Vincent_windows
	RifeVincentClass3Order4
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
	case FlatTop:
		flatTop(s)
	case RifeVincentClass1Order1:
		rifeVincent(s, class1, order1, nonDecibel)
	case RifeVincentClass1Order2:
		rifeVincent(s, class1, order2, nonDecibel)
	case RifeVincentClass1Order3:
		rifeVincent(s, class1, order3, nonDecibel)
	case RifeVincentClass1Order4:
		rifeVincent(s, class1, order4, nonDecibel)
	case RifeVincentClass2Decibel36:
		rifeVincent(s, class2, nonOrder, dB36)
	case RifeVincentClass2Decibel42:
		rifeVincent(s, class2, nonOrder, dB42)
	case RifeVincentClass2Decibel48:
		rifeVincent(s, class2, nonOrder, dB48)
	case RifeVincentClass2Decibel54:
		rifeVincent(s, class2, nonOrder, dB54)
	case RifeVincentClass2Decibel60:
		rifeVincent(s, class2, nonOrder, dB60)
	case RifeVincentClass2Decibel66:
		rifeVincent(s, class2, nonOrder, dB66)
	case RifeVincentClass3Order2:
		rifeVincent(s, class3, order2, nonDecibel)
	case RifeVincentClass3Order3:
		rifeVincent(s, class3, order3, nonDecibel)
	case RifeVincentClass3Order4:
		rifeVincent(s, class3, order4, nonDecibel)
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
	case FlatTop:
		return flatTopNew(s)
	case RifeVincentClass1Order1:
		return rifeVincentNew(s, class1, order1, nonDecibel)
	case RifeVincentClass1Order2:
		return rifeVincentNew(s, class1, order2, nonDecibel)
	case RifeVincentClass1Order3:
		return rifeVincentNew(s, class1, order3, nonDecibel)
	case RifeVincentClass1Order4:
		return rifeVincentNew(s, class1, order4, nonDecibel)
	case RifeVincentClass2Decibel36:
		return rifeVincentNew(s, class2, nonOrder, dB36)
	case RifeVincentClass2Decibel42:
		return rifeVincentNew(s, class2, nonOrder, dB42)
	case RifeVincentClass2Decibel48:
		return rifeVincentNew(s, class2, nonOrder, dB48)
	case RifeVincentClass2Decibel54:
		return rifeVincentNew(s, class2, nonOrder, dB54)
	case RifeVincentClass2Decibel60:
		return rifeVincentNew(s, class2, nonOrder, dB60)
	case RifeVincentClass2Decibel66:
		return rifeVincentNew(s, class2, nonOrder, dB66)
	case RifeVincentClass3Order2:
		return rifeVincentNew(s, class3, order2, nonDecibel)
	case RifeVincentClass3Order3:
		return rifeVincentNew(s, class3, order3, nonDecibel)
	case RifeVincentClass3Order4:
		return rifeVincentNew(s, class3, order4, nonDecibel)
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
