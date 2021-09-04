package gowindow_test

import (
	"fmt"

	"github.com/yut-kt/gowindow"
)

func ExampleRectangular() {
	for _, x := range gowindow.Rectangular([]float64{1, 2, 3, 4}) {
		fmt.Printf("%f ", x)
	}
	fmt.Println()
	// Output:
	// 1.000000 2.000000 3.000000 4.000000
}

func ExampleTriangular() {
	for _, x := range gowindow.Triangular([]float64{1, 2, 3, 4, 5}) {
		fmt.Printf("%f ", x)
	}
	fmt.Println()
	// Output:
	// 0.000000 1.000000 3.000000 2.000000 0.000000
}

func ExampleBartlett() {
	for _, x := range gowindow.Bartlett([]float64{1, 2, 3, 4, 5}) {
		fmt.Printf("%f ", x)
	}
	fmt.Println()
	// Output:
	// 0.000000 1.000000 3.000000 2.000000 0.000000
}

func ExampleFejer() {
	for _, x := range gowindow.Fejer([]float64{1, 2, 3, 4, 5}) {
		fmt.Printf("%f ", x)
	}
	fmt.Println()
	// Output:
	// 0.000000 1.000000 3.000000 2.000000 0.000000
}

func ExampleTriangularD() {
	s := []float64{1, 2, 3, 4, 5}
	gowindow.TriangularD(s)
	for i := range s {
		fmt.Printf("%f ", s[i])
	}
	fmt.Println()
	// Output:
	// 0.000000 1.000000 3.000000 2.000000 0.000000
}

func ExampleBartlettD() {
	s := []float64{1, 2, 3, 4, 5}
	gowindow.BartlettD(s)
	for i := range s {
		fmt.Printf("%f ", s[i])
	}
	fmt.Println()
	// Output:
	// 0.000000 1.000000 3.000000 2.000000 0.000000
}

func ExampleFejerD() {
	s := []float64{1, 2, 3, 4, 5}
	gowindow.FejerD(s)
	for i := range s {
		fmt.Printf("%f ", s[i])
	}
	fmt.Println()
	// Output:
	// 0.000000 1.000000 3.000000 2.000000 0.000000
}

func ExampleHamming() {
	for _, x := range gowindow.Hamming([]float64{1, 2, 3, 4}) {
		fmt.Printf("%f ", x)
	}
	fmt.Println()
	// Output:
	// 0.080000 1.540000 2.310000 0.320000
}

func ExampleHammingD() {
	s := []float64{1, 2, 3, 4}
	gowindow.HammingD(s)
	for i := range s {
		fmt.Printf("%f ", s[i])
	}
	fmt.Println()
	// Output:
	// 0.080000 1.540000 2.310000 0.320000
}

func ExampleHanning() {
	for _, x := range gowindow.Hanning([]float64{1, 2, 3, 4}) {
		fmt.Printf("%f ", x)
	}
	fmt.Println()
	// Output:
	// 0.000000 1.500000 2.250000 0.000000
}

func ExampleHann() {
	for _, x := range gowindow.Hann([]float64{1, 2, 3, 4}) {
		fmt.Printf("%f ", x)
	}
	fmt.Println()
	// Output:
	// 0.000000 1.500000 2.250000 0.000000
}

func ExampleHanningD() {
	s := []float64{1, 2, 3, 4}
	gowindow.HanningD(s)
	for i := range s {
		fmt.Printf("%f ", s[i])
	}
	fmt.Println()
	// Output:
	// 0.000000 1.500000 2.250000 0.000000
}

func ExampleHannD() {
	s := []float64{1, 2, 3, 4}
	gowindow.HannD(s)
	for i := range s {
		fmt.Printf("%f ", s[i])
	}
	fmt.Println()
	// Output:
	// 0.000000 1.500000 2.250000 0.000000
}

func ExampleBlackman() {
	for _, x := range gowindow.Blackman([]float64{1, 2, 3, 4}) {
		fmt.Printf("%f ", x)
	}
	fmt.Println()
	// Output:
	// -0.000000 1.260000 1.890000 -0.000000
}

func ExampleBlackmanD() {
	s := []float64{1, 2, 3, 4}
	gowindow.BlackmanD(s)
	for i := range s {
		fmt.Printf("%f ", s[i])
	}
	fmt.Println()
	// Output:
	// -0.000000 1.260000 1.890000 -0.000000
}
