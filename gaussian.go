package gowindow

import (
	"math"
)

func gaussian(s []float64, sd float64) {
	halfN := float64(len(s)-1) / 2.0
	for n := range s {
		s[n] *= math.Exp(-0.5 * math.Pow((float64(n)-halfN)/(sd*halfN), 2))
	}
}

func gaussianNew(s []float64, sd float64) []float64 {
	sw := make([]float64, len(s))
	halfN := float64(len(s)-1) / 2.0
	for n := range s {
		sw[n] = s[n] * math.Exp(-0.5*math.Pow((float64(n)-halfN)/(sd*halfN), 2))
	}
	return sw
}

func gauss(s []float64, sd float64) {
	gaussian(s, sd)
}

func gaussNew(s []float64, sd float64) []float64 {
	return gaussianNew(s, sd)
}
