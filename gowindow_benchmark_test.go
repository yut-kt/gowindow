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
