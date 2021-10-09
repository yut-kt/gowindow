package gowindow

func sumSlice(s []float64) float64 {
	var sum float64
	for i := range s {
		sum += s[i]
	}
	return sum
}

func makeFloatSlice(size int, initV float64) []float64 {
	s := make([]float64, size)
	for i := range s {
		s[i] = initV
	}
	return s
}

// KaiserBesselDerived is failure work
// implementation is incorrect because the window value is incorrect
//func KaiserBesselDerived(s []float64, alpha float64) []float64 {
//	sw := append([]float64{}, s...)
//
//	N := len(s)
//	halfN := (N - 1) / 2
//
//	weightSlice := makeFloatSlice(len(s), 1)
//	denominator := sumSlice(Kaiser(weightSlice[:halfN+1], alpha))
//	for n := range s {
//		var numerator float64
//		if n < halfN {
//			numerator = sumSlice(Kaiser(weightSlice[:n], alpha))
//		} else {
//			numerator = sumSlice(Kaiser(weightSlice[:N-n], alpha))
//		}
//		sw[n] *= math.Sqrt(numerator/denominator)
//	}
//	return sw
//}
