package gowindow

import (
	"math"
)

func kaiser(s []float64, alpha float64) {
	piAlpha := math.Pi * alpha
	N := float64(len(s))
	for k := range s {
		fk := float64(k)
		s[k] *= math.J0(piAlpha*math.Sqrt(1-math.Pow(2*fk/(N-1)-1, 2))) / math.J0(piAlpha)
	}
}

func kaiserNew(s []float64, alpha float64) []float64 {
	piAlpha := math.Pi * alpha
	N := float64(len(s))
	sw := make([]float64, len(s))
	for k := range s {
		fk := float64(k)
		sw[k] = s[k] * (math.J0(piAlpha*math.Sqrt(1-math.Pow(2*fk/(N-1)-1, 2))) / math.J0(piAlpha))
	}
	return sw
}
