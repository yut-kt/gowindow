package gowindow

import "math"

func cosineSum(s []float64, v float64) float64 {
	weight := s[0]
	for n := range s[1:] {
		n++
		w := s[n] * math.Cos(float64(n)*v)
		if n%2 == 1 {
			w *= -1
		}
		weight += w
	}
	return weight
}
