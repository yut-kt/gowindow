package gowindow

import "math"

func lanczos(s []float64) {
	N := float64(len(s))
	halfN := (N - 1) / 2
	for k := range s {
		fk := float64(k)
		if fk != halfN {
			numerator := math.Sin(math.Pi * (2*fk/(N-1) - 1))
			denominator := math.Pi * (2*fk/(N-1) - 1)
			s[k] *= numerator / denominator
		}
	}
}

func lanczosNew(s []float64) []float64 {
	N := float64(len(s))
	halfN := (N - 1) / 2
	sw := make([]float64, len(s))
	for k := range s {
		fk := float64(k)
		if fk != halfN {
			numerator := math.Sin(math.Pi * (2*fk/(N-1) - 1))
			denominator := math.Pi * (2*fk/(N-1) - 1)
			sw[k] = s[k] * (numerator / denominator)
		} else {
			sw[k] = s[k]
		}
	}
	return sw
}
