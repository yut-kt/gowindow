package gowindow

import "math"

// https://www.recordingblogs.com/wiki/hann-poisson-window
func hannPoisson(s []float64, alpha float64) {
	M := float64(len(s)-1) / 2
	for k := range s {
		s[k] *= 0.5 * (1 - math.Cos(math.Pi*float64(k)/M)) * math.Exp(-alpha*math.Abs(float64(k)-M)/M)
	}
}

func hannPoissonNew(s []float64, alpha float64) []float64 {
	M := float64(len(s)-1) / 2
	sw := make([]float64, len(s))
	for k := range s {
		sw[k] = s[k] * (0.5 * (1 - math.Cos(math.Pi*float64(k)/M)) * math.Exp(-alpha*math.Abs(float64(k)-M)/M))
	}
	return sw
}
