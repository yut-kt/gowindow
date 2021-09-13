package gowindow

import "math"

// https://www.recordingblogs.com/wiki/bartlett-hann-window
func bartlettHann(s []float64) {
	N := float64(len(s))
	v := 2.0 * math.Pi / (N - 1)
	for k := range s {
		fk := float64(k)
		s[k] *= 0.62 - 0.48*math.Abs(fk/(N-1)-0.5) - 0.38*math.Cos(fk*v)
	}
}

func bartlettHannNew(s []float64) []float64 {
	N := float64(len(s))
	v := 2.0 * math.Pi / (N - 1)
	sw := make([]float64, len(s))
	for k := range s {
		fk := float64(k)
		sw[k] = s[k] * (0.62 - 0.48*math.Abs(fk/(N-1)-0.5) - 0.38*math.Cos(fk*v))
	}
	return sw
}
