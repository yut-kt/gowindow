package gowindow

import "math"

func flatTop(s []float64) {
	w := []float64{0.21557895, 0.4166315, 0.277263158, 0.083578947, 0.006947368}
	v := 2.0 * math.Pi / float64(len(s)-1)
	for n := range s {
		val := v * float64(n)
		s[n] *= cosineSum(w, val)
	}
}

func flatTopNew(s []float64) []float64 {
	w := []float64{0.21557895, 0.4166315, 0.277263158, 0.083578947, 0.006947368}
	sw := make([]float64, len(s))
	v := 2.0 * math.Pi / float64(len(s)-1)
	for n := range s {
		val := v * float64(n)
		sw[n] = s[n] * cosineSum(w, val)
	}
	return sw
}
