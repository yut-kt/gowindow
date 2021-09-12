package gowindow

import (
	"math"
)

// https://www.recordingblogs.com/wiki/tukey-window
func tukey(s []float64, alpha float64) {
	M := float64(len(s)-1) / 2.0
	alphaM := alpha * M
	for k := range s {
		abs := math.Abs(float64(k) - M)
		if abs >= alphaM {
			s[k] *= 0.5 + 0.5*math.Cos(math.Pi*(abs-alphaM)/(M-alphaM))
		}
	}
}

func tukeyNew(s []float64, alpha float64) []float64 {
	M := float64(len(s)-1) / 2.0
	alphaM := alpha * M
	sw := make([]float64, len(s))
	copy(sw, s)
	for k := range s {
		abs := math.Abs(float64(k) - M)
		if abs >= alphaM {
			sw[k] = s[k]*0.5 + 0.5*math.Cos(math.Pi*(abs-alphaM)/(M-alphaM))
		}
	}
	return sw
}
