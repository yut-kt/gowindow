package gowindow_test

import (
	"fmt"
	"log"

	"github.com/yut-kt/gowindow"
)

func getTestSlice() []float64 {
	return []float64{1, 1, 1, 1, 1, 1, 1, 1, 1}
}

func ExampleNew_apply_rectangular() {
	s := getTestSlice()
	gowindow.New(gowindow.Rectangular).Apply(s)
	for i := range s {
		fmt.Printf("%f ", s[i])
	}
	fmt.Println()
	// Output:
	// 1.000000 1.000000 1.000000 1.000000 1.000000 1.000000 1.000000 1.000000 1.000000
}

func ExampleNew_applyNew_rectangular() {
	s := getTestSlice()
	for _, x := range gowindow.New(gowindow.Rectangular).ApplyNew(s) {
		fmt.Printf("%f ", x)
	}
	fmt.Println()
	// Output:
	// 1.000000 1.000000 1.000000 1.000000 1.000000 1.000000 1.000000 1.000000 1.000000
}

func ExampleNew_apply_triangular() {
	s := getTestSlice()
	gowindow.New(gowindow.Triangular).Apply(s)
	for i := range s {
		fmt.Printf("%f ", s[i])
	}
	fmt.Println()
	// Output:
	// 0.000000 0.250000 0.500000 0.750000 1.000000 0.750000 0.500000 0.250000 0.000000
}
func ExampleNew_applyNew_triangular() {
	s := getTestSlice()
	for _, x := range gowindow.New(gowindow.Triangular).ApplyNew(s) {
		fmt.Printf("%f ", x)
	}
	fmt.Println()
	// Output:
	// 0.000000 0.250000 0.500000 0.750000 1.000000 0.750000 0.500000 0.250000 0.000000
}

func ExampleNew_apply_bartlett() {
	s := getTestSlice()
	gowindow.New(gowindow.Bartlett).Apply(s)
	for i := range s {
		fmt.Printf("%f ", s[i])
	}
	fmt.Println()
	// Output:
	// 0.000000 0.250000 0.500000 0.750000 1.000000 0.750000 0.500000 0.250000 0.000000
}

func ExampleNew_applyNew_bartlett() {
	s := getTestSlice()
	for _, x := range gowindow.New(gowindow.Bartlett).ApplyNew(s) {
		fmt.Printf("%f ", x)
	}
	fmt.Println()
	// Output:
	// 0.000000 0.250000 0.500000 0.750000 1.000000 0.750000 0.500000 0.250000 0.000000
}

func ExampleNew_apply_fejer() {
	s := getTestSlice()
	gowindow.New(gowindow.Fejer).Apply(s)
	for i := range s {
		fmt.Printf("%f ", s[i])
	}
	fmt.Println()
	// Output:
	// 0.000000 0.250000 0.500000 0.750000 1.000000 0.750000 0.500000 0.250000 0.000000
}

func ExampleNew_applyNew_fejer() {
	s := getTestSlice()
	for _, x := range gowindow.New(gowindow.Fejer).ApplyNew(s) {
		fmt.Printf("%f ", x)
	}
	fmt.Println()
	// Output:
	// 0.000000 0.250000 0.500000 0.750000 1.000000 0.750000 0.500000 0.250000 0.000000
}

func ExampleNew_apply_parzen() {
	s := getTestSlice()
	gowindow.New(gowindow.Parzen).Apply(s)
	for i := range s {
		fmt.Printf("%f ", s[i])
	}
	fmt.Println()
	// Output:
	// 0.002000 0.054000 0.250000 0.622000 0.946000 0.946000 0.622000 0.250000 0.054000
}

func ExampleNew_applyNew_parzen() {
	s := getTestSlice()
	for _, x := range gowindow.New(gowindow.Parzen).ApplyNew(s) {
		fmt.Printf("%f ", x)
	}
	fmt.Println()
	// Output:
	// 0.002000 0.054000 0.250000 0.622000 0.946000 0.946000 0.622000 0.250000 0.054000
}

func ExampleNew_apply_deLaValleePoussin() {
	s := getTestSlice()
	gowindow.New(gowindow.DeLaValleePoussin).Apply(s)
	for i := range s {
		fmt.Printf("%f ", s[i])
	}
	fmt.Println()
	// Output:
	// 0.002000 0.054000 0.250000 0.622000 0.946000 0.946000 0.622000 0.250000 0.054000
}

func ExampleNew_applyNew_deLaValleePoussin() {
	s := getTestSlice()
	for _, x := range gowindow.New(gowindow.DeLaValleePoussin).ApplyNew(s) {
		fmt.Printf("%f ", x)
	}
	fmt.Println()
	// Output:
	// 0.002000 0.054000 0.250000 0.622000 0.946000 0.946000 0.622000 0.250000 0.054000
}

func ExampleNew_apply_welch() {
	s := getTestSlice()
	gowindow.New(gowindow.Welch).Apply(s)
	for i := range s {
		fmt.Printf("%f ", s[i])
	}
	fmt.Println()
	// Output:
	// 0.000000 0.437500 0.750000 0.937500 1.000000 0.937500 0.750000 0.437500 0.000000
}

func ExampleNew_applyNew_welch() {
	s := getTestSlice()
	for _, x := range gowindow.New(gowindow.Welch).ApplyNew(s) {
		fmt.Printf("%f ", x)
	}
	fmt.Println()
	// Output:
	// 0.000000 0.437500 0.750000 0.937500 1.000000 0.937500 0.750000 0.437500 0.000000
}

func ExampleNew_apply_sine() {
	s := getTestSlice()
	gowindow.New(gowindow.Sine).Apply(s)
	for i := range s {
		fmt.Printf("%f ", s[i])
	}
	fmt.Println()
	// Output:
	// 0.000000 0.382683 0.707107 0.923880 1.000000 0.923880 0.707107 0.382683 0.000000
}

func ExampleNew_applyNew_sine() {
	s := getTestSlice()
	for _, x := range gowindow.New(gowindow.Sine).ApplyNew(s) {
		fmt.Printf("%f ", x)
	}
	fmt.Println()
	// Output:
	// 0.000000 0.382683 0.707107 0.923880 1.000000 0.923880 0.707107 0.382683 0.000000
}

func ExampleNew_apply_hanning() {
	s := getTestSlice()
	gowindow.New(gowindow.Hanning).Apply(s)
	for i := range s {
		fmt.Printf("%f ", s[i])
	}
	fmt.Println()
	// Output:
	// 0.000000 0.146447 0.500000 0.853553 1.000000 0.853553 0.500000 0.146447 0.000000
}

func ExampleNew_applyNew_hanning() {
	s := getTestSlice()
	for _, x := range gowindow.New(gowindow.Hanning).ApplyNew(s) {
		fmt.Printf("%f ", x)
	}
	fmt.Println()
	// Output:
	// 0.000000 0.146447 0.500000 0.853553 1.000000 0.853553 0.500000 0.146447 0.000000
}

func ExampleNew_apply_hann() {
	s := getTestSlice()
	gowindow.New(gowindow.Hann).Apply(s)
	for i := range s {
		fmt.Printf("%f ", s[i])
	}
	fmt.Println()
	// Output:
	// 0.000000 0.146447 0.500000 0.853553 1.000000 0.853553 0.500000 0.146447 0.000000
}

func ExampleNew_applyNew_hann() {
	s := getTestSlice()
	for _, x := range gowindow.New(gowindow.Hann).ApplyNew(s) {
		fmt.Printf("%f ", x)
	}
	fmt.Println()
	// Output:
	// 0.000000 0.146447 0.500000 0.853553 1.000000 0.853553 0.500000 0.146447 0.000000
}

func ExampleNew_apply_hamming() {
	s := getTestSlice()
	gowindow.New(gowindow.Hamming).Apply(s)
	for i := range s {
		fmt.Printf("%f ", s[i])
	}
	fmt.Println()
	// Output:
	// 0.080000 0.214731 0.540000 0.865269 1.000000 0.865269 0.540000 0.214731 0.080000
}

func ExampleNew_applyNew_hamming() {
	s := getTestSlice()
	for _, x := range gowindow.New(gowindow.Hamming).ApplyNew(s) {
		fmt.Printf("%f ", x)
	}
	fmt.Println()
	// Output:
	// 0.080000 0.214731 0.540000 0.865269 1.000000 0.865269 0.540000 0.214731 0.080000
}

func ExampleNew_apply_blackman() {
	s := getTestSlice()
	gowindow.New(gowindow.Blackman).Apply(s)
	for i := range s {
		fmt.Printf("%f ", s[i])
	}
	fmt.Println()
	// Output:
	// -0.000000 0.066447 0.340000 0.773553 1.000000 0.773553 0.340000 0.066447 -0.000000
}

func ExampleNew_applyNew_blackman() {
	s := getTestSlice()
	for _, x := range gowindow.New(gowindow.Blackman).ApplyNew(s) {
		fmt.Printf("%f ", x)
	}
	fmt.Println()
	// Output:
	// -0.000000 0.066447 0.340000 0.773553 1.000000 0.773553 0.340000 0.066447 -0.000000
}

func ExampleNew_apply_nuttall() {
	s := getTestSlice()
	gowindow.New(gowindow.Nuttall).Apply(s)
	for i := range s {
		fmt.Printf("%f ", s[i])
	}
	fmt.Println()
	// Output:
	// -0.000000 0.020039 0.211536 0.691497 1.000000 0.691497 0.211536 0.020039 -0.000000
}

func ExampleNew_applyNew_nuttall() {
	s := getTestSlice()
	for _, x := range gowindow.New(gowindow.Nuttall).ApplyNew(s) {
		fmt.Printf("%f ", x)
	}
	fmt.Println()
	// Output:
	// -0.000000 0.020039 0.211536 0.691497 1.000000 0.691497 0.211536 0.020039 -0.000000
}

func ExampleNew_apply_blackmanNuttall() {
	s := getTestSlice()
	gowindow.New(gowindow.BlackmanNuttall).Apply(s)
	for i := range s {
		fmt.Printf("%f ", s[i])
	}
	fmt.Println()
	// Output:
	// 0.000363 0.025206 0.226982 0.701958 1.000000 0.701958 0.226982 0.025206 0.000363
}

func ExampleNew_applyNew_blackmanNuttall() {
	s := getTestSlice()
	for _, x := range gowindow.New(gowindow.BlackmanNuttall).ApplyNew(s) {
		fmt.Printf("%f ", x)
	}
	fmt.Println()
	// Output:
	// 0.000363 0.025206 0.226982 0.701958 1.000000 0.701958 0.226982 0.025206 0.000363
}

func ExampleNew_apply_blackmanHarris() {
	s := getTestSlice()
	gowindow.New(gowindow.BlackmanHarris).Apply(s)
	for i := range s {
		fmt.Printf("%f ", s[i])
	}
	fmt.Println()
	// Output:
	// 0.000060 0.021736 0.217470 0.695764 1.000000 0.695764 0.217470 0.021736 0.000060
}

func ExampleNew_applyNew_blackmanHarris() {
	s := getTestSlice()
	for _, x := range gowindow.New(gowindow.BlackmanHarris).ApplyNew(s) {
		fmt.Printf("%f ", x)
	}
	fmt.Println()
	// Output:
	// 0.000060 0.021736 0.217470 0.695764 1.000000 0.695764 0.217470 0.021736 0.000060
}

func ExampleNew_apply_flatTop() {
	s := getTestSlice()
	gowindow.New(gowindow.FlatTop).Apply(s)
	for i := range s {
		fmt.Printf("%f ", s[i])
	}
	fmt.Println()
	// Output:
	// -0.000421 -0.026872 -0.054737 0.444135 1.000000 0.444135 -0.054737 -0.026872 -0.000421
}

func ExampleNew_applyNew_flatTop() {
	s := getTestSlice()
	for _, x := range gowindow.New(gowindow.FlatTop).ApplyNew(s) {
		fmt.Printf("%f ", x)
	}
	fmt.Println()
	// Output:
	// -0.000421 -0.026872 -0.054737 0.444135 1.000000 0.444135 -0.054737 -0.026872 -0.000421
}

func ExampleNew_apply_rifeVincent() {
	options := []gowindow.Option{
		{Class: gowindow.Class1, Order: gowindow.Order1},
		{Class: gowindow.Class1, Order: gowindow.Order2},
		{Class: gowindow.Class1, Order: gowindow.Order3},
		{Class: gowindow.Class1, Order: gowindow.Order4},
		{Class: gowindow.Class2, Decibel: gowindow.Decibel36},
		{Class: gowindow.Class2, Decibel: gowindow.Decibel42},
		{Class: gowindow.Class2, Decibel: gowindow.Decibel48},
		{Class: gowindow.Class2, Decibel: gowindow.Decibel54},
		{Class: gowindow.Class2, Decibel: gowindow.Decibel60},
		{Class: gowindow.Class2, Decibel: gowindow.Decibel66},
		{Class: gowindow.Class3, Order: gowindow.Order2},
		{Class: gowindow.Class3, Order: gowindow.Order3},
		{Class: gowindow.Class3, Order: gowindow.Order4},
	}
	for _, o := range options {
		s := getTestSlice()
		w := gowindow.New(gowindow.RifeVincent)
		err := w.SetOption(&o)
		if err != nil {
			log.Fatal(err)
		}
		w.Apply(s)
		fmt.Printf("%e\n", s)
	}
	// Output:
	// [2.000000e+00 1.707107e+00 1.000000e+00 2.928932e-01 0.000000e+00 2.928932e-01 1.000000e+00 1.707107e+00 2.000000e+00]
	// [5.551115e-17 5.719096e-02 6.666667e-01 1.942809e+00 2.666667e+00 1.942809e+00 6.666667e-01 5.719096e-02 5.551115e-17]
	// [-2.775558e-17 1.005051e-02 4.000000e-01 1.989949e+00 3.200000e+00 1.989949e+00 4.000000e-01 1.005051e-02 -2.775558e-17]
	// [-3.816392e-17 1.682129e-03 2.285714e-01 1.941175e+00 3.657143e+00 1.941175e+00 2.285714e-01 1.682129e-03 -3.816392e-17]
	// [-2.024927e+39 2.025112e+39 -2.025559e+39 2.026007e+39 -2.026193e+39 2.026007e+39 -2.025559e+39 2.025112e+39 -2.024927e+39]
	// [-2.003435e+46 3.226211e+42 2.003891e+46 -3.233214e+42 -2.004349e+46 -3.233214e+42 2.003891e+46 3.226211e+42 -2.003435e+46]
	// [-2.182846e+53 -2.182956e+53 -2.183224e+53 -2.183492e+53 -2.183603e+53 -2.183492e+53 -2.183224e+53 -2.182956e+53 -2.182846e+53]
	// [-2.538939e+60 -2.448396e+56 2.539285e+60 2.451575e+56 -2.539632e+60 2.451575e+56 2.539285e+60 -2.448396e+56 -2.538939e+60]
	// [-3.097092e+67 3.097191e+67 -3.097432e+67 3.097673e+67 -3.097773e+67 3.097673e+67 -3.097432e+67 3.097191e+67 -3.097092e+67]
	// [-3.493872e+74 2.237626e+70 3.494188e+74 -2.239554e+70 -3.494505e+74 -2.239554e+70 3.494188e+74 2.237626e+70 -3.493872e+74]
	// [2.775558e-17 1.536992e-01 8.031500e-01 1.846301e+00 2.393700e+00 1.846301e+00 8.031500e-01 1.536992e-01 2.775558e-17]
	// [1.000000e-06 2.816375e-02 5.024630e-01 1.971836e+00 2.995073e+00 1.971836e+00 5.024630e-01 2.816375e-02 1.000000e-06]
	// [-3.546000e-03 2.294166e-03 2.924730e-01 1.961864e+00 3.490284e+00 1.961864e+00 2.924730e-01 2.294166e-03 -3.546000e-03]
}

func ExampleNew_applyNew_rifeVincent() {
	options := []gowindow.Option{
		{Class: gowindow.Class1, Order: gowindow.Order1},
		{Class: gowindow.Class1, Order: gowindow.Order2},
		{Class: gowindow.Class1, Order: gowindow.Order3},
		{Class: gowindow.Class1, Order: gowindow.Order4},
		{Class: gowindow.Class2, Decibel: gowindow.Decibel36},
		{Class: gowindow.Class2, Decibel: gowindow.Decibel42},
		{Class: gowindow.Class2, Decibel: gowindow.Decibel48},
		{Class: gowindow.Class2, Decibel: gowindow.Decibel54},
		{Class: gowindow.Class2, Decibel: gowindow.Decibel60},
		{Class: gowindow.Class2, Decibel: gowindow.Decibel66},
		{Class: gowindow.Class3, Order: gowindow.Order2},
		{Class: gowindow.Class3, Order: gowindow.Order3},
		{Class: gowindow.Class3, Order: gowindow.Order4},
	}

	for _, o := range options {
		s := getTestSlice()
		w := gowindow.New(gowindow.RifeVincent)
		err := w.SetOption(&o)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%e\n", w.ApplyNew(s))
	}
	// Output:
	// [2.000000e+00 1.707107e+00 1.000000e+00 2.928932e-01 0.000000e+00 2.928932e-01 1.000000e+00 1.707107e+00 2.000000e+00]
	// [5.551115e-17 5.719096e-02 6.666667e-01 1.942809e+00 2.666667e+00 1.942809e+00 6.666667e-01 5.719096e-02 5.551115e-17]
	// [-2.775558e-17 1.005051e-02 4.000000e-01 1.989949e+00 3.200000e+00 1.989949e+00 4.000000e-01 1.005051e-02 -2.775558e-17]
	// [-3.816392e-17 1.682129e-03 2.285714e-01 1.941175e+00 3.657143e+00 1.941175e+00 2.285714e-01 1.682129e-03 -3.816392e-17]
	// [-2.024927e+39 2.025112e+39 -2.025559e+39 2.026007e+39 -2.026193e+39 2.026007e+39 -2.025559e+39 2.025112e+39 -2.024927e+39]
	// [-2.003435e+46 3.226211e+42 2.003891e+46 -3.233214e+42 -2.004349e+46 -3.233214e+42 2.003891e+46 3.226211e+42 -2.003435e+46]
	// [-2.182846e+53 -2.182956e+53 -2.183224e+53 -2.183492e+53 -2.183603e+53 -2.183492e+53 -2.183224e+53 -2.182956e+53 -2.182846e+53]
	// [-2.538939e+60 -2.448396e+56 2.539285e+60 2.451575e+56 -2.539632e+60 2.451575e+56 2.539285e+60 -2.448396e+56 -2.538939e+60]
	// [-3.097092e+67 3.097191e+67 -3.097432e+67 3.097673e+67 -3.097773e+67 3.097673e+67 -3.097432e+67 3.097191e+67 -3.097092e+67]
	// [-3.493872e+74 2.237626e+70 3.494188e+74 -2.239554e+70 -3.494505e+74 -2.239554e+70 3.494188e+74 2.237626e+70 -3.493872e+74]
	// [2.775558e-17 1.536992e-01 8.031500e-01 1.846301e+00 2.393700e+00 1.846301e+00 8.031500e-01 1.536992e-01 2.775558e-17]
	// [1.000000e-06 2.816375e-02 5.024630e-01 1.971836e+00 2.995073e+00 1.971836e+00 5.024630e-01 2.816375e-02 1.000000e-06]
	// [-3.546000e-03 2.294166e-03 2.924730e-01 1.961864e+00 3.490284e+00 1.961864e+00 2.924730e-01 2.294166e-03 -3.546000e-03]
}

func ExampleNew_apply_gaussian() {
	s := getTestSlice()
	w := gowindow.New(gowindow.Gaussian)
	err := w.SetOption(&gowindow.Option{SD: 0.4})
	w.Apply(s)
	if err != nil {
		log.Fatal(err)
	}
	for n := range s {
		fmt.Printf("%f ", s[n])
	}
	fmt.Println()
	// Output:
	// 0.043937 0.172422 0.457833 0.822578 1.000000 0.822578 0.457833 0.172422 0.043937
}

func ExampleNew_applyNew_gaussian() {
	s := getTestSlice()
	w := gowindow.New(gowindow.Gaussian)
	err := w.SetOption(&gowindow.Option{SD: 0.4})
	if err != nil {
		log.Fatal(err)
	}
	for _, x := range w.ApplyNew(s) {
		fmt.Printf("%f ", x)
	}
	fmt.Println()
	// Output:
	// 0.043937 0.172422 0.457833 0.822578 1.000000 0.822578 0.457833 0.172422 0.043937
}

func ExampleNew_apply_gauss() {
	s := getTestSlice()
	w := gowindow.New(gowindow.Gauss)
	err := w.SetOption(&gowindow.Option{SD: 0.4})
	w.Apply(s)
	if err != nil {
		log.Fatal(err)
	}
	for n := range s {
		fmt.Printf("%f ", s[n])
	}
	fmt.Println()
	// Output:
	// 0.043937 0.172422 0.457833 0.822578 1.000000 0.822578 0.457833 0.172422 0.043937
}

func ExampleNew_applyNew_gauss() {
	s := getTestSlice()
	w := gowindow.New(gowindow.Gauss)
	err := w.SetOption(&gowindow.Option{SD: 0.4})
	if err != nil {
		log.Fatal(err)
	}
	for _, x := range w.ApplyNew(s) {
		fmt.Printf("%f ", x)
	}
	fmt.Println()
	// Output:
	// 0.043937 0.172422 0.457833 0.822578 1.000000 0.822578 0.457833 0.172422 0.043937
}

func ExampleNew_apply_confinedGaussian() {
	s := getTestSlice()
	w := gowindow.New(gowindow.ConfinedGaussian)
	err := w.SetOption(&gowindow.Option{SDt: 0.1})
	w.Apply(s)
	if err != nil {
		log.Fatal(err)
	}
	for n := range s {
		fmt.Printf("%f ", s[n])
	}
	fmt.Println()
	// Output:
	// 0.016812 0.105341 0.367878 0.778801 1.000000 0.778801 0.367878 0.105341 0.016812
}

func ExampleNew_applyNew_confinedGaussian() {
	s := getTestSlice()
	w := gowindow.New(gowindow.ConfinedGaussian)
	err := w.SetOption(&gowindow.Option{SDt: 0.1})
	if err != nil {
		log.Fatal(err)
	}
	for _, x := range w.ApplyNew(s) {
		fmt.Printf("%f ", x)
	}
	fmt.Println()
	// Output:
	// 0.016812 0.105341 0.367878 0.778801 1.000000 0.778801 0.367878 0.105341 0.016812
}

func ExampleNew_apply_approximateConfinedGaussian() {
	s := getTestSlice()
	w := gowindow.New(gowindow.ApproximateConfinedGaussian)
	err := w.SetOption(&gowindow.Option{SDt: 0.1})
	w.Apply(s)
	if err != nil {
		log.Fatal(err)
	}
	for n := range s {
		fmt.Printf("%f ", s[n])
	}
	fmt.Println()
	// Output:
	// 0.016812 0.105341 0.367878 0.778801 1.000000 0.778801 0.367878 0.105341 0.016812
}

func ExampleNew_applyNew_approximateConfinedGaussian() {
	s := getTestSlice()
	w := gowindow.New(gowindow.ApproximateConfinedGaussian)
	err := w.SetOption(&gowindow.Option{SDt: 0.1})
	if err != nil {
		log.Fatal(err)
	}
	for _, x := range w.ApplyNew(s) {
		fmt.Printf("%f ", x)
	}
	fmt.Println()
	// Output:
	// 0.016812 0.105341 0.367878 0.778801 1.000000 0.778801 0.367878 0.105341 0.016812
}

func ExampleNew_apply_generalizedNormal() {
	s := getTestSlice()
	w := gowindow.New(gowindow.GeneralizedNormal)
	err := w.SetOption(&gowindow.Option{SD: 0.4, P: 2})
	w.Apply(s)
	if err != nil {
		log.Fatal(err)
	}
	for n := range s {
		fmt.Printf("%f ", s[n])
	}
	fmt.Println()
	// Output:
	// 0.001930 0.029729 0.209611 0.676634 1.000000 0.676634 0.209611 0.029729 0.001930
}

func ExampleNew_applyNew_generalizedNormal() {
	s := getTestSlice()
	w := gowindow.New(gowindow.GeneralizedNormal)
	err := w.SetOption(&gowindow.Option{SD: 0.4, P: 2})
	if err != nil {
		log.Fatal(err)
	}
	for _, x := range w.ApplyNew(s) {
		fmt.Printf("%f ", x)
	}
	fmt.Println()
	// Output:
	// 0.001930 0.029729 0.209611 0.676634 1.000000 0.676634 0.209611 0.029729 0.001930
}

func ExampleNew_apply_tukey() {
	s := getTestSlice()
	w := gowindow.New(gowindow.Tukey)
	err := w.SetOption(&gowindow.Option{Alpha: 0.5})
	w.Apply(s)
	if err != nil {
		log.Fatal(err)
	}
	for n := range s {
		fmt.Printf("%f ", s[n])
	}
	fmt.Println()
	// Output:
	// 0.000000 0.500000 1.000000 1.000000 1.000000 1.000000 1.000000 0.500000 0.000000
}

func ExampleNew_applyNew_tukey() {
	s := getTestSlice()
	w := gowindow.New(gowindow.Tukey)
	err := w.SetOption(&gowindow.Option{Alpha: 0.5})
	if err != nil {
		log.Fatal(err)
	}
	for _, x := range w.ApplyNew(s) {
		fmt.Printf("%f ", x)
	}
	fmt.Println()
	// Output:
	// 0.000000 0.500000 1.000000 1.000000 1.000000 1.000000 1.000000 0.500000 0.000000
}

func ExampleNew_apply_planckTaper() {
	s := getTestSlice()
	w := gowindow.New(gowindow.PlanckTaper)
	err := w.SetOption(&gowindow.Option{Epsilon: 0.1})
	w.Apply(s)
	if err != nil {
		log.Fatal(err)
	}
	for n := range s {
		fmt.Printf("%f ", s[n])
	}
	fmt.Println()
	// Output:
	// 0.000000 1.000000 1.000000 1.000000 1.000000 1.000000 1.000000 1.000000 0.000000
}

func ExampleNew_applyNew_planckTaper() {
	s := getTestSlice()
	w := gowindow.New(gowindow.PlanckTaper)
	err := w.SetOption(&gowindow.Option{Epsilon: 0.1})
	if err != nil {
		log.Fatal(err)
	}
	for _, x := range w.ApplyNew(s) {
		fmt.Printf("%f ", x)
	}
	fmt.Println()
	// Output:
	// 0.000000 1.000000 1.000000 1.000000 1.000000 1.000000 1.000000 1.000000 0.000000
}

func ExampleNew_apply_kaiser() {
	s := getTestSlice()
	w := gowindow.New(gowindow.Kaiser)
	err := w.SetOption(&gowindow.Option{Alpha: 2})
	w.Apply(s)
	if err != nil {
		log.Fatal(err)
	}
	for n := range s {
		fmt.Printf("%f ", s[n])
	}
	fmt.Println()
	// Output:
	// 4.539740 -1.735678 -0.122286 0.785754 1.000000 0.785754 -0.122286 -1.735678 4.539740
}

func ExampleNew_applyNew_kaiser() {
	s := getTestSlice()
	w := gowindow.New(gowindow.Kaiser)
	err := w.SetOption(&gowindow.Option{Alpha: 2})
	if err != nil {
		log.Fatal(err)
	}
	for _, x := range w.ApplyNew(s) {
		fmt.Printf("%f ", x)
	}
	fmt.Println()
	// Output:
	// 4.539740 -1.735678 -0.122286 0.785754 1.000000 0.785754 -0.122286 -1.735678 4.539740
}

func ExampleNew_apply_dolphChebyshev() {
	s := getTestSlice()
	w := gowindow.New(gowindow.DolphChebyshev)
	err := w.SetOption(&gowindow.Option{OmegaZero: 0.1})
	w.Apply(s)
	if err != nil {
		log.Fatal(err)
	}
	for n := range s {
		fmt.Printf("%f ", s[n])
	}
	fmt.Println()
	// Output:
	// 0.197401 0.174267 0.105529 0.033874 0.003743 0.033874 0.105529 0.174267 0.197401
}

func ExampleNew_applyNew_dolphChebyshev() {
	s := getTestSlice()
	w := gowindow.New(gowindow.DolphChebyshev)
	err := w.SetOption(&gowindow.Option{OmegaZero: 0.1})
	if err != nil {
		log.Fatal(err)
	}
	for _, x := range w.ApplyNew(s) {
		fmt.Printf("%f ", x)
	}
	fmt.Println()
	// Output:
	// 0.197401 0.174267 0.105529 0.033874 0.003743 0.033874 0.105529 0.174267 0.197401
}

func ExampleNew_apply_ultraspherical() {
	s := getTestSlice()
	w := gowindow.New(gowindow.Ultraspherical)
	err := w.SetOption(&gowindow.Option{Mu: -0.5, XZero: 1})
	w.Apply(s)
	if err != nil {
		log.Fatal(err)
	}
	for n := range s {
		fmt.Printf("%f ", s[n])
	}
	fmt.Println()
	// Output:
	// 0.000118 0.001023 0.000114 0.002987 -0.000173 -0.008020 -0.000173 0.002987 0.000114
}

func ExampleNew_applyNew_ultraspherical() {
	s := getTestSlice()
	w := gowindow.New(gowindow.Ultraspherical)
	err := w.SetOption(&gowindow.Option{Mu: -0.5, XZero: 1})
	if err != nil {
		log.Fatal(err)
	}
	for _, x := range w.ApplyNew(s) {
		fmt.Printf("%f ", x)
	}
	fmt.Println()
	// Output:
	// 0.000118 0.001023 0.000114 0.002987 -0.000173 -0.008020 -0.000173 0.002987 0.000114
}

func ExampleNew_apply_exponential() {
	s := getTestSlice()
	w := gowindow.New(gowindow.Exponential)
	err := w.SetOption(&gowindow.Option{T: float64(len(s)) / 2})
	w.Apply(s)
	if err != nil {
		log.Fatal(err)
	}
	for n := range s {
		fmt.Printf("%f ", s[n])
	}
	fmt.Println()
	// Output:
	// 0.411112 0.513417 0.641180 0.800737 1.000000 0.800737 0.641180 0.513417 0.411112
}

func ExampleNew_applyNew_exponential() {
	s := getTestSlice()
	w := gowindow.New(gowindow.Exponential)
	err := w.SetOption(&gowindow.Option{T: float64(len(s)) / 2})
	if err != nil {
		log.Fatal(err)
	}
	for _, x := range w.ApplyNew(s) {
		fmt.Printf("%f ", x)
	}
	fmt.Println()
	// Output:
	// 0.411112 0.513417 0.641180 0.800737 1.000000 0.800737 0.641180 0.513417 0.411112
}

func ExampleNew_apply_poisson() {
	s := getTestSlice()
	w := gowindow.New(gowindow.Poisson)
	err := w.SetOption(&gowindow.Option{T: float64(len(s)) / 2})
	w.Apply(s)
	if err != nil {
		log.Fatal(err)
	}
	for n := range s {
		fmt.Printf("%f ", s[n])
	}
	fmt.Println()
	// Output:
	// 0.411112 0.513417 0.641180 0.800737 1.000000 0.800737 0.641180 0.513417 0.411112
}

func ExampleNew_applyNew_poisson() {
	s := getTestSlice()
	w := gowindow.New(gowindow.Poisson)
	err := w.SetOption(&gowindow.Option{T: float64(len(s)) / 2})
	if err != nil {
		log.Fatal(err)
	}
	for _, x := range w.ApplyNew(s) {
		fmt.Printf("%f ", x)
	}
	fmt.Println()
	// Output:
	// 0.411112 0.513417 0.641180 0.800737 1.000000 0.800737 0.641180 0.513417 0.411112
}

func ExampleNew_apply_bartlettHann() {
	s := getTestSlice()
	gowindow.New(gowindow.BartlettHann).Apply(s)
	for i := range s {
		fmt.Printf("%f ", s[i])
	}
	fmt.Println()
	// Output:
	// 0.000000 0.171299 0.500000 0.828701 1.000000 0.828701 0.500000 0.171299 0.000000
}

func ExampleNew_applyNew_bartlettHann() {
	s := getTestSlice()
	for _, x := range gowindow.New(gowindow.BartlettHann).ApplyNew(s) {
		fmt.Printf("%f ", x)
	}
	fmt.Println()
	// Output:
	// 0.000000 0.171299 0.500000 0.828701 1.000000 0.828701 0.500000 0.171299 0.000000
}

func ExampleNew_apply_planckBessel() {
	s := getTestSlice()
	w := gowindow.New(gowindow.PlanckBessel)
	err := w.SetOption(&gowindow.Option{Epsilon: 0.1, Alpha: 4.45})
	w.Apply(s)
	if err != nil {
		log.Fatal(err)
	}
	for n := range s {
		fmt.Printf("%f ", s[n])
	}
	fmt.Println()
	// Output:
	// 2.241608 -0.845019 0.409901 1.229007 1.000000 1.229007 0.409901 -0.845019 3.400478
}

func ExampleNew_applyNew_planckBessel() {
	s := getTestSlice()
	w := gowindow.New(gowindow.PlanckBessel)
	err := w.SetOption(&gowindow.Option{Epsilon: 0.1, Alpha: 4.45})
	if err != nil {
		log.Fatal(err)
	}
	for _, x := range w.ApplyNew(s) {
		fmt.Printf("%f ", x)
	}
	fmt.Println()
	// Output:
	// 2.241608 -0.845019 0.409901 1.229007 1.000000 1.229007 0.409901 -0.845019 3.400478
}

func ExampleNew_apply_missedSwitchImplementation() {
	s := getTestSlice()
	gowindow.New(gowindow.None).Apply(s)
	for i := range s {
		fmt.Printf("%f ", s[i])
	}
	fmt.Println()
	// Output:
	// 1.000000 1.000000 1.000000 1.000000 1.000000 1.000000 1.000000 1.000000 1.000000
}

func ExampleNew_applyNew_missedSwitchImplementation() {
	s := getTestSlice()
	for _, x := range gowindow.New(gowindow.None).ApplyNew(s) {
		fmt.Printf("%f ", x)
	}
	fmt.Println()
	// Output:
	//
}
