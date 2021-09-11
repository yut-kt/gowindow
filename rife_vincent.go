package gowindow

import (
	"math"
)

type class int

const (
	class1 class = iota
	class2
	class3
)

type order int

const (
	nonOrder order = iota
	order1
	order2
	order3
	order4
)

type decibel int

const (
	nonDecibel decibel = iota
	dB36
	dB42
	dB48
	dB54
	dB60
	dB66
)

func rifeVincent(s []float64, c class, o order, d decibel) {
	v := 2.0 * math.Pi / float64(len(s)-1)
	coefficients := getCoefficients(c, o, d)
	for n := range s {
		val := v * float64(n)
		s[n] *= cosineSum(coefficients, val)
	}
}

func rifeVincentNew(s []float64, c class, o order, d decibel) []float64 {
	sw := make([]float64, len(s))
	v := 2.0 * math.Pi / float64(len(s)-1)
	coefficients := getCoefficients(c, o, d)
	for n := range s {
		val := v * float64(n)
		sw[n] = s[n] * cosineSum(coefficients, val)
	}
	return sw
}

func getCoefficients(c class, o order, d decibel) []float64 {
	switch c {
	case class1:
		switch o {
		case order1:
			return []float64{1, -1}
		case order2:
			return []float64{1, 4.0 / 3.0, 1.0 / 3.0}
		case order3:
			return []float64{1, 3.0 / 2.0, 3.0 / 5.0, 1.0 / 10.0}
		case order4:
			return []float64{1, 8.0 / 5.0, 4.0 / 5.0, 8.0 / 35.0, 1.0 / 35.0}
		}
	case class2:
		switch d {
		case dB36:
			return getClass2Coefficients(36, 3)
		case dB42:
			return getClass2Coefficients(42, 4)
		case dB48:
			return getClass2Coefficients(48, 5)
		case dB54:
			return getClass2Coefficients(54, 6)
		case dB60:
			return getClass2Coefficients(60, 7)
		case dB66:
			return getClass2Coefficients(66, 9)
		}
	case class3:
		switch o {
		case order2:
			return []float64{1, 1.19685, 0.19685}
		case order3:
			return []float64{1, 1.43596, 0.497537, 0.061576}
		case order4:
			return []float64{1, 1.56627, 0.725448, 0.180645, 0.017921}
		}
	}
	return []float64{}
}

func getClass2Coefficients(P, R float64) []float64 {
	coefficients := make([]float64, int(P)+1)
	lambda := 1.0 / math.Pi * math.Log(R+math.Sqrt(R*R-1))
	variance := math.Pow(P+1, 2) / (math.Pow(lambda, 2) + math.Pow(P+0.5, 2))
	for l := 0.0; l <= P; l++ {
		numerator := 1.0
		for i := 1.0; i <= P; i++ {
			numerator *= 1 - math.Pow(l/variance, 2)/(math.Pow(lambda, 2)+math.Pow(i-0.5, 2))
		}

		denominator := 1.0
		for i := 1.0; i <= P; i++ {
			if i == l {
				continue
			}
			numerator *= 1 - math.Pow(l, 2)/math.Pow(i, 2)
		}
		coefficients[int(l)] = -numerator / denominator
	}
	return coefficients
}
