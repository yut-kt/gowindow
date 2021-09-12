package gowindow

import "errors"

type window struct {
	w windows
	o *Option
}

// Option use as an argument to SetOption
type Option struct {
	// only Rife-Vincent window
	Class   class
	Order   order
	Decibel decibel
	// only Gaussian and GeneralizedNormal window
	SD float64 // σ
	// only ConfinedGaussian window
	SDt float64
	// only GeneralizedNormal window
	P int
	// only Tukey and Kaiser window
	Alpha float64 // α
	// only PlanckTaper window
	Epsilon float64 // ε
}

type windows int

const (
	// Rectangular https://en.wikipedia.org/wiki/Window_function#Rectangular_window
	Rectangular windows = iota
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
	// RifeVincent https://en.wikipedia.org/wiki/Window_function#Rife%E2%80%93Vincent_windows
	RifeVincent
	// Gaussian https://en.wikipedia.org/wiki/Window_function#Gaussian_window
	Gaussian
	// Gauss https://en.wikipedia.org/wiki/Window_function#Gaussian_window
	Gauss
	// ConfinedGaussian https://en.wikipedia.org/wiki/Window_function#Confined_Gaussian_window
	ConfinedGaussian
	// ApproximateConfinedGaussian https://en.wikipedia.org/wiki/Window_function#Approximate_confined_Gaussian_window
	ApproximateConfinedGaussian
	// GeneralizedNormal https://en.wikipedia.org/wiki/Window_function#Generalized_normal_window
	GeneralizedNormal
	// Tukey https://en.wikipedia.org/wiki/Window_function#Tukey_window
	Tukey
	// PlanckTaper https://en.wikipedia.org/wiki/Window_function#Planck-taper_window
	PlanckTaper
	// Kaiser https://en.wikipedia.org/wiki/Window_function#Kaiser_window
	Kaiser
	// None is test window for when missed in switch implementation
	None
)

// New is constructor
func New(w windows) *window {
	return &window{w: w}
}

// SetOption is use when external options are needed
func (w *window) SetOption(o *Option) error {
	if err := w.validateOption(o); err != nil {
		return err
	}
	w.o = o
	return nil
}

func (w *window) validateOption(o *Option) error {
	switch w.w {
	case RifeVincent:
		// TODO: validation
	case Gaussian:
		if o.SD > 0.5 {
			return errors.New("SD is illegal value")
		}
	case ApproximateConfinedGaussian:
		if o.SDt >= 0.14 {
			return errors.New("SDt is illegal value")
		}
	case PlanckTaper:
		if o.Epsilon <= 0 || o.Epsilon > 0.5 {
			return errors.New("epsilon is illegal value")
		}
	}
	return nil
}

func (w *window) applyWindow(s []float64) {
	switch w.w {
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
	case RifeVincent:
		rifeVincent(s, w.o.Class, w.o.Order, w.o.Decibel)
	case Gaussian:
		gaussian(s, w.o.SD)
	case Gauss:
		gauss(s, w.o.SD)
	case ConfinedGaussian:
		confinedGaussian(s, w.o.SDt)
	case ApproximateConfinedGaussian:
		approximateConfinedGaussian(s, w.o.SDt)
	case GeneralizedNormal:
		generalizedNormal(s, w.o.SD, float64(w.o.P))
	case Tukey:
		tukey(s, w.o.Alpha)
	case PlanckTaper:
		planckTaper(s, w.o.Epsilon)
	case Kaiser:
		kaiser(s, w.o.Alpha)
	}
}

func (w window) applyNewWindow(s []float64) []float64 {
	switch w.w {
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
	case RifeVincent:
		return rifeVincentNew(s, w.o.Class, w.o.Order, w.o.Decibel)
	case Gaussian:
		return gaussianNew(s, w.o.SD)
	case Gauss:
		return gaussNew(s, w.o.SD)
	case ConfinedGaussian:
		return confinedGaussianNew(s, w.o.SDt)
	case ApproximateConfinedGaussian:
		return approximateConfinedGaussianNew(s, w.o.SDt)
	case GeneralizedNormal:
		return generalizedNormalNew(s, w.o.SD, float64(w.o.P))
	case Tukey:
		return tukeyNew(s, w.o.Alpha)
	case PlanckTaper:
		return planckTaperNew(s, w.o.Epsilon)
	case Kaiser:
		return kaiserNew(s, w.o.Alpha)
	}
	// missed in switch implementation
	return []float64{}
}

// Apply func to apply window func Destructively
func (w *window) Apply(s []float64) {
	w.applyWindow(s)
}

// ApplyNew func to apply window func Non-Destructively
func (w *window) ApplyNew(s []float64) []float64 {
	return w.applyNewWindow(s)
}
