package gowindow

// reference formula
// https://ja.wikipedia.org/wiki/%E7%AA%93%E9%96%A2%E6%95%B0#%E3%82%A6%E3%82%A7%E3%83%AB%E3%83%81%E7%AA%93
func welch(s []float64) {
	decN := float64(len(s) - 1)
	for k := range s {
		x := float64(k) / decN
		s[k] *= 4*x - 4*x*x
	}
}

func welchNew(s []float64) []float64 {
	decN := float64(len(s) - 1)
	sw := make([]float64, len(s))
	for k := range s {
		x := float64(k) / decN
		sw[k] = s[k] * (4*x - 4*x*x)
	}
	return sw
}
