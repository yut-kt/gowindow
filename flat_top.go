package gowindow

import "math"

func flatTop(s []float64) {
	v := 2.0 * math.Pi / float64(len(s)-1)
	for n := range s {
		val := v * float64(n)
		s[n] *= 0.21557895 - 0.41663158*math.Cos(val) + 0.277263158*math.Cos(2*val) - 0.083578947*math.Cos(3*val) + 0.006947368*math.Cos(4*val)
	}
}

func flatTopNew(s []float64) []float64 {
	sw := make([]float64, len(s))
	v := 2.0 * math.Pi / float64(len(s)-1)
	for n := range s {
		val := v * float64(n)
		sw[n] = s[n] * (0.21557895 - 0.41663158*math.Cos(val) + 0.277263158*math.Cos(2*val) - 0.083578947*math.Cos(3*val) + 0.006947368*math.Cos(4*val))
	}
	return sw
}
