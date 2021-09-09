package gowindow

import (
	"math"
)

func parzen(s []float64) {
	N := float64(len(s))
	for n := range s {
		s[n] *= parzenCalculation(float64(n)-N/2.0, N)
	}
}

func parzenNew(s []float64) []float64 {
	sw := make([]float64, len(s))
	N := float64(len(s))
	for n := range s {
		sw[n] = s[n] * parzenCalculation(float64(n)-N/2.0, N)
	}
	return sw
}

func parzenCalculation(n, N float64) float64 {
	nAbs := math.Abs(n)
	LQuarter, LHalf := float64(N+1)/4, float64(N+1)/2
	if nAbs <= LQuarter {
		return 1.0 - 6.0*math.Pow(n/LHalf, 2)*(1.0-nAbs/LHalf)
	}
	return 2.0 * math.Pow(1.0-nAbs/LHalf, 3)
}
