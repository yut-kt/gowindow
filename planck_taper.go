package gowindow

import "math"

// https://www.recordingblogs.com/wiki/planck-taper-window
func planckTaper(s []float64, epsilon float64) {
	N := float64(len(s))
	Za := func(k float64) float64 {
		return epsilon * (N - 1) * (1.0/k + 1/(k-epsilon*(N-1)))
	}
	Zb := func(k float64) float64 {
		return epsilon * (N - 1) * (1/(N-1-k) + 1/((1-epsilon)*(N-1)-k))

	}
	point1 := epsilon * (N - 1)
	point2 := (1 - epsilon) * (N - 1)
	for k := range s {
		fk := float64(k)
		if fk == 0 || fk == N-1 {
			s[k] *= 0
		} else if fk < point1 {
			s[k] *= 1 / (math.Exp(Za(fk) + 1))
		} else if point2 < fk {
			s[k] *= 1 / (math.Exp(Zb(fk) + 1))
		}
	}
}

func planckTaperNew(s []float64, epsilon float64) []float64 {
	N := float64(len(s))
	Za := func(k float64) float64 {
		return epsilon * (N - 1) * (1.0/k + 1/(k-epsilon*(N-1)))
	}
	Zb := func(k float64) float64 {
		return epsilon * (N - 1) * (1/(N-1-k) + 1/((1-epsilon)*(N-1)-k))

	}
	point1 := epsilon * (N - 1)
	point2 := (1 - epsilon) * (N - 1)
	sw := make([]float64, len(s))
	copy(sw, s)
	for k := range s {
		fk := float64(k)
		if fk == 0 || fk == N-1 {
			sw[k] = 0
		} else if fk < point1 {
			sw[k] = s[k] * (1 / (math.Exp(Za(fk) + 1)))
		} else if point2 < fk {
			sw[k] = s[k] * (1 / (math.Exp(Zb(fk) + 1)))
		}
	}
	return sw
}
