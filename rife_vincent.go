package gowindow

import (
	"math"
)

type class int

const (
	// Class1 is defined by minimizing the high-Order side-lobe amplitude
	Class1 class = iota
	// Class2 minimizes the main-lobe width for a given maximum side-lobe
	Class2
	// Class3 is a compromise for which Order K = 2 resembles the § Blackman window
	Class3
)

// order use only for Class1, Class3
type order int

const (
	// NonOrder is only Class2
	NonOrder order = iota
	// Order1 is (K = 1)
	Order1
	// Order2 is (K = 2)
	Order2
	// Order3 is (K = 3)
	Order3
	// Order4 is (K = 4)
	Order4
)

// decibel use only for Class2
type decibel int

const (
	// NonDecibel is only Class1, Class3
	NonDecibel decibel = iota
	Decibel36
	Decibel42
	Decibel48
	Decibel54
	Decibel60
	Decibel66
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
	case Class1:
		switch o {
		case Order1:
			return []float64{1, -1}
		case Order2:
			return []float64{1, 4.0 / 3.0, 1.0 / 3.0}
		case Order3:
			return []float64{1, 3.0 / 2.0, 3.0 / 5.0, 1.0 / 10.0}
		case Order4:
			return []float64{1, 8.0 / 5.0, 4.0 / 5.0, 8.0 / 35.0, 1.0 / 35.0}
		}
	case Class2:
		switch d {
		case Decibel36:
			return getClass2Coefficients(36, 3)
		case Decibel42:
			return getClass2Coefficients(42, 4)
		case Decibel48:
			return getClass2Coefficients(48, 5)
		case Decibel54:
			return getClass2Coefficients(54, 6)
		case Decibel60:
			return getClass2Coefficients(60, 7)
		case Decibel66:
			return getClass2Coefficients(66, 9)
		}
	case Class3:
		switch o {
		case Order2:
			return []float64{1, 1.19685, 0.19685}
		case Order3:
			return []float64{1, 1.43596, 0.497537, 0.061576}
		case Order4:
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
