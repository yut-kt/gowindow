package gowindow_test

import (
	"testing"

	"github.com/yut-kt/gowindow"
)

func BenchmarkApply_rectangular(b *testing.B) {
	gowindow.Apply([]float64{1, 2, 3, 4}, gowindow.Rectangular)
}

func BenchmarkApplyNew_rectangular(b *testing.B) {
	gowindow.ApplyNew([]float64{1, 2, 3, 4}, gowindow.Rectangular)
}

func BenchmarkApply_triangular(b *testing.B) {
	gowindow.Apply([]float64{1, 2, 3, 4}, gowindow.Triangular)
}

func BenchmarkApplyNew_triangular(b *testing.B) {
	gowindow.ApplyNew([]float64{1, 2, 3, 4}, gowindow.Triangular)
}

func BenchmarkApply_bartlett(b *testing.B) {
	gowindow.Apply([]float64{1, 2, 3, 4}, gowindow.Bartlett)
}

func BenchmarkApplyNew_bartlett(b *testing.B) {
	gowindow.ApplyNew([]float64{1, 2, 3, 4}, gowindow.Bartlett)
}

func BenchmarkApply_fejer(b *testing.B) {
	gowindow.Apply([]float64{1, 2, 3, 4}, gowindow.Fejer)
}

func BenchmarkApplyNew_fejer(b *testing.B) {
	gowindow.ApplyNew([]float64{1, 2, 3, 4}, gowindow.Fejer)
}

func BenchmarkApply_parzen(b *testing.B) {
	gowindow.Apply([]float64{1, 2, 3, 4}, gowindow.Parzen)
}

func BenchmarkApplyNew_parzen(b *testing.B) {
	gowindow.ApplyNew([]float64{1, 2, 3, 4}, gowindow.Parzen)
}

func BenchmarkApply_deLaValleePoussin(b *testing.B) {
	gowindow.Apply([]float64{1, 2, 3, 4}, gowindow.DeLaValleePoussin)
}

func BenchmarkApplyNew_deLaValleePoussin(b *testing.B) {
	gowindow.ApplyNew([]float64{1, 2, 3, 4}, gowindow.DeLaValleePoussin)
}

func BenchmarkApply_welch(b *testing.B) {
	gowindow.Apply([]float64{1, 2, 3, 4}, gowindow.Welch)
}

func BenchmarkApplyNew_welch(b *testing.B) {
	gowindow.ApplyNew([]float64{1, 2, 3, 4}, gowindow.Welch)
}

func BenchmarkApply_hanning(b *testing.B) {
	gowindow.Apply([]float64{1, 2, 3, 4}, gowindow.Hanning)
}

func BenchmarkApplyNew_hanning(b *testing.B) {
	gowindow.ApplyNew([]float64{1, 2, 3, 4}, gowindow.Hanning)
}

func BenchmarkApply_hann(b *testing.B) {
	gowindow.Apply([]float64{1, 2, 3, 4}, gowindow.Hann)
}

func BenchmarkApplyNew_hann(b *testing.B) {
	gowindow.ApplyNew([]float64{1, 2, 3, 4}, gowindow.Hann)
}

func BenchmarkApply_hamming(b *testing.B) {
	gowindow.Apply([]float64{1, 2, 3, 4}, gowindow.Hamming)
}

func BenchmarkApplyNew_hamming(b *testing.B) {
	gowindow.ApplyNew([]float64{1, 2, 3, 4}, gowindow.Hamming)
}

func BenchmarkApply_blackman(b *testing.B) {
	gowindow.Apply([]float64{1, 2, 3, 4}, gowindow.Blackman)
}

func BenchmarkApplyNew_blackman(b *testing.B) {
	gowindow.ApplyNew([]float64{1, 2, 3, 4}, gowindow.Blackman)
}
