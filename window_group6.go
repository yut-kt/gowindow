package gowindow

import (
	"errors"
	"math"
)

// Group6 is adjustable windows.
// https://en.wikipedia.org/wiki/Window_function#Adjustable_windows

type gaussianFunction struct {
	SDt   float64
	L     float64
	HalfN float64
}

// G is Gaussian function
func (g *gaussianFunction) G(x float64) float64 {
	return math.Exp(-math.Pow((x-g.HalfN)/2*g.L*g.SDt, 2))
}

// Gaussian window is since the support of a Gaussian function extends to infinity,
// it must either be truncated at the ends of the window, or itself windowed with another zero-ended window.
// Since the log of Gaussian produces a parabola, this can be used for nearly exact quadratic interpolation in frequency estimation.
// sd[σ] <= 0.5
// https://en.wikipedia.org/wiki/Window_function#Gaussian_window
func Gaussian(s []float64, sd float64) ([]float64, error) {
	if sd > 0.5 {
		return nil, errors.New("invalid error: sd <= 0.5")
	}
	sw := append([]float64{}, s...)

	halfN := float64(len(s)-1) / 2.0
	for n := range s {
		sw[n] *= math.Exp(-0.5 * math.Pow((float64(n)-halfN)/(sd*halfN), 2))
	}
	return sw, nil
}

// Gauss window is the same window function as Gaussian.
func Gauss(s []float64, sd float64) ([]float64, error) {
	return Gaussian(s, sd)
}

// ConfinedGaussian window yields the smallest possible root-mean-square frequency width σω for a given temporal width (N + 1) SDt[σt].
// These windows optimize the RMS time-frequency bandwidth products.
// They are computed as the minimum eigenvectors of a parameter-dependent matrix.
// The ConfinedGaussian window family contains the § Sine window and the § Gaussian window in the limiting cases of large and small SDt[σt], respectively.
// https://en.wikipedia.org/wiki/Window_function#Confined_Gaussian_window
func ConfinedGaussian(s []float64, SDt float64) []float64 {
	sw := append([]float64{}, s...)

	L := float64(len(s) + 1)
	g := gaussianFunction{SDt: SDt, L: L, HalfN: float64(len(s)-1) / 2}
	for n := range s {
		nFloat := float64(n)
		numerator := g.G(nFloat+L) + g.G(nFloat-L)
		denominator := g.G(-0.5+L) + g.G(-0.5-L)
		sw[n] *= g.G(nFloat) - g.G(-0.5)*(numerator/denominator)
	}
	return sw
}

// ApproximateConfinedGaussian is ConfinedGaussian window of temporal width L×SDt[σt] is well approximated.
// The standard deviation of the approximate window is asymptotically equal (i.e. large values of N) to L×σt for σt < 0.14.
// https://en.wikipedia.org/wiki/Window_function#Approximate_confined_Gaussian_window
func ApproximateConfinedGaussian(s []float64, SDt float64) []float64 {
	sw := append([]float64{}, s...)

	L := float64(len(s) + 1)
	g := gaussianFunction{SDt: SDt, L: L, HalfN: float64(len(s)-1) / 2.0}
	for n := range s {
		nFloat := float64(n)
		numerator := g.G(-0.5) * (g.G(nFloat+L) + g.G(nFloat-L))
		denominator := g.G(-0.5+L) + g.G(-0.5-L)
		sw[n] *= g.G(nFloat) - numerator/denominator
	}
	return sw
}

// GeneralizedNormal is more generalized version of the Gaussian window.
// For any even p.At p=2, this is a Gaussian window and as p approaches ∞, this approximates to a rectangular window.
// The Fourier transform of this window does not exist in a closed form for a general p.
// However, it demonstrates the other benefits of being smooth, adjustable bandwidth.
// Like the § Tukey window, this window naturally offers a "flat top" to control the amplitude attenuation of a time-series
// (on which we don't have a control with Gaussian window).
// In essence, it offers a good (controllable) compromise, in terms of spectral leakage,
// frequency resolution and amplitude attenuation, between the Gaussian window and the rectangular window.
// https://en.wikipedia.org/wiki/Window_function#Generalized_normal_window
func GeneralizedNormal(s []float64, sd, p float64) []float64 {
	sw := append([]float64{}, s...)

	halfN := float64(len(s)-1) / 2.0
	for n := range s {
		sw[n] *= math.Exp(-math.Pow((float64(n)-halfN)/(sd*halfN), p))
	}
	return sw
}

// Tukey window, also known as the cosine-tapered window,
// can be regarded as a cosine lobe of width Nα/2 (spanning Nα/2 + 1 observations) that is convolved with a rectangular window of width N(1 − α/2).
// https://en.wikipedia.org/wiki/Window_function#Tukey_window
func Tukey(s []float64, alpha float64) []float64 {
	sw := append([]float64{}, s...)

	halfN := float64(len(s)-1) / 2
	alphaHalfN := alpha * halfN
	for n := range s {
		abs := math.Abs(float64(n) - halfN)
		if abs >= alphaHalfN {
			sw[n] *= 0.5 + 0.5*math.Cos(math.Pi*(abs-alphaHalfN)/(halfN-alphaHalfN))
		}
	}
	return sw
}

// PlanckTaper window is a bump function that has been widely used in the theory of partitions of unity in manifolds.
// It is smooth (a C^∞ function) everywhere, but is exactly zero out-side of a compact region,
// exactly one over an interval within that region, and varies smoothly and monotonically between those limits.
// Its use as a window function in signal processing was first suggested in the context of gravitational-wave astronomy,
// inspired by the Planck distribution.
// https://en.wikipedia.org/wiki/Window_function#Planck-taper_window
func PlanckTaper(s []float64, epsilon float64) []float64 {
	sw := append([]float64{}, s...)

	N := float64(len(s))
	Za := func(k float64) float64 {
		return epsilon * (N - 1) * (1.0/k + 1/(k-epsilon*(N-1)))
	}
	Zb := func(k float64) float64 {
		return epsilon * (N - 1) * (1/(N-1-k) + 1/((1-epsilon)*(N-1)-k))

	}
	point1 := epsilon * (N - 1)
	point2 := (1 - epsilon) * (N - 1)
	for n := range s {
		nFloat := float64(n)
		if nFloat == 0 || nFloat == N-1 {
			sw[n] *= 0
		} else if nFloat < point1 {
			sw[n] *= 1 / (math.Exp(Za(nFloat) + 1))
		} else if point2 < nFloat {
			sw[n] *= 1 / (math.Exp(Zb(nFloat) + 1))
		}
	}
	return sw
}

// Kaiser window is a simple approximation of the DPSS window using Bessel functions, discovered by James Kaiser.
// https://en.wikipedia.org/wiki/Window_function#Kaiser_window
func Kaiser(s []float64, alpha float64) []float64 {
	sw := append([]float64{}, s...)

	piAlpha := math.Pi * alpha
	N := float64(len(s))
	for n := range s {
		sw[n] *= math.J0(piAlpha*math.Sqrt(1-math.Pow(2*float64(n)/(N-1)-1, 2))) / math.J0(piAlpha)
	}
	return sw
}

// Exponential window increases exponentially towards the center of the window and decreases exponentially in the second half.
// Since the exponential function never reaches zero,
// the values of the window at its limits are non-zero (it can be seen as the multiplication of an exponential function by a rectangular window).
// https://en.wikipedia.org/wiki/Window_function#Exponential_or_Poisson_window
func Exponential(s []float64, t float64) []float64 {
	sw := append([]float64{}, s...)

	N := len(s)
	for n := range s {
		sw[n] *= math.Exp(-math.Abs(float64(n)-float64(N-1)/2.0) / t)
	}
	return sw
}

// Poisson window is the same window function as Exponential.
func Poisson(s []float64, t float64) []float64 {
	return Exponential(s, t)
}
