package gowindow

func confinedGaussian(s []float64, SDt float64) {
	L := float64(len(s) + 1)
	g := gaussianFunction{SDt: SDt, L: L, HalfN: float64(len(s)-1) / 2.0}
	for n := range s {
		fn := float64(n)
		numerator := g.G(fn+L) + g.G(fn-L)
		denominator := g.G(-0.5+L) + g.G(-0.5-L)
		s[n] *= g.G(fn) - g.G(-0.5)*(numerator/denominator)
	}
}

func confinedGaussianNew(s []float64, SDt float64) []float64 {
	L := float64(len(s) + 1)
	g := gaussianFunction{SDt: SDt, L: L, HalfN: float64(len(s)-1) / 2.0}
	sw := make([]float64, len(s))
	for n := range s {
		fn := float64(n)
		numerator := g.G(fn+L) + g.G(fn-L)
		denominator := g.G(-0.5+L) + g.G(-0.5-L)
		sw[n] = s[n] * (g.G(fn) - g.G(-0.5)*(numerator/denominator))
	}
	return sw
}
