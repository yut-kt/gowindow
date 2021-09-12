package gowindow

import "math"

// https://www.recordingblogs.com/wiki/dolph-chebychev-window
func dolphChebyshev(s []float64, omegaZero float64) {
	M := float64(len(s)-1) / 2.0
	T2M := func(x float64) float64 {
		if math.Abs(x) <= 1 {
			return math.Cos(2.0 * M * math.Acos(x))
		}
		return math.Cosh(2.0 * M * math.Acosh(x))
	}

	N := float64(len(s))
	omegaNs := make([]float64, int(M))
	for n := range omegaNs {
		if n == 0 {
			omegaNs[n] = omegaZero
		} else {
			omegaNs[n] = float64(n) * math.Pi / (N - 1)
		}
	}

	for k := range s {
		sum := 0.0
		denominator := T2M(1.0 / math.Cos(omegaNs[0]/2.0))
		for n := 0; n < int(M); n++ {
			numerator := T2M(math.Cos(omegaNs[n]/2.0) / math.Cos(omegaNs[0]/2.0))
			sum += numerator / denominator * math.Cos((float64(k)-M)*omegaNs[n])
		}
		s[k] *= sum / N
	}
}

func dolphChebyshevNew(s []float64, omegaZero float64) []float64 {
	M := float64(len(s)-1) / 2.0
	T2M := func(x float64) float64 {
		if math.Abs(x) <= 1 {
			return math.Cos(2.0 * M * math.Acos(x))
		}
		return math.Cosh(2.0 * M * math.Acosh(x))
	}

	N := float64(len(s))
	omegaNs := make([]float64, int(M))
	for n := range omegaNs {
		if n == 0 {
			omegaNs[n] = omegaZero
		} else {
			omegaNs[n] = float64(n) * math.Pi / (N - 1)
		}
	}

	sw := make([]float64, len(s))
	for k := range s {
		sum := 0.0
		denominator := T2M(1.0 / math.Cos(omegaNs[0]/2.0))
		for n := 0; n < int(M); n++ {
			numerator := T2M(math.Cos(omegaNs[n]/2.0) / math.Cos(omegaNs[0]/2.0))
			sum += numerator / denominator * math.Cos((float64(k)-M)*omegaNs[n])
		}
		sw[k] = s[k] * (sum / N)
	}
	return sw
}
