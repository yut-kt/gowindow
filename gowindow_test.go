package gowindow_test

import (
	"fmt"

	"github.com/yut-kt/gowindow"
)

func ExampleApply_rectangular() {
	s := []float64{1, 2, 3, 4, 5}
	gowindow.Apply(s, gowindow.Rectangular)
	for i := range s {
		fmt.Printf("%f ", s[i])
	}
	fmt.Println()
	// Output:
	// 1.000000 2.000000 3.000000 4.000000 5.000000
}

func ExampleApplyNew_rectangular() {
	for _, x := range gowindow.ApplyNew([]float64{1, 2, 3, 4, 5}, gowindow.Rectangular) {
		fmt.Printf("%f ", x)
	}
	fmt.Println()
	// Output:
	// 1.000000 2.000000 3.000000 4.000000 5.000000
}

func ExampleApply_triangular() {
	s := []float64{1, 2, 3, 4, 5}
	gowindow.Apply(s, gowindow.Triangular)
	for i := range s {
		fmt.Printf("%f ", s[i])
	}
	fmt.Println()
	// Output:
	// 0.000000 1.000000 3.000000 2.000000 0.000000
}
func ExampleApplyNew_triangular() {
	for _, x := range gowindow.ApplyNew([]float64{1, 2, 3, 4, 5}, gowindow.Triangular) {
		fmt.Printf("%f ", x)
	}
	fmt.Println()
	// Output:
	// 0.000000 1.000000 3.000000 2.000000 0.000000
}

func ExampleApply_bartlett() {
	s := []float64{1, 2, 3, 4, 5}
	gowindow.Apply(s, gowindow.Bartlett)
	for i := range s {
		fmt.Printf("%f ", s[i])
	}
	fmt.Println()
	// Output:
	// 0.000000 1.000000 3.000000 2.000000 0.000000
}

func ExampleApplyNew_bartlett() {
	for _, x := range gowindow.ApplyNew([]float64{1, 2, 3, 4, 5}, gowindow.Bartlett) {
		fmt.Printf("%f ", x)
	}
	fmt.Println()
	// Output:
	// 0.000000 1.000000 3.000000 2.000000 0.000000
}

func ExampleApply_fejer() {
	s := []float64{1, 2, 3, 4, 5}
	gowindow.Apply(s, gowindow.Fejer)
	for i := range s {
		fmt.Printf("%f ", s[i])
	}
	fmt.Println()
	// Output:
	// 0.000000 1.000000 3.000000 2.000000 0.000000
}

func ExampleApplyNew_fejer() {
	for _, x := range gowindow.ApplyNew([]float64{1, 2, 3, 4, 5}, gowindow.Fejer) {
		fmt.Printf("%f ", x)
	}
	fmt.Println()
	// Output:
	// 0.000000 1.000000 3.000000 2.000000 0.000000
}

func ExampleApply_parzen() {
	s := []float64{1, 2, 3, 4, 5}
	gowindow.Apply(s, gowindow.Parzen)
	for i := range s {
		fmt.Printf("%f ", s[i])
	}
	fmt.Println()
	// Output:
	// 0.009259 0.500000 2.583333 3.444444 1.250000
}

func ExampleApplyNew_parzen() {
	for _, x := range gowindow.ApplyNew([]float64{1, 2, 3, 4, 5}, gowindow.Parzen) {
		fmt.Printf("%f ", x)
	}
	fmt.Println()
	// Output:
	// 0.009259 0.500000 2.583333 3.444444 1.250000
}

func ExampleApply_deLaValleePoussin() {
	s := []float64{1, 2, 3, 4, 5}
	gowindow.Apply(s, gowindow.DeLaValleePoussin)
	for i := range s {
		fmt.Printf("%f ", s[i])
	}
	fmt.Println()
	// Output:
	// 0.009259 0.500000 2.583333 3.444444 1.250000
}

func ExampleApplyNew_deLaValleePoussin() {
	for _, x := range gowindow.ApplyNew([]float64{1, 2, 3, 4, 5}, gowindow.DeLaValleePoussin) {
		fmt.Printf("%f ", x)
	}
	fmt.Println()
	// Output:
	// 0.009259 0.500000 2.583333 3.444444 1.250000
}

func ExampleApply_welch() {
	s := []float64{1, 1, 1, 1, 1}
	gowindow.Apply(s, gowindow.Welch)
	for i := range s {
		fmt.Printf("%f ", s[i])
	}
	fmt.Println()
	// Output:
	// 0.000000 0.750000 1.000000 0.750000 0.000000
}

func ExampleApplyNew_welch() {
	for _, x := range gowindow.ApplyNew([]float64{1, 1, 1, 1, 1}, gowindow.Welch) {
		fmt.Printf("%f ", x)
	}
	fmt.Println()
	// Output:
	// 0.000000 0.750000 1.000000 0.750000 0.000000
}

func ExampleApply_sine() {
	s := []float64{1, 1, 1, 1, 1}
	gowindow.Apply(s, gowindow.Sine)
	for i := range s {
		fmt.Printf("%f ", s[i])
	}
	fmt.Println()
	// Output:
	// 0.000000 0.707107 1.000000 0.707107 0.000000
}

func ExampleApplyNew_sine() {
	for _, x := range gowindow.ApplyNew([]float64{1, 1, 1, 1, 1}, gowindow.Sine) {
		fmt.Printf("%f ", x)
	}
	fmt.Println()
	// Output:
	// 0.000000 0.707107 1.000000 0.707107 0.000000
}

func ExampleApply_hanning() {
	s := []float64{1, 2, 3, 4}
	gowindow.Apply(s, gowindow.Hanning)
	for i := range s {
		fmt.Printf("%f ", s[i])
	}
	fmt.Println()
	// Output:
	// 0.000000 1.500000 2.250000 0.000000
}

func ExampleApplyNew_hanning() {
	for _, x := range gowindow.ApplyNew([]float64{1, 2, 3, 4}, gowindow.Hanning) {
		fmt.Printf("%f ", x)
	}
	fmt.Println()
	// Output:
	// 0.000000 1.500000 2.250000 0.000000
}

func ExampleApply_hann() {
	s := []float64{1, 2, 3, 4}
	gowindow.Apply(s, gowindow.Hann)
	for i := range s {
		fmt.Printf("%f ", s[i])
	}
	fmt.Println()
	// Output:
	// 0.000000 1.500000 2.250000 0.000000
}

func ExampleApplyNew_hann() {
	for _, x := range gowindow.ApplyNew([]float64{1, 2, 3, 4}, gowindow.Hann) {
		fmt.Printf("%f ", x)
	}
	fmt.Println()
	// Output:
	// 0.000000 1.500000 2.250000 0.000000
}

func ExampleApply_hamming() {
	s := []float64{1, 2, 3, 4}
	gowindow.Apply(s, gowindow.Hamming)
	for i := range s {
		fmt.Printf("%f ", s[i])
	}
	fmt.Println()
	// Output:
	// 0.080000 1.540000 2.310000 0.320000
}

func ExampleApplyNew_hamming() {
	for _, x := range gowindow.ApplyNew([]float64{1, 2, 3, 4}, gowindow.Hamming) {
		fmt.Printf("%f ", x)
	}
	fmt.Println()
	// Output:
	// 0.080000 1.540000 2.310000 0.320000
}

func ExampleApply_blackman() {
	s := []float64{1, 2, 3, 4}
	gowindow.Apply(s, gowindow.Blackman)
	for i := range s {
		fmt.Printf("%f ", s[i])
	}
	fmt.Println()
	// Output:
	// -0.000000 1.260000 1.890000 -0.000000
}

func ExampleApplyNew_blackman() {
	for _, x := range gowindow.ApplyNew([]float64{1, 2, 3, 4}, gowindow.Blackman) {
		fmt.Printf("%f ", x)
	}
	fmt.Println()
	// Output:
	// -0.000000 1.260000 1.890000 -0.000000
}
