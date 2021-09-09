package gowindow_test

import (
	"testing"

	"github.com/yut-kt/gowindow"
)

func BenchmarkRectangular(b *testing.B) {
	gowindow.Rectangular([]float64{1, 2, 3, 4})
}

func BenchmarkHamming(b *testing.B) {
	gowindow.Hanning([]float64{1, 2, 3, 4})
}

func BenchmarkHammingD(b *testing.B) {
	gowindow.HanningD([]float64{1, 2, 3, 4})
}

func BenchmarkHanning(b *testing.B) {
	gowindow.Hanning([]float64{1, 2, 3, 4})
}

func BenchmarkHann(b *testing.B) {
	gowindow.Hann([]float64{1, 2, 3, 4})
}

func BenchmarkHanningD(b *testing.B) {
	gowindow.HanningD([]float64{1, 2, 3, 4})
}

func BenchmarkHannD(b *testing.B) {
	gowindow.HannD([]float64{1, 2, 3, 4})
}

func BenchmarkBlackman(b *testing.B) {
	gowindow.Blackman([]float64{1, 2, 3, 4})
}

func BenchmarkBlackmanD(b *testing.B) {
	gowindow.BlackmanD([]float64{1, 2, 3, 4})
}
