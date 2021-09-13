package gowindow

import "math"

func ultraspherical(s []float64, mu, xZero float64) {
	N := len(s)
	for n := range s {
		fn := float64(n)
		sum := 0.0
		for k := 1; k <= N/2; k++ {
			fk := float64(k)
			sum += c(mu, N, xZero*math.Cos((fk*math.Pi)/float64(N+1))) * math.Cos(2.0*fn*math.Pi*fk/float64(N+1))
		}
		s[n] *= 1.0 / float64(N+1) * (c(mu, N, xZero) + sum)
	}
}

func ultrasphericalNew(s []float64, mu, xZero float64) []float64 {
	N := len(s)
	sw := make([]float64, len(s))
	for n := range s {
		fn := float64(n)
		sum := 0.0
		for k := 1; k <= N/2; k++ {
			fk := float64(k)
			sum += c(mu, N, xZero*math.Cos((fk*math.Pi)/float64(N+1))) * math.Cos(2.0*fn*math.Pi*fk/float64(N+1))
		}
		sw[n] = s[n] * (1.0 / float64(N+1) * (c(mu, N, xZero) + sum))
	}
	return sw
}

func c(alpha float64, n int, x float64) float64 {
	switch n {
	case 0:
		return 1.0
	case 1:
		return 2.0 * alpha * x
	}

	fn := float64(n)
	return 1.0 / fn * (2.0*x*(fn+alpha-1.0)*c(alpha, n-1, x) - (fn+2.0*alpha-2.0)*c(alpha, n-2, x))
}
