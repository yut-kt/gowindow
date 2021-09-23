package gowindow_test

import (
	"fmt"
	"github.com/yut-kt/gowindow"
	"log"
)

func getTestSlice() []float64 {
	s := make([]float64, 9)
	for i := range s {
		s[i] = 1
	}
	return s
}

// Group1 is Rectangular window

func ExampleRectangular() {
	s := getTestSlice()
	for _, x := range gowindow.Rectangular(s) {
		fmt.Printf("%f ", x)
	}
	fmt.Println()
	// Output:
	// 1.000000 1.000000 1.000000 1.000000 1.000000 1.000000 1.000000 1.000000 1.000000
}

func ExampleBoxcar() {
	s := getTestSlice()
	for _, x := range gowindow.Boxcar(s) {
		fmt.Printf("%f ", x)
	}
	fmt.Println()
	// Output:
	// 1.000000 1.000000 1.000000 1.000000 1.000000 1.000000 1.000000 1.000000 1.000000
}

func ExampleDirichlet() {
	s := getTestSlice()
	for _, x := range gowindow.Dirichlet(s) {
		fmt.Printf("%f ", x)
	}
	fmt.Println()
	// Output:
	// 1.000000 1.000000 1.000000 1.000000 1.000000 1.000000 1.000000 1.000000 1.000000
}

// Group2 is B-spline windows

func ExampleTriangular() {
	s := getTestSlice()
	for _, x := range gowindow.Triangular(s) {
		fmt.Printf("%f ", x)
	}
	// Output:
	// 0.000000 0.250000 0.500000 0.750000 1.000000 0.750000 0.500000 0.250000 0.000000
}

func ExampleBartlett() {
	s := getTestSlice()
	for _, x := range gowindow.Bartlett(s) {
		fmt.Printf("%f ", x)
	}
	fmt.Println()
	// Output:
	// 0.000000 0.250000 0.500000 0.750000 1.000000 0.750000 0.500000 0.250000 0.000000
}

func ExampleFejer() {
	s := getTestSlice()
	for _, x := range gowindow.Fejer(s) {
		fmt.Printf("%f ", x)
	}
	fmt.Println()
	// Output:
	// 0.000000 0.250000 0.500000 0.750000 1.000000 0.750000 0.500000 0.250000 0.000000
}

func ExampleParzen() {
	s := getTestSlice()
	for _, x := range gowindow.Parzen(s) {
		fmt.Printf("%f ", x)
	}
	fmt.Println()
	// Output:
	// 0.250000 0.472656 0.718750 0.917969 1.000000 0.917969 0.718750 0.472656 0.250000
}

func ExampleDeLaValleePoussin() {
	s := getTestSlice()
	for _, x := range gowindow.DeLaValleePoussin(s) {
		fmt.Printf("%f ", x)
	}
	fmt.Println()
	// Output:
	// 0.250000 0.472656 0.718750 0.917969 1.000000 0.917969 0.718750 0.472656 0.250000
}

func ExampleWelch() {
	s := getTestSlice()
	for _, x := range gowindow.Welch(s) {
		fmt.Printf("%f ", x)
	}
	fmt.Println()
	// Output:
	// 0.000000 0.437500 0.750000 0.937500 1.000000 0.937500 0.750000 0.437500 0.000000
}

func ExamplePowerOfSine() {
	s := getTestSlice()
	for _, x := range gowindow.PowerOfCosine(s, 1) {
		fmt.Printf("%f ", x)
	}
	fmt.Println()
	// Output:
	// 0.000000 0.382683 0.707107 0.923880 1.000000 0.923880 0.707107 0.382683 0.000000
}

func ExampleNew_applyNew_powerOfSine() {
	s := getTestSlice()
	for _, x := range gowindow.PowerOfSine(s, 1) {
		fmt.Printf("%f ", x)
	}
	fmt.Println()
	// Output:
	// 0.000000 0.382683 0.707107 0.923880 1.000000 0.923880 0.707107 0.382683 0.000000
}

func ExampleSine() {
	s := getTestSlice()
	for _, x := range gowindow.Sine(s) {
		fmt.Printf("%f ", x)
	}
	fmt.Println()
	// Output:
	// 0.000000 0.382683 0.707107 0.923880 1.000000 0.923880 0.707107 0.382683 0.000000
}

func ExampleCosine() {
	s := getTestSlice()
	for _, x := range gowindow.Cosine(s) {
		fmt.Printf("%f ", x)
	}
	fmt.Println()
	// Output:
	// 0.000000 0.382683 0.707107 0.923880 1.000000 0.923880 0.707107 0.382683 0.000000
}

func ExampleHann() {
	s := getTestSlice()
	for _, x := range gowindow.Hann(s) {
		fmt.Printf("%f ", x)
	}
	fmt.Println()
	// Output:
	// 0.000000 0.146447 0.500000 0.853553 1.000000 0.853553 0.500000 0.146447 0.000000
}

func ExampleHanning() {
	s := getTestSlice()
	for _, x := range gowindow.Hanning(s) {
		fmt.Printf("%f ", x)
	}
	fmt.Println()
	// Output:
	// 0.000000 0.146447 0.500000 0.853553 1.000000 0.853553 0.500000 0.146447 0.000000
}

func ExampleHamming() {
	s := getTestSlice()
	for _, x := range gowindow.Hamming(s) {
		fmt.Printf("%f ", x)
	}
	fmt.Println()
	// Output:
	// 0.076720 0.211931 0.538360 0.864789 1.000000 0.864789 0.538360 0.211931 0.076720
}

func ExampleHammingOriginal() {
	s := getTestSlice()
	for _, x := range gowindow.HammingOriginal(s) {
		fmt.Printf("%f ", x)
	}
	fmt.Println()
	// Output:
	// 0.080000 0.214731 0.540000 0.865269 1.000000 0.865269 0.540000 0.214731 0.080000
}

func ExampleBlackman() {
	s := getTestSlice()
	for _, x := range gowindow.Blackman(s) {
		fmt.Printf("%f ", x)
	}
	fmt.Println()
	// Output:
	// -0.000000 0.066447 0.340000 0.773553 1.000000 0.773553 0.340000 0.066447 -0.000000
}

func ExampleBlackmanExact() {
	s := getTestSlice()
	for _, x := range gowindow.BlackmanExact(s) {
		fmt.Printf("%f ", x)
	}
	fmt.Println()
	// Output:
	// 0.006879 0.075469 0.349742 0.777712 1.000000 0.777712 0.349742 0.075469 0.006879
}

func ExampleNuttall() {
	s := getTestSlice()
	for _, x := range gowindow.Nuttall(s) {
		fmt.Printf("%f ", x)
	}
	fmt.Println()
	// Output:
	// -0.000000 0.020039 0.211536 0.691497 1.000000 0.691497 0.211536 0.020039 -0.000000
}

func ExampleBlackmanNuttall() {
	s := getTestSlice()
	for _, x := range gowindow.BlackmanNuttall(s) {
		fmt.Printf("%f ", x)
	}
	fmt.Println()
	// Output:
	// 0.000363 0.025206 0.226982 0.701958 1.000000 0.701958 0.226982 0.025206 0.000363
}

func ExampleBlackmanHarris() {
	s := getTestSlice()
	for _, x := range gowindow.BlackmanHarris(s) {
		fmt.Printf("%f ", x)
	}
	fmt.Println()
	// Output:
	// 0.000060 0.021736 0.217470 0.695764 1.000000 0.695764 0.217470 0.021736 0.000060
}

func ExampleNew_applyNew_flatTop() {
	s := getTestSlice()
	for _, x := range gowindow.FlatTop(s) {
		fmt.Printf("%f ", x)
	}
	fmt.Println()
	// Output:
	// -0.000421 -0.026872 -0.054737 0.444135 1.000000 0.444135 -0.054737 -0.026872 -0.000421
}

func ExampleRifeVincentClass1() {
	s := getTestSlice()
	orders := []int{1, 2, 3, 4}
	for i := range orders {
		if sw, err := gowindow.RifeVincentClass1(s, orders[i]); err != nil {
			log.Fatal(err)
		} else {
			fmt.Printf("%e\n", sw)
		}
	}
	// Output:
	// [2.000000e+00 1.707107e+00 1.000000e+00 2.928932e-01 0.000000e+00 2.928932e-01 1.000000e+00 1.707107e+00 2.000000e+00]
	// [5.551115e-17 5.719096e-02 6.666667e-01 1.942809e+00 2.666667e+00 1.942809e+00 6.666667e-01 5.719096e-02 5.551115e-17]
	// [-2.775558e-17 1.005051e-02 4.000000e-01 1.989949e+00 3.200000e+00 1.989949e+00 4.000000e-01 1.005051e-02 -2.775558e-17]
	// [-3.816392e-17 1.682129e-03 2.285714e-01 1.941175e+00 3.657143e+00 1.941175e+00 2.285714e-01 1.682129e-03 -3.816392e-17]
}

func ExampleRifeVincentClass2() {
	s := getTestSlice()
	decibels := []int{36, 42, 48, 54, 60, 66}
	for i := range decibels {
		if sw, err := gowindow.RifeVincentClass2(s, decibels[i]); err != nil {
			log.Fatal(err)
		} else {
			fmt.Printf("%e\n", sw)
		}
	}
	// Output:
	// [-2.024927e+39 2.025112e+39 -2.025559e+39 2.026007e+39 -2.026193e+39 2.026007e+39 -2.025559e+39 2.025112e+39 -2.024927e+39]
	// [-2.003435e+46 3.226211e+42 2.003891e+46 -3.233214e+42 -2.004349e+46 -3.233214e+42 2.003891e+46 3.226211e+42 -2.003435e+46]
	// [-2.182846e+53 -2.182956e+53 -2.183224e+53 -2.183492e+53 -2.183603e+53 -2.183492e+53 -2.183224e+53 -2.182956e+53 -2.182846e+53]
	// [-2.538939e+60 -2.448396e+56 2.539285e+60 2.451575e+56 -2.539632e+60 2.451575e+56 2.539285e+60 -2.448396e+56 -2.538939e+60]
	// [-3.097092e+67 3.097191e+67 -3.097432e+67 3.097673e+67 -3.097773e+67 3.097673e+67 -3.097432e+67 3.097191e+67 -3.097092e+67]
	// [-3.493872e+74 2.237626e+70 3.494188e+74 -2.239554e+70 -3.494505e+74 -2.239554e+70 3.494188e+74 2.237626e+70 -3.493872e+74]
}

func ExampleRifeVincentClass3() {
	s := getTestSlice()
	orders := []int{2, 3, 4}
	for i := range orders {
		if sw, err := gowindow.RifeVincentClass3(s, orders[i]); err != nil {
			log.Fatal(err)
		} else {
			fmt.Printf("%e\n", sw)
		}
	}
	// Output:
	// [2.775558e-17 1.536992e-01 8.031500e-01 1.846301e+00 2.393700e+00 1.846301e+00 8.031500e-01 1.536992e-01 2.775558e-17]
	// [1.000000e-06 2.816375e-02 5.024630e-01 1.971836e+00 2.995073e+00 1.971836e+00 5.024630e-01 2.816375e-02 1.000000e-06]
	// [-3.546000e-03 2.294166e-03 2.924730e-01 1.961864e+00 3.490284e+00 1.961864e+00 2.924730e-01 2.294166e-03 -3.546000e-03]
}

func ExampleGaussian() {
	s := getTestSlice()
	gaussian, err := gowindow.Gaussian(s, 0.4)
	if err != nil {
		log.Fatal(err)
	}
	for i := range gaussian {
		fmt.Printf("%f ", gaussian[i])
	}
	fmt.Println()
	// Output:
	// 0.043937 0.172422 0.457833 0.822578 1.000000 0.822578 0.457833 0.172422 0.043937
}

func ExampleNew_applyNew_gauss() {
	s := getTestSlice()
	gaussian, err := gowindow.Gauss(s, 0.4)
	if err != nil {
		log.Fatal(err)
	}
	for i := range gaussian {
		fmt.Printf("%f ", gaussian[i])
	}
	fmt.Println()
	// Output:
	// 0.043937 0.172422 0.457833 0.822578 1.000000 0.822578 0.457833 0.172422 0.043937
}

func ExampleConfinedGaussian() {
	s := getTestSlice()
	for _, x := range gowindow.ConfinedGaussian(s, 0.1) {
		fmt.Printf("%f ", x)
	}
	fmt.Println()
	// Output:
	// 0.016812 0.105341 0.367878 0.778801 1.000000 0.778801 0.367878 0.105341 0.016812
}

func ExampleApproximateConfinedGaussian() {
	s := getTestSlice()
	for _, x := range gowindow.ApproximateConfinedGaussian(s, 0.1) {
		fmt.Printf("%f ", x)
	}
	fmt.Println()
	// Output:
	// 0.016812 0.105341 0.367878 0.778801 1.000000 0.778801 0.367878 0.105341 0.016812
}

func ExampleGeneralizedNormal() {
	s := getTestSlice()
	for _, x := range gowindow.GeneralizedNormal(s, 0.4, 2) {
		fmt.Printf("%f ", x)
	}
	fmt.Println()
	// Output:
	// 0.001930 0.029729 0.209611 0.676634 1.000000 0.676634 0.209611 0.029729 0.001930
}

func ExampleTukey() {
	s := getTestSlice()
	for _, x := range gowindow.Tukey(s, 0.5) {
		fmt.Printf("%f ", x)
	}
	fmt.Println()
	// Output:
	// 0.000000 0.500000 1.000000 1.000000 1.000000 1.000000 1.000000 0.500000 0.000000
}

func ExamplePlanckTaper() {
	s := getTestSlice()
	for _, x := range gowindow.PlanckTaper(s, 0.1) {
		fmt.Printf("%f ", x)
	}
	fmt.Println()
	// Output:
	// 0.000000 1.000000 1.000000 1.000000 1.000000 1.000000 1.000000 1.000000 0.000000
}

func ExampleKaiser() {
	s := getTestSlice()
	for _, x := range gowindow.Kaiser(s, 2) {
		fmt.Printf("%f ", x)
	}
	fmt.Println()
	// Output:
	// 4.539740 -1.735678 -0.122286 0.785754 1.000000 0.785754 -0.122286 -1.735678 4.539740
}

func ExampleExponential() {
	s := getTestSlice()
	for _, x := range gowindow.Exponential(s, float64(len(s)/2)) {
		fmt.Printf("%f ", x)
	}
	fmt.Println()
	// Output:
	// 0.367879 0.472367 0.606531 0.778801 1.000000 0.778801 0.606531 0.472367 0.367879
}

func ExamplePoisson() {
	s := getTestSlice()
	for _, x := range gowindow.Poisson(s, float64(len(s)/2)) {
		fmt.Printf("%f ", x)
	}
	fmt.Println()
	// Output:
	// 0.367879 0.472367 0.606531 0.778801 1.000000 0.778801 0.606531 0.472367 0.367879
}

func ExampleBartlettHann() {
	s := getTestSlice()
	for _, x := range gowindow.BartlettHann(s) {
		fmt.Printf("%f ", x)
	}
	fmt.Println()
	// Output:
	// 0.000000 0.171299 0.500000 0.828701 1.000000 0.828701 0.500000 0.171299 0.000000
}

func ExampleHannPoisson() {
	s := getTestSlice()
	for _, x := range gowindow.HannPoisson(s, 2) {
		fmt.Printf("%f ", x)
	}
	fmt.Println()
	// Output:
	// 0.000000 0.032677 0.183940 0.517706 1.000000 0.517706 0.183940 0.032677 0.000000
}

func ExampleLanczos() {
	s := getTestSlice()
	for _, x := range gowindow.Lanczos(s) {
		fmt.Printf("%f ", x)
	}
	fmt.Println()
	// Output:
	// 0.000000 0.300105 0.636620 0.900316 1.000000 0.900316 0.636620 0.300105 0.000000
}
