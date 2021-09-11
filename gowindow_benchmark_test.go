package gowindow_test

import (
	"math/rand"
	"testing"

	"github.com/yut-kt/gowindow"
)

func makeNSlice() []float64 {
	const sliceSize = 1024
	s := make([]float64, sliceSize)
	for i := range s {
		s[i] = rand.Float64() * float64(rand.Intn(5))
	}
	return s
}

func BenchmarkApply_rectangular(b *testing.B) {
	s := makeNSlice()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		gowindow.Apply(s, gowindow.Rectangular)
	}
}

func BenchmarkApplyNew_rectangular(b *testing.B) {
	s := makeNSlice()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		gowindow.ApplyNew(s, gowindow.Rectangular)
	}
}

func BenchmarkApply_triangular(b *testing.B) {
	s := makeNSlice()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		gowindow.Apply(s, gowindow.Triangular)
	}
}

func BenchmarkApplyNew_triangular(b *testing.B) {
	s := makeNSlice()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		gowindow.ApplyNew(s, gowindow.Triangular)
	}
}

func BenchmarkApply_bartlett(b *testing.B) {
	s := makeNSlice()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		gowindow.Apply(s, gowindow.Bartlett)
	}
}

func BenchmarkApplyNew_bartlett(b *testing.B) {
	s := makeNSlice()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		gowindow.ApplyNew(s, gowindow.Bartlett)
	}
}

func BenchmarkApply_fejer(b *testing.B) {
	s := makeNSlice()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		gowindow.Apply(s, gowindow.Fejer)
	}
}

func BenchmarkApplyNew_fejer(b *testing.B) {
	s := makeNSlice()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		gowindow.ApplyNew(s, gowindow.Fejer)
	}
}

func BenchmarkApply_parzen(b *testing.B) {
	s := makeNSlice()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		gowindow.Apply(s, gowindow.Parzen)
	}
}

func BenchmarkApplyNew_parzen(b *testing.B) {
	s := makeNSlice()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		gowindow.ApplyNew(s, gowindow.Parzen)
	}
}

func BenchmarkApply_deLaValleePoussin(b *testing.B) {
	s := makeNSlice()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		gowindow.Apply(s, gowindow.DeLaValleePoussin)
	}
}

func BenchmarkApplyNew_deLaValleePoussin(b *testing.B) {
	s := makeNSlice()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		gowindow.ApplyNew(s, gowindow.DeLaValleePoussin)
	}
}

func BenchmarkApply_welch(b *testing.B) {
	s := makeNSlice()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		gowindow.Apply(s, gowindow.Welch)
	}
}

func BenchmarkApplyNew_welch(b *testing.B) {
	s := makeNSlice()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		gowindow.ApplyNew(s, gowindow.Welch)
	}
}

func BenchmarkApply_sine(b *testing.B) {
	s := makeNSlice()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		gowindow.Apply(s, gowindow.Sine)
	}
}

func BenchmarkApplyNew_sine(b *testing.B) {
	s := makeNSlice()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		gowindow.ApplyNew(s, gowindow.Sine)
	}
}

func BenchmarkApply_hanning(b *testing.B) {
	s := makeNSlice()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		gowindow.Apply(s, gowindow.Hanning)
	}
}

func BenchmarkApplyNew_hanning(b *testing.B) {
	s := makeNSlice()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		gowindow.ApplyNew(s, gowindow.Hanning)
	}
}

func BenchmarkApply_hann(b *testing.B) {
	s := makeNSlice()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		gowindow.Apply(s, gowindow.Hann)
	}
}

func BenchmarkApplyNew_hann(b *testing.B) {
	s := makeNSlice()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		gowindow.ApplyNew(s, gowindow.Hann)
	}
}

func BenchmarkApply_hamming(b *testing.B) {
	s := makeNSlice()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		gowindow.Apply(s, gowindow.Hamming)
	}
}

func BenchmarkApplyNew_hamming(b *testing.B) {
	s := makeNSlice()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		gowindow.ApplyNew(s, gowindow.Hamming)
	}
}

func BenchmarkApply_blackman(b *testing.B) {
	s := makeNSlice()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		gowindow.Apply(s, gowindow.Blackman)
	}
}

func BenchmarkApplyNew_blackman(b *testing.B) {
	s := makeNSlice()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		gowindow.ApplyNew(s, gowindow.Blackman)
	}
}

func BenchmarkApply_nuttall(b *testing.B) {
	s := makeNSlice()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		gowindow.Apply(s, gowindow.Nuttall)
	}
}

func BenchmarkApplyNew_nuttall(b *testing.B) {
	s := makeNSlice()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		gowindow.ApplyNew(s, gowindow.Nuttall)
	}
}

func BenchmarkApply_blackmanNuttall(b *testing.B) {
	s := makeNSlice()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		gowindow.Apply(s, gowindow.BlackmanNuttall)
	}
}

func BenchmarkApplyNew_blackmanNuttall(b *testing.B) {
	s := makeNSlice()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		gowindow.ApplyNew(s, gowindow.BlackmanNuttall)
	}
}

func BenchmarkApply_blackmanHarris(b *testing.B) {
	s := makeNSlice()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		gowindow.Apply(s, gowindow.BlackmanHarris)
	}
}

func BenchmarkApplyNew_blackmanHarris(b *testing.B) {
	s := makeNSlice()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		gowindow.ApplyNew(s, gowindow.BlackmanHarris)
	}
}

func BenchmarkApply_flatTop(b *testing.B) {
	s := makeNSlice()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		gowindow.Apply(s, gowindow.FlatTop)
	}
}

func BenchmarkApplyNew_flatTop(b *testing.B) {
	s := makeNSlice()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		gowindow.ApplyNew(s, gowindow.FlatTop)
	}
}

func BenchmarkApply_rifeVincentClass1Order1(b *testing.B) {
	s := makeNSlice()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		gowindow.Apply(s, gowindow.RifeVincentClass1Order1)
	}
}

func BenchmarkApplyNew_rifeVincentClass1Order1(b *testing.B) {
	s := makeNSlice()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		gowindow.ApplyNew(s, gowindow.RifeVincentClass1Order1)
	}
}

func BenchmarkApply_rifeVincentClass1Order2(b *testing.B) {
	s := makeNSlice()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		gowindow.Apply(s, gowindow.RifeVincentClass1Order2)
	}
}

func BenchmarkApplyNew_rifeVincentClass1Order2(b *testing.B) {
	s := makeNSlice()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		gowindow.ApplyNew(s, gowindow.RifeVincentClass1Order2)
	}
}

func BenchmarkApply_rifeVincentClass1Order3(b *testing.B) {
	s := makeNSlice()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		gowindow.Apply(s, gowindow.RifeVincentClass1Order3)
	}
}

func BenchmarkApplyNew_rifeVincentClass1Order3(b *testing.B) {
	s := makeNSlice()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		gowindow.ApplyNew(s, gowindow.RifeVincentClass1Order3)
	}
}

func BenchmarkApply_rifeVincentClass1Order4(b *testing.B) {
	s := makeNSlice()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		gowindow.Apply(s, gowindow.RifeVincentClass1Order4)
	}
}

func BenchmarkApplyNew_rifeVincentClass1Order4(b *testing.B) {
	s := makeNSlice()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		gowindow.ApplyNew(s, gowindow.RifeVincentClass1Order4)
	}
}

func BenchmarkApply_rifeVincentClass2Decibel36(b *testing.B) {
	s := makeNSlice()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		gowindow.Apply(s, gowindow.RifeVincentClass2Decibel36)
	}
}
func BenchmarkApplyNew_rifeVincentClass2Decibel36(b *testing.B) {
	s := makeNSlice()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		gowindow.ApplyNew(s, gowindow.RifeVincentClass2Decibel36)
	}
}

func BenchmarkApply_rifeVincentClass2Decibel42(b *testing.B) {
	s := makeNSlice()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		gowindow.Apply(s, gowindow.RifeVincentClass2Decibel42)
	}
}
func BenchmarkApplyNew_rifeVincentClass2Decibel42(b *testing.B) {
	s := makeNSlice()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		gowindow.ApplyNew(s, gowindow.RifeVincentClass2Decibel42)
	}
}

func BenchmarkApply_rifeVincentClass2Decibel48(b *testing.B) {
	s := makeNSlice()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		gowindow.Apply(s, gowindow.RifeVincentClass2Decibel48)
	}
}
func BenchmarkApplyNew_rifeVincentClass2Decibel48(b *testing.B) {
	s := makeNSlice()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		gowindow.ApplyNew(s, gowindow.RifeVincentClass2Decibel48)
	}
}

func BenchmarkApply_rifeVincentClass2Decibel54(b *testing.B) {
	s := makeNSlice()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		gowindow.Apply(s, gowindow.RifeVincentClass2Decibel54)
	}
}
func BenchmarkApplyNew_rifeVincentClass2Decibel54(b *testing.B) {
	s := makeNSlice()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		gowindow.ApplyNew(s, gowindow.RifeVincentClass2Decibel54)
	}
}

func BenchmarkApply_rifeVincentClass2Decibel60(b *testing.B) {
	s := makeNSlice()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		gowindow.Apply(s, gowindow.RifeVincentClass2Decibel60)
	}
}
func BenchmarkApplyNew_rifeVincentClass2Decibel60(b *testing.B) {
	s := makeNSlice()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		gowindow.ApplyNew(s, gowindow.RifeVincentClass2Decibel60)
	}
}

func BenchmarkApply_rifeVincentClass2Decibel66(b *testing.B) {
	s := makeNSlice()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		gowindow.Apply(s, gowindow.RifeVincentClass2Decibel66)
	}
}
func BenchmarkApplyNew_rifeVincentClass2Decibel66(b *testing.B) {
	s := makeNSlice()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		gowindow.ApplyNew(s, gowindow.RifeVincentClass2Decibel66)
	}
}

func BenchmarkApply_rifeVincentClass3Order2(b *testing.B) {
	s := makeNSlice()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		gowindow.Apply(s, gowindow.RifeVincentClass3Order2)
	}
}

func BenchmarkApplyNew_rifeVincentClass3Order2(b *testing.B) {
	s := makeNSlice()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		gowindow.ApplyNew(s, gowindow.RifeVincentClass3Order2)
	}
}

func BenchmarkApply_rifeVincentClass3Order3(b *testing.B) {
	s := makeNSlice()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		gowindow.Apply(s, gowindow.RifeVincentClass3Order3)
	}
}

func BenchmarkApplyNew_rifeVincentClass3Order3(b *testing.B) {
	s := makeNSlice()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		gowindow.ApplyNew(s, gowindow.RifeVincentClass3Order3)
	}
}

func BenchmarkApply_rifeVincentClass3Order4(b *testing.B) {
	s := makeNSlice()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		gowindow.Apply(s, gowindow.RifeVincentClass3Order4)
	}
}

func BenchmarkApplyNew_rifeVincentClass3Order4(b *testing.B) {
	s := makeNSlice()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		gowindow.ApplyNew(s, gowindow.RifeVincentClass3Order4)
	}
}
