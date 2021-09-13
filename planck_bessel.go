package gowindow

import "math"

func planckBessel(s []float64, epsilon, alpha float64) {
	tmp := make([]float64, len(s))
	for i := 0; i < len(tmp); i++ {
		tmp[i] = 1
	}
	kaiserWindow := kaiserNew(tmp, alpha)

	N := float64(len(s))
	Za := func(n float64) float64 {
		return 2 * epsilon * (1/(1+2*n/(N-1)) + 1/(1-2*epsilon+2*n/(N-1)))
	}
	Zb := func(n float64) float64 {
		return 2 * epsilon * (1/(1-2*n/(N-1)) + 1/(1-2*epsilon-2*n/(N-1)))
	}
	for n := range s {
		fn := float64(n)
		var weight float64
		if fn < epsilon*(N-1) {
			weight = 1 / (math.Exp(Za(fn)) + 1)
		} else if epsilon*(N-1) < fn && fn < (1-epsilon)*(N-1) {
			weight = 1
		} else if (1-epsilon)*(N-1) < fn {
			weight = 1 / (math.Exp(Zb(fn)) + 1)
		} else {
			weight = 0
		}
		s[n] *= weight * kaiserWindow[n]
	}
}

func planckBesselNew(s []float64, epsilon, alpha float64) []float64 {
	tmp := make([]float64, len(s))
	for i := 0; i < len(tmp); i++ {
		tmp[i] = 1
	}
	kaiserWindow := kaiserNew(tmp, alpha)

	N := float64(len(s))
	Za := func(n float64) float64 {
		return 2 * epsilon * (1/(1+2*n/(N-1)) + 1/(1-2*epsilon+2*n/(N-1)))
	}
	Zb := func(n float64) float64 {
		return 2 * epsilon * (1/(1-2*n/(N-1)) + 1/(1-2*epsilon-2*n/(N-1)))
	}

	sw := make([]float64, len(s))
	copy(sw, s)
	for n := range s {
		fn := float64(n)
		var weight float64
		if fn < epsilon*(N-1) {
			weight = 1 / (math.Exp(Za(fn)) + 1)
		} else if epsilon*(N-1) < fn && fn < (1-epsilon)*(N-1) {
			weight = 1
		} else if (1-epsilon)*(N-1) < fn {
			weight = 1 / (math.Exp(Zb(fn)) + 1)
		} else {
			weight = 0
		}
		sw[n] *= weight * kaiserWindow[n]
	}
	return sw
}
