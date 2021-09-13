package gowindow

import "math"

func powerOfCosine(s []float64, alpha float64) {
	v := math.Pi / float64(len(s)-1)
	halfPi := math.Pi / 2.0
	for n := range s {
		s[n] *= math.Pow(math.Cos(float64(n)*v-halfPi), alpha)
	}
}

func powerOfCosineNew(s []float64, alpha float64) []float64 {
	v := math.Pi / float64(len(s)-1)
	halfPi := math.Pi / 2.0
	sw := make([]float64, len(s))
	for n := range s {
		sw[n] = s[n] * math.Pow(math.Cos(float64(n)*v-halfPi), alpha)
	}
	return sw
}
