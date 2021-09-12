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

type gaussianFunction struct {
	SDt   float64
	L     float64
	HalfN float64
}

func (g *gaussianFunction) G(x float64) float64 {
	return math.Exp(-math.Pow((x-g.HalfN)/2.0*g.L*g.SDt, 2))
}
