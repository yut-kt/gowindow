package gowindow

import (
	"math"
)

// Triangular is function to apply Triangular Window
// Non-Destructive function
func Triangular(s []float64) []float64 {
	sw := make([]float64, len(s))
	for k := range s {
		x := float64(k) / float64(len(s)-1)
		sw[k] = s[k] * (1 - 2*math.Abs(x-0.5))
	}
	return sw
}

// Bartlett is function to apply Bartlett Window
// Non-Destructive function
func Bartlett(s []float64) []float64 {
	return Triangular(s)
}

// Fejer is function to apply Fejer Window
// Non-Destructive function
func Fejer(s []float64) []float64 {
	return Triangular(s)
}

// TriangularD is function to apply Triangular Window
// Destructive function
func TriangularD(s []float64) {
	for k := range s {
		x := float64(k) / float64(len(s)-1)
		s[k] *= 1 - 2*math.Abs(x-0.5)
	}
}

// BartlettD is function to apply Bartlett Window
// Destructive function
func BartlettD(s []float64) {
	TriangularD(s)
}

// FejerD is function to apply Fejer Window
// Destructive function
func FejerD(s []float64) {
	TriangularD(s)
}
