package gowindow

import (
	"errors"
	"math"
)

// Group5 is Cosine-sum windows
// This family is also known as generalized cosine windows.
// https://en.wikipedia.org/wiki/Window_function#Cosine-sum_windows

// cosineSum from:
// Σ^K_k=0 (-1)^k * ak * cos(2πkn/N): The arg v is (2πn/N)
func cosineSum(a []float64, v float64) float64 {
	w := a[0]
	for k := range a[1:] {
		k++
		kFloat := float64(k)
		w += math.Pow(-1, kFloat) * a[k] * math.Cos(kFloat*v)
	}
	return w
}

// Hann window is customary cosine-sum windows for case K = 1 have the form:
// a1 = (1 - a0)
// https://en.wikipedia.org/wiki/Window_function#Hann_and_Hamming_windows
func Hann(s []float64) []float64 {
	sw := append([]float64{}, s...)
	a := []float64{0.5, 0.5}

	v := 2 * math.Pi / float64(len(s)-1)
	for n := range s {
		sw[n] *= cosineSum(a, float64(n)*v)
	}
	return sw
}

// Hanning window is the same window function as Hann.
func Hanning(s []float64) []float64 {
	return Hann(s)
}

// Hamming window is places a zero-crossing at frequency 5π/(N − 1),
// which cancels the first side-lobe of the Hann window,
// giving it a height of about one-fifth that of the Hann window.
// The Hamming window is often called the Hamming blip when used for pulse shaping.
// In the equiripple sense, the optimal values for the coefficients are a0 = 0.53836 and a1 = 0.46164.
// https://en.wikipedia.org/wiki/Window_function#Hann_and_Hamming_windows
func Hamming(s []float64) []float64 {
	sw := append([]float64{}, s...)
	a := []float64{0.53836, 0.46164}

	v := 2 * math.Pi / float64(len(s)-1)
	for n := range s {
		sw[n] *= cosineSum(a, float64(n)*v)
	}
	return sw
}

// HammingOriginal window is a Hamming window with different coefficients.
// Approximation of the coefficients to two decimal places substantially lowers the level of side-lobes, to a nearly equiripple condition.
// https://en.wikipedia.org/wiki/Window_function#Hann_and_Hamming_windows
func HammingOriginal(s []float64) []float64 {
	sw := append([]float64{}, s...)
	a := []float64{0.54, 0.46}

	v := 2 * math.Pi / float64(len(s)-1)
	for n := range s {
		sw[n] *= cosineSum(a, float64(n)*v)
	}
	return sw
}

// Blackman window is defined as:
// w[n] = a0 - a1*cos(2πn/N) + a2*cos(4πn/N)
// a0=(1-α)/2; a1=1/2; a2=α/2
// By common convention, the unqualified term Blackman window
// refers to Blackman's "not very serious proposal" of α = 0.16 (a0 = 0.42, a1 = 0.5, a2 = 0.08)
// https://en.wikipedia.org/wiki/Window_function#Blackman_window
func Blackman(s []float64) []float64 {
	sw := append([]float64{}, s...)
	a := []float64{0.42, 0.5, 0.08}

	v := 2 * math.Pi / float64(len(s)-1)
	for n := range s {
		sw[n] *= cosineSum(a, float64(n)*v)
	}
	return sw
}

// BlackmanExact window closely approximates the exact Blackman,
// with a0 = 7938/18608 ≈ 0.42659, a1 = 9240/18608 ≈ 0.49656, and a2 = 1430/18608 ≈ 0.076849.
// These exact values place zeros at the third and fourth side-lobes,
// but result in a discontinuity at the edges and a 6 dB/oct fall-off.
// The truncated coefficients do not null the side-lobes as well, but have an improved 18 dB/oct fall-off.
// https://en.wikipedia.org/wiki/Window_function#Blackman_window
func BlackmanExact(s []float64) []float64 {
	sw := append([]float64{}, s...)
	//a := []float64{0.42659, 0.49656, 0.076849}
	a := []float64{7938.0 / 18608, 9240.0 / 18608, 1430.0 / 18608}

	v := 2 * math.Pi / float64(len(s)-1)
	for n := range s {
		sw[n] *= cosineSum(a, float64(n)*v)
	}
	return sw
}

// Nuttall window is continuous form of Nuttall window, and its first derivative are continuous everywhere, like the Hann function.
// That is, the function goes to 0 at x = ±N/2, unlike the BlackmanNuttall, BlackmanHarris, and Hamming windows.
// The Blackman window (α = 0.16) is also continuous with continuous derivative at the edge, but the BlackmanExact window is not.
// https://en.wikipedia.org/wiki/Window_function#Nuttall_window,_continuous_first_derivative
func Nuttall(s []float64) []float64 {
	sw := append([]float64{}, s...)
	a := []float64{0.355768, 0.487396, 0.144232, 0.012604}

	v := 2 * math.Pi / float64(len(s)-1)
	for n := range s {
		sw[n] *= cosineSum(a, float64(n)*v)
	}
	return sw
}

// BlackmanNuttall window is defined as:
// a0 - a1*cos(2πn/N) + a2*cos(4πn/N) - a3*cos(6πn/N)
// a0=0.3635819; a1=0.4891775; a2=0.1365995; a3=0.0106411;
// https://en.wikipedia.org/wiki/Window_function#Blackman%E2%80%93Nuttall_window
func BlackmanNuttall(s []float64) []float64 {
	sw := append([]float64{}, s...)
	a := []float64{0.3635819, 0.4891775, 0.1365995, 0.0106411}

	v := 2 * math.Pi / float64(len(s)-1)
	for n := range s {
		sw[n] *= cosineSum(a, float64(n)*v)
	}
	return sw
}

// BlackmanHarris is generalization of the Hamming family,
// produced by adding more shifted sinc functions, meant to minimize side-lobe levels.
// https://en.wikipedia.org/wiki/Window_function#Blackman%E2%80%93Harris_window
func BlackmanHarris(s []float64) []float64 {
	sw := append([]float64{}, s...)
	a := []float64{0.35875, 0.48829, 0.14128, 0.01168}

	v := 2 * math.Pi / float64(len(s)-1)
	for n := range s {
		sw[n] *= cosineSum(a, float64(n)*v)
	}
	return sw
}

// FlatTop window is a partially negative-valued window that has minimal scalloping loss in the frequency domain.
// That property is desirable for the measurement of amplitudes of sinusoidal frequency components.
// Drawbacks of the broad bandwidth are poor frequency resolution and high § Noise bandwidth.
// Flat top windows can be designed using low-pass filter design methods,
// or they may be of the usual cosine-sum variety.
// https://en.wikipedia.org/wiki/Window_function#Flat_top_window
func FlatTop(s []float64) []float64 {
	sw := append([]float64{}, s...)
	a := []float64{0.21557895, 0.41663158, 0.277263158, 0.083578947, 0.006947368}

	v := 2 * math.Pi / float64(len(s)-1)
	for n := range s {
		sw[n] *= cosineSum(a, float64(n)*v)
	}
	return sw
}

// From here on, Rife-Vincent windows.
// Rife–Vincent windows are customarily scaled for unity average value, instead of unity peak value.

// RifeVincentClass1 is defined by minimizing the high-Order side-lobe amplitude.
// Order(K) is 1~4.
func RifeVincentClass1(s []float64, order int) ([]float64, error) {
	if order < 1 || 4 < order {
		return nil, errors.New("order is 1 ~ 4")
	}
	sw := append([]float64{}, s...)
	a := [][]float64{
		{1, -1},
		{1, 4.0 / 3, 1.0 / 3},
		{1, 1.5, 0.6, 0.1},
		{1, 1.6, 0.8, 8.0 / 35.0, 1.0 / 35.0},
	}[order-1]

	v := 2 * math.Pi / float64(len(s)-1)
	for n := range s {
		sw[n] *= cosineSum(a, float64(n)*v)
	}
	return sw, nil
}

// RifeVincentClass2 minimizes the main-lobe width for a given maximum side-lobe.
// ※Deprecated due to lack of confirmation.
// decibel[dB] is any of 36, 42, 48, 54, 60, 66.
func RifeVincentClass2(s []float64, decibel int) ([]float64, error) {
	sw := append([]float64{}, s...)
	m := map[float64]float64{36: 3, 42: 4, 48: 5, 54: 6, 60: 7, 66: 9}
	P := float64(decibel)
	R, ok := m[P]
	if !ok {
		return nil, errors.New("decibel[dB] is any of 36, 42, 48, 54, 60, 66")
	}
	a := func(P, R float64) []float64 {
		coefficients := make([]float64, int(P)+1)
		lambda := 1.0 / math.Pi * math.Log(R+math.Sqrt(R*R-1))
		variance := math.Pow(P+1, 2) / (math.Pow(lambda, 2) + math.Pow(P+0.5, 2))
		for l := 0.0; l <= P; l++ {
			numerator := 1.0
			for i := 1.0; i <= P; i++ {
				numerator *= 1 - math.Pow(l/variance, 2)/(math.Pow(lambda, 2)+math.Pow(i-0.5, 2))
			}

			denominator := 1.0
			for i := 1.0; i <= P; i++ {
				if i == l {
					continue
				}
				numerator *= 1 - math.Pow(l, 2)/math.Pow(i, 2)
			}
			coefficients[int(l)] = -numerator / denominator
		}
		return coefficients
	}(P, R)

	v := 2 * math.Pi / float64(len(s)-1)
	for n := range s {
		sw[n] *= cosineSum(a, float64(n)*v)
	}
	return sw, nil
}

// RifeVincentClass3 is a compromise for which Order K = 2 resembles the § Blackman window.
// Order(K) is 2~4.
func RifeVincentClass3(s []float64, order int) ([]float64, error) {
	if order < 2 || 4 < order {
		return nil, errors.New("order is 2 ~ 4")
	}
	sw := append([]float64{}, s...)
	a := [][]float64{
		{1, 1.19685, 0.19685},
		{1, 1.43596, 0.497537, 0.061576},
		{1, 1.56627, 0.725448, 0.180645, 0.017921},
	}[order-2]

	v := 2 * math.Pi / float64(len(s)-1)
	for n := range s {
		sw[n] *= cosineSum(a, float64(n)*v)
	}
	return sw, nil
}

// Akaike window made by akaike
// https://ja.wikipedia.org/wiki/%E8%B5%A4%E6%B1%A0%E5%BC%98%E6%AC%A1
func Akaike(s []float64) []float64 {
	sw := append([]float64{}, s...)

	a := []float64{0.625, 0.5, 0.125}

	v := 2 * math.Pi / float64(len(s)-1)
	for n := range s {
		sw[n] *= cosineSum(a, float64(n)*v)
	}
	return sw
}