package gowindow_test

import (
	"github.com/yut-kt/gowindow"
	"math/rand"
	"testing"
)

func makeBenchSlice() []float64 {
	const sliceSize = 1024
	s := make([]float64, sliceSize)
	for i := range s {
		s[i] = rand.Float64() * float64(rand.Intn(5))
	}
	return s
}

func BenchmarkRectangular(b *testing.B) {
	s := makeBenchSlice()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		gowindow.Rectangular(s)
	}
}

func BenchmarkTriangular(b *testing.B) {
	s := makeBenchSlice()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		gowindow.Triangular(s)
	}
}

func BenchmarkBartlett(b *testing.B) {
	s := makeBenchSlice()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		gowindow.Bartlett(s)
	}
}

func BenchmarkParzen(b *testing.B) {
	s := makeBenchSlice()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		gowindow.Parzen(s)
	}
}

func BenchmarkWelch(b *testing.B) {
	s := makeBenchSlice()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		gowindow.Welch(s)
	}
}

func BenchmarkSine(b *testing.B) {
	s := makeBenchSlice()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		gowindow.Sine(s)
	}
}

func BenchmarkPowerOfSine(b *testing.B) {
	s := makeBenchSlice()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		gowindow.PowerOfSine(s, 1)
	}
}

func BenchmarkPowerOfCosine(b *testing.B) {
	s := makeBenchSlice()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		gowindow.PowerOfCosine(s, 1)
	}
}

func BenchmarkHann(b *testing.B) {
	s := makeBenchSlice()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		gowindow.Hann(s)
	}
}

func BenchmarkHamming(b *testing.B) {
	s := makeBenchSlice()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		gowindow.Hamming(s)
	}
}

func BenchmarkBlackman(b *testing.B) {
	s := makeBenchSlice()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		gowindow.Blackman(s)
	}
}

func BenchmarkNuttall(b *testing.B) {
	s := makeBenchSlice()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		gowindow.Nuttall(s)
	}
}

func BenchmarkBlackmanNuttall(b *testing.B) {
	s := makeBenchSlice()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		gowindow.BlackmanNuttall(s)
	}
}

func BenchmarkBlackmanHarris(b *testing.B) {
	s := makeBenchSlice()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		gowindow.BlackmanHarris(s)
	}
}

func BenchmarkFlatTop(b *testing.B) {
	s := makeBenchSlice()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		gowindow.FlatTop(s)
	}
}

func BenchmarkRifeVincentClass1(b *testing.B) {
	s := makeBenchSlice()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = gowindow.RifeVincentClass1(s, 1)
	}
}

func BenchmarkRifeVincentClass3(b *testing.B) {
	s := makeBenchSlice()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = gowindow.RifeVincentClass3(s, 2)
	}
}

func BenchmarkGaussian(b *testing.B) {
	s := makeBenchSlice()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = gowindow.Gaussian(s, 0.4)
	}
}

func BenchmarkConfinedGaussian(b *testing.B) {
	s := makeBenchSlice()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		gowindow.ConfinedGaussian(s, 0.1)
	}
}

func BenchmarkApproximateConfinedGaussian(b *testing.B) {
	s := makeBenchSlice()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		gowindow.ApproximateConfinedGaussian(s, 0.1)
	}
}

func BenchmarkGeneralizedNormal(b *testing.B) {
	s := makeBenchSlice()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		gowindow.GeneralizedNormal(s, 0.4, 2)
	}
}

func BenchmarkTukey(b *testing.B) {
	s := makeBenchSlice()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		gowindow.Tukey(s, 0.5)
	}
}

func BenchmarkPlanckTaper(b *testing.B) {
	s := makeBenchSlice()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		gowindow.PlanckTaper(s, 0.1)
	}
}

func BenchmarkKaiser(b *testing.B) {
	s := makeBenchSlice()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		gowindow.Kaiser(s, 2)
	}
}

func BenchmarkExponential(b *testing.B) {
	s := makeBenchSlice()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		gowindow.Exponential(s, 512)
	}
}

func BenchmarkBartlettHann(b *testing.B) {
	s := makeBenchSlice()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		gowindow.BartlettHann(s)
	}
}

func BenchmarkHannPoisson(b *testing.B) {
	s := makeBenchSlice()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		gowindow.HannPoisson(s, 2)
	}
}

func BenchmarkLanczos(b *testing.B) {
	s := makeBenchSlice()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		gowindow.Lanczos(s)
	}
}
