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
		gowindow.New(gowindow.Rectangular).Apply(s)
	}
}

func BenchmarkApplyNew_rectangular(b *testing.B) {
	s := makeNSlice()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		gowindow.New(gowindow.Rectangular).ApplyNew(s)
	}
}

func BenchmarkApply_triangular(b *testing.B) {
	s := makeNSlice()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		gowindow.New(gowindow.Triangular).Apply(s)
	}
}

func BenchmarkApplyNew_triangular(b *testing.B) {
	s := makeNSlice()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		gowindow.New(gowindow.Triangular).ApplyNew(s)
	}
}

func BenchmarkApply_bartlett(b *testing.B) {
	s := makeNSlice()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		gowindow.New(gowindow.Bartlett).Apply(s)
	}
}

func BenchmarkApplyNew_bartlett(b *testing.B) {
	s := makeNSlice()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		gowindow.New(gowindow.Bartlett).ApplyNew(s)
	}
}

func BenchmarkApply_fejer(b *testing.B) {
	s := makeNSlice()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		gowindow.New(gowindow.Fejer).Apply(s)
	}
}

func BenchmarkApplyNew_fejer(b *testing.B) {
	s := makeNSlice()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		gowindow.New(gowindow.Fejer).ApplyNew(s)
	}
}

func BenchmarkApply_parzen(b *testing.B) {
	s := makeNSlice()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		gowindow.New(gowindow.Parzen).Apply(s)
	}
}

func BenchmarkApplyNew_parzen(b *testing.B) {
	s := makeNSlice()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		gowindow.New(gowindow.Parzen).ApplyNew(s)
	}
}

func BenchmarkApply_deLaValleePoussin(b *testing.B) {
	s := makeNSlice()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		gowindow.New(gowindow.DeLaValleePoussin).Apply(s)
	}
}

func BenchmarkApplyNew_deLaValleePoussin(b *testing.B) {
	s := makeNSlice()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		gowindow.New(gowindow.DeLaValleePoussin).ApplyNew(s)
	}
}

func BenchmarkApply_welch(b *testing.B) {
	s := makeNSlice()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		gowindow.New(gowindow.Welch).Apply(s)
	}
}

func BenchmarkApplyNew_welch(b *testing.B) {
	s := makeNSlice()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		gowindow.New(gowindow.Welch).ApplyNew(s)
	}
}

func BenchmarkApply_sine(b *testing.B) {
	s := makeNSlice()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		gowindow.New(gowindow.Sine).Apply(s)
	}
}

func BenchmarkApplyNew_sine(b *testing.B) {
	s := makeNSlice()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		gowindow.New(gowindow.Sine).ApplyNew(s)
	}
}

func BenchmarkApply_hanning(b *testing.B) {
	s := makeNSlice()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		gowindow.New(gowindow.Hanning).Apply(s)
	}
}

func BenchmarkApplyNew_hanning(b *testing.B) {
	s := makeNSlice()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		gowindow.New(gowindow.Hanning).ApplyNew(s)
	}
}

func BenchmarkApply_hann(b *testing.B) {
	s := makeNSlice()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		gowindow.New(gowindow.Hann).Apply(s)
	}
}

func BenchmarkApplyNew_hann(b *testing.B) {
	s := makeNSlice()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		gowindow.New(gowindow.Hann).ApplyNew(s)
	}
}

func BenchmarkApply_hamming(b *testing.B) {
	s := makeNSlice()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		gowindow.New(gowindow.Hamming).Apply(s)
	}
}

func BenchmarkApplyNew_hamming(b *testing.B) {
	s := makeNSlice()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		gowindow.New(gowindow.Hamming).ApplyNew(s)
	}
}

func BenchmarkApply_blackman(b *testing.B) {
	s := makeNSlice()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		gowindow.New(gowindow.Blackman).Apply(s)
	}
}

func BenchmarkApplyNew_blackman(b *testing.B) {
	s := makeNSlice()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		gowindow.New(gowindow.Blackman).ApplyNew(s)
	}
}

func BenchmarkApply_nuttall(b *testing.B) {
	s := makeNSlice()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		gowindow.New(gowindow.Nuttall).Apply(s)
	}
}

func BenchmarkApplyNew_nuttall(b *testing.B) {
	s := makeNSlice()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		gowindow.New(gowindow.Nuttall).ApplyNew(s)
	}
}

func BenchmarkApply_blackmanNuttall(b *testing.B) {
	s := makeNSlice()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		gowindow.New(gowindow.BlackmanNuttall).Apply(s)
	}
}

func BenchmarkApplyNew_blackmanNuttall(b *testing.B) {
	s := makeNSlice()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		gowindow.New(gowindow.BlackmanNuttall).ApplyNew(s)
	}
}

func BenchmarkApply_blackmanHarris(b *testing.B) {
	s := makeNSlice()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		gowindow.New(gowindow.BlackmanHarris).Apply(s)
	}
}

func BenchmarkApplyNew_blackmanHarris(b *testing.B) {
	s := makeNSlice()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		gowindow.New(gowindow.BlackmanHarris).ApplyNew(s)
	}
}

func BenchmarkApply_flatTop(b *testing.B) {
	s := makeNSlice()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		gowindow.New(gowindow.FlatTop).Apply(s)
	}
}

func BenchmarkApplyNew_flatTop(b *testing.B) {
	s := makeNSlice()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		gowindow.New(gowindow.FlatTop).ApplyNew(s)
	}
}

func BenchmarkApply_rifeVincentClass1Order1(b *testing.B) {
	s := makeNSlice()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		w := gowindow.New(gowindow.RifeVincent)
		_ = w.SetOption(&gowindow.Option{Class: gowindow.Class1, Order: gowindow.Order1})
		w.Apply(s)
	}
}

func BenchmarkApplyNew_rifeVincentClass1Order1(b *testing.B) {
	s := makeNSlice()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		w := gowindow.New(gowindow.RifeVincent)
		_ = w.SetOption(&gowindow.Option{Class: gowindow.Class1, Order: gowindow.Order1})
		w.ApplyNew(s)
	}
}

func BenchmarkApply_rifeVincentClass1Order2(b *testing.B) {
	s := makeNSlice()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		w := gowindow.New(gowindow.RifeVincent)
		_ = w.SetOption(&gowindow.Option{Class: gowindow.Class1, Order: gowindow.Order2})
		w.Apply(s)
	}
}

func BenchmarkApplyNew_rifeVincentClass1Order2(b *testing.B) {
	s := makeNSlice()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		w := gowindow.New(gowindow.RifeVincent)
		_ = w.SetOption(&gowindow.Option{Class: gowindow.Class1, Order: gowindow.Order2})
		w.ApplyNew(s)
	}
}

func BenchmarkApply_rifeVincentClass1Order3(b *testing.B) {
	s := makeNSlice()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		w := gowindow.New(gowindow.RifeVincent)
		_ = w.SetOption(&gowindow.Option{Class: gowindow.Class1, Order: gowindow.Order3})
		w.Apply(s)
	}
}

func BenchmarkApplyNew_rifeVincentClass1Order3(b *testing.B) {
	s := makeNSlice()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		w := gowindow.New(gowindow.RifeVincent)
		_ = w.SetOption(&gowindow.Option{Class: gowindow.Class1, Order: gowindow.Order3})
		w.ApplyNew(s)
	}
}

func BenchmarkApply_rifeVincentClass1Order4(b *testing.B) {
	s := makeNSlice()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		w := gowindow.New(gowindow.RifeVincent)
		_ = w.SetOption(&gowindow.Option{Class: gowindow.Class1, Order: gowindow.Order4})
		w.Apply(s)
	}
}

func BenchmarkApplyNew_rifeVincentClass1Order4(b *testing.B) {
	s := makeNSlice()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		w := gowindow.New(gowindow.RifeVincent)
		_ = w.SetOption(&gowindow.Option{Class: gowindow.Class1, Order: gowindow.Order4})
		w.ApplyNew(s)
	}
}

func BenchmarkApply_rifeVincentClass2Decibel36(b *testing.B) {
	s := makeNSlice()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		w := gowindow.New(gowindow.RifeVincent)
		_ = w.SetOption(&gowindow.Option{Class: gowindow.Class2, Decibel: gowindow.Decibel36})
		w.Apply(s)
	}
}
func BenchmarkApplyNew_rifeVincentClass2Decibel36(b *testing.B) {
	s := makeNSlice()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		w := gowindow.New(gowindow.RifeVincent)
		_ = w.SetOption(&gowindow.Option{Class: gowindow.Class2, Decibel: gowindow.Decibel36})
		w.ApplyNew(s)
	}
}

func BenchmarkApply_rifeVincentClass2Decibel42(b *testing.B) {
	s := makeNSlice()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		w := gowindow.New(gowindow.RifeVincent)
		_ = w.SetOption(&gowindow.Option{Class: gowindow.Class2, Decibel: gowindow.Decibel42})
		w.Apply(s)
	}
}
func BenchmarkApplyNew_rifeVincentClass2Decibel42(b *testing.B) {
	s := makeNSlice()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		w := gowindow.New(gowindow.RifeVincent)
		_ = w.SetOption(&gowindow.Option{Class: gowindow.Class2, Decibel: gowindow.Decibel42})
		w.ApplyNew(s)
	}
}

func BenchmarkApply_rifeVincentClass2Decibel48(b *testing.B) {
	s := makeNSlice()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		w := gowindow.New(gowindow.RifeVincent)
		_ = w.SetOption(&gowindow.Option{Class: gowindow.Class2, Decibel: gowindow.Decibel48})
		w.Apply(s)
	}
}
func BenchmarkApplyNew_rifeVincentClass2Decibel48(b *testing.B) {
	s := makeNSlice()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		w := gowindow.New(gowindow.RifeVincent)
		_ = w.SetOption(&gowindow.Option{Class: gowindow.Class2, Decibel: gowindow.Decibel48})
		w.ApplyNew(s)
	}
}

func BenchmarkApply_rifeVincentClass2Decibel54(b *testing.B) {
	s := makeNSlice()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		w := gowindow.New(gowindow.RifeVincent)
		_ = w.SetOption(&gowindow.Option{Class: gowindow.Class2, Decibel: gowindow.Decibel54})
		w.Apply(s)
	}
}
func BenchmarkApplyNew_rifeVincentClass2Decibel54(b *testing.B) {
	s := makeNSlice()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		w := gowindow.New(gowindow.RifeVincent)
		_ = w.SetOption(&gowindow.Option{Class: gowindow.Class2, Decibel: gowindow.Decibel54})
		w.ApplyNew(s)
	}
}

func BenchmarkApply_rifeVincentClass2Decibel60(b *testing.B) {
	s := makeNSlice()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		w := gowindow.New(gowindow.RifeVincent)
		_ = w.SetOption(&gowindow.Option{Class: gowindow.Class2, Decibel: gowindow.Decibel60})
		w.Apply(s)
	}
}
func BenchmarkApplyNew_rifeVincentClass2Decibel60(b *testing.B) {
	s := makeNSlice()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		w := gowindow.New(gowindow.RifeVincent)
		_ = w.SetOption(&gowindow.Option{Class: gowindow.Class2, Decibel: gowindow.Decibel60})
		w.ApplyNew(s)
	}
}

func BenchmarkApply_rifeVincentClass2Decibel66(b *testing.B) {
	s := makeNSlice()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		w := gowindow.New(gowindow.RifeVincent)
		_ = w.SetOption(&gowindow.Option{Class: gowindow.Class2, Decibel: gowindow.Decibel66})
		w.Apply(s)
	}
}
func BenchmarkApplyNew_rifeVincentClass2Decibel66(b *testing.B) {
	s := makeNSlice()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		w := gowindow.New(gowindow.RifeVincent)
		_ = w.SetOption(&gowindow.Option{Class: gowindow.Class2, Decibel: gowindow.Decibel66})
		w.ApplyNew(s)
	}
}

func BenchmarkApply_rifeVincentClass3Order2(b *testing.B) {
	s := makeNSlice()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		w := gowindow.New(gowindow.RifeVincent)
		_ = w.SetOption(&gowindow.Option{Class: gowindow.Class3, Order: gowindow.Order2})
		w.Apply(s)
	}
}

func BenchmarkApplyNew_rifeVincentClass3Order2(b *testing.B) {
	s := makeNSlice()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		w := gowindow.New(gowindow.RifeVincent)
		_ = w.SetOption(&gowindow.Option{Class: gowindow.Class3, Order: gowindow.Order2})
		w.ApplyNew(s)
	}
}

func BenchmarkApply_rifeVincentClass3Order3(b *testing.B) {
	s := makeNSlice()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		w := gowindow.New(gowindow.RifeVincent)
		_ = w.SetOption(&gowindow.Option{Class: gowindow.Class3, Order: gowindow.Order3})
		w.Apply(s)
	}
}

func BenchmarkApplyNew_rifeVincentClass3Order3(b *testing.B) {
	s := makeNSlice()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		w := gowindow.New(gowindow.RifeVincent)
		_ = w.SetOption(&gowindow.Option{Class: gowindow.Class3, Order: gowindow.Order3})
		w.ApplyNew(s)
	}
}

func BenchmarkApply_rifeVincentClass3Order4(b *testing.B) {
	s := makeNSlice()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		w := gowindow.New(gowindow.RifeVincent)
		_ = w.SetOption(&gowindow.Option{Class: gowindow.Class3, Order: gowindow.Order4})
		w.Apply(s)
	}
}

func BenchmarkApplyNew_rifeVincentClass3Order4(b *testing.B) {
	s := makeNSlice()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		w := gowindow.New(gowindow.RifeVincent)
		_ = w.SetOption(&gowindow.Option{Class: gowindow.Class3, Order: gowindow.Order4})
		w.ApplyNew(s)
	}
}

func BenchmarkApply_gaussian(b *testing.B) {
	s := makeNSlice()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		w := gowindow.New(gowindow.Gaussian)
		_ = w.SetOption(&gowindow.Option{SD: 0.4})
		w.Apply(s)
	}
}

func BenchmarkApplyNew_gaussian(b *testing.B) {
	s := makeNSlice()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		w := gowindow.New(gowindow.Gaussian)
		_ = w.SetOption(&gowindow.Option{SD: 0.4})
		w.ApplyNew(s)
	}
}

func BenchmarkApply_confinedGaussian(b *testing.B) {
	s := makeNSlice()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		w := gowindow.New(gowindow.ConfinedGaussian)
		_ = w.SetOption(&gowindow.Option{SDt: 0.1})
		w.Apply(s)
	}
}

func BenchmarkApplyNew_confinedGaussian(b *testing.B) {
	s := makeNSlice()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		w := gowindow.New(gowindow.ConfinedGaussian)
		_ = w.SetOption(&gowindow.Option{SDt: 0.1})
		w.ApplyNew(s)
	}
}

func BenchmarkApply_approximateConfinedGaussian(b *testing.B) {
	s := makeNSlice()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		w := gowindow.New(gowindow.ApproximateConfinedGaussian)
		_ = w.SetOption(&gowindow.Option{SDt: 0.1})
		w.Apply(s)
	}
}

func BenchmarkApplyNew_approximateConfinedGaussian(b *testing.B) {
	s := makeNSlice()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		w := gowindow.New(gowindow.ApproximateConfinedGaussian)
		_ = w.SetOption(&gowindow.Option{SDt: 0.1})
		w.ApplyNew(s)
	}
}

func BenchmarkApply_generalizedNormal(b *testing.B) {
	s := makeNSlice()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		w := gowindow.New(gowindow.GeneralizedNormal)
		_ = w.SetOption(&gowindow.Option{SD: 0.4, P: 2})
		w.Apply(s)
	}
}

func BenchmarkApplyNew_generalizedNormal(b *testing.B) {
	s := makeNSlice()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		w := gowindow.New(gowindow.GeneralizedNormal)
		_ = w.SetOption(&gowindow.Option{SD: 0.4, P: 2})
		w.ApplyNew(s)
	}
}

func BenchmarkApply_tukey(b *testing.B) {
	s := makeNSlice()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		w := gowindow.New(gowindow.Tukey)
		_ = w.SetOption(&gowindow.Option{Alpha: 0.5})
		w.Apply(s)
	}
}

func BenchmarkApplyNew_tukey(b *testing.B) {
	s := makeNSlice()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		w := gowindow.New(gowindow.Tukey)
		_ = w.SetOption(&gowindow.Option{Alpha: 0.5})
		w.ApplyNew(s)
	}
}

func BenchmarkApply_planckTaper(b *testing.B) {
	s := makeNSlice()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		w := gowindow.New(gowindow.PlanckTaper)
		_ = w.SetOption(&gowindow.Option{Epsilon: 0.1})
		w.Apply(s)
	}
}

func BenchmarkApplyNew_planckTaper(b *testing.B) {
	s := makeNSlice()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		w := gowindow.New(gowindow.PlanckTaper)
		_ = w.SetOption(&gowindow.Option{Epsilon: 0.1})
		w.ApplyNew(s)
	}
}

func BenchmarkApply_kaiser(b *testing.B) {
	s := makeNSlice()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		w := gowindow.New(gowindow.Kaiser)
		_ = w.SetOption(&gowindow.Option{Alpha: 2})
		w.Apply(s)
	}
}

func BenchmarkApplyNew_kaiser(b *testing.B) {
	s := makeNSlice()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		w := gowindow.New(gowindow.Kaiser)
		_ = w.SetOption(&gowindow.Option{Alpha: 2})
		w.ApplyNew(s)
	}
}

func BenchmarkApply_dolphChebyshev(b *testing.B) {
	s := makeNSlice()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		w := gowindow.New(gowindow.DolphChebyshev)
		_ = w.SetOption(&gowindow.Option{OmegaZero: 0.1})
		w.Apply(s)
	}
}

func BenchmarkApplyNew_dolphChebyshev(b *testing.B) {
	s := makeNSlice()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		w := gowindow.New(gowindow.DolphChebyshev)
		_ = w.SetOption(&gowindow.Option{OmegaZero: 0.1})
		w.ApplyNew(s)
	}
}

func BenchmarkApply_ultraspherical(b *testing.B) {
	s := makeNSlice()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		w := gowindow.New(gowindow.Ultraspherical)
		_ = w.SetOption(&gowindow.Option{Mu: -0.5, XZero: 1})
		w.Apply(s)
	}
}

func BenchmarkApplyNew_ultraspherical(b *testing.B) {
	s := makeNSlice()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		w := gowindow.New(gowindow.Ultraspherical)
		_ = w.SetOption(&gowindow.Option{Mu: -0.5, XZero: 1})
		w.ApplyNew(s)
	}
}

func BenchmarkApply_exponential(b *testing.B) {
	s := makeNSlice()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		w := gowindow.New(gowindow.Exponential)
		_ = w.SetOption(&gowindow.Option{T: 512})
		w.Apply(s)
	}
}

func BenchmarkApplyNew_exponential(b *testing.B) {
	s := makeNSlice()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		w := gowindow.New(gowindow.Exponential)
		_ = w.SetOption(&gowindow.Option{T: 512})
		w.ApplyNew(s)
	}
}

func BenchmarkApply_poisson(b *testing.B) {
	s := makeNSlice()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		w := gowindow.New(gowindow.Poisson)
		_ = w.SetOption(&gowindow.Option{T: 512})
		w.Apply(s)
	}
}

func BenchmarkApplyNew_poisson(b *testing.B) {
	s := makeNSlice()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		w := gowindow.New(gowindow.Poisson)
		_ = w.SetOption(&gowindow.Option{T: 512})
		w.ApplyNew(s)
	}
}

func BenchmarkApply_bartlettHann(b *testing.B) {
	s := makeNSlice()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		gowindow.New(gowindow.BartlettHann).Apply(s)
	}
}

func BenchmarkApplyNew_bartlettHann(b *testing.B) {
	s := makeNSlice()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		gowindow.New(gowindow.BartlettHann).ApplyNew(s)
	}
}

func BenchmarkApply_planckBessel(b *testing.B) {
	s := makeNSlice()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		w := gowindow.New(gowindow.PlanckBessel)
		_ = w.SetOption(&gowindow.Option{Epsilon: 0.1, Alpha: 4.45})
		w.Apply(s)
	}
}

func BenchmarkApplyNew_planckBessel(b *testing.B) {
	s := makeNSlice()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		w := gowindow.New(gowindow.PlanckBessel)
		_ = w.SetOption(&gowindow.Option{Epsilon: 0.1, Alpha: 4.45})
		w.ApplyNew(s)
	}
}

func BenchmarkApply_hannPoisson(b *testing.B) {
	s := makeNSlice()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		w := gowindow.New(gowindow.HannPoisson)
		_ = w.SetOption(&gowindow.Option{Alpha: 2})
		w.Apply(s)
	}
}

func BenchmarkApplyNew_hannPoisson(b *testing.B) {
	s := makeNSlice()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		w := gowindow.New(gowindow.HannPoisson)
		_ = w.SetOption(&gowindow.Option{Alpha: 2})
		w.ApplyNew(s)
	}
}
