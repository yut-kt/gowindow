package gowindow

import (
	"math"
)

func generalizedNormal(s []float64, sd float64, p float64) {
	halfN := float64(len(s)-1) / 2.0
	for n := range s {
		s[n] *= math.Exp(-math.Pow((float64(n)-halfN)/(sd*halfN), p))
	}
}

func generalizedNormalNew(s []float64, sd float64, p float64) []float64 {
	sw := make([]float64, len(s))
	halfN := float64(len(s)-1) / 2.0
	for n := range s {
		sw[n] = s[n] * math.Exp(-math.Pow((float64(n)-halfN)/(sd*halfN), p))
	}
	return sw
}
