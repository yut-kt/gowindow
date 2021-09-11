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

func ExampleApply_nuttall() {
	s := []float64{1, 2, 3, 4}
	gowindow.Apply(s, gowindow.Nuttall)
	for i := range s {
		fmt.Printf("%f ", s[i])
	}
	fmt.Println()
	// Output:
	// -0.000000 1.029492 1.544238 -0.000000
}

func ExampleApplyNew_nuttall() {
	for _, x := range gowindow.ApplyNew([]float64{1, 2, 3, 4}, gowindow.Nuttall) {
		fmt.Printf("%f ", x)
	}
	fmt.Println()
	// Output:
	// -0.000000 1.029492 1.544238 -0.000000
}

func ExampleApply_blackmanNuttall() {
	s := []float64{1, 1, 1, 1, 1, 1, 1, 1, 1}
	gowindow.Apply(s, gowindow.BlackmanNuttall)
	for i := range s {
		fmt.Printf("%f ", s[i])
	}
	fmt.Println()
	// Output:
	// 0.000363 0.025206 0.226982 0.701958 1.000000 0.701958 0.226982 0.025206 0.000363
}

func ExampleApplyNew_blackmanNuttall() {
	s := []float64{1, 1, 1, 1, 1, 1, 1, 1, 1}
	for _, x := range gowindow.ApplyNew(s, gowindow.BlackmanNuttall) {
		fmt.Printf("%f ", x)
	}
	fmt.Println()
	// Output:
	// 0.000363 0.025206 0.226982 0.701958 1.000000 0.701958 0.226982 0.025206 0.000363
}

func ExampleApply_blackmanHarris() {
	s := []float64{1, 1, 1, 1, 1, 1, 1, 1, 1}
	gowindow.Apply(s, gowindow.BlackmanHarris)
	for i := range s {
		fmt.Printf("%f ", s[i])
	}
	fmt.Println()
	// Output:
	// 0.000060 0.021736 0.217470 0.695764 1.000000 0.695764 0.217470 0.021736 0.000060
}

func ExampleApplyNew_blackmanHarris() {
	s := []float64{1, 1, 1, 1, 1, 1, 1, 1, 1}
	for _, x := range gowindow.ApplyNew(s, gowindow.BlackmanHarris) {
		fmt.Printf("%f ", x)
	}
	fmt.Println()
	// Output:
	// 0.000060 0.021736 0.217470 0.695764 1.000000 0.695764 0.217470 0.021736 0.000060
}

func ExampleApply_flatTop() {
	s := []float64{1, 1, 1, 1, 1, 1, 1, 1, 1}
	gowindow.Apply(s, gowindow.FlatTop)
	for i := range s {
		fmt.Printf("%f ", s[i])
	}
	fmt.Println()
	// Output:
	// -0.000421 -0.026872 -0.054737 0.444135 1.000000 0.444135 -0.054737 -0.026872 -0.000421
}

func ExampleApplyNew_flatTop() {
	s := []float64{1, 1, 1, 1, 1, 1, 1, 1, 1}
	for _, x := range gowindow.ApplyNew(s, gowindow.FlatTop) {
		fmt.Printf("%f ", x)
	}
	fmt.Println()
	// Output:
	// -0.000421 -0.026872 -0.054737 0.444135 1.000000 0.444135 -0.054737 -0.026872 -0.000421
}

func ExampleApply_rifeVincentClass1Order1() {
	s := []float64{1, 1, 1, 1, 1, 1, 1, 1, 1}
	gowindow.Apply(s, gowindow.RifeVincentClass1Order1)
	for i := range s {
		fmt.Printf("%f ", s[i])
	}
	fmt.Println()
	// Output:
	// 2.000000 1.707107 1.000000 0.292893 0.000000 0.292893 1.000000 1.707107 2.000000
}

func ExampleApplyNew_rifeVincentClass1Order1() {
	s := []float64{1, 1, 1, 1, 1, 1, 1, 1, 1}
	for _, x := range gowindow.ApplyNew(s, gowindow.RifeVincentClass1Order1) {
		fmt.Printf("%f ", x)
	}
	fmt.Println()
	// Output:
	// 2.000000 1.707107 1.000000 0.292893 0.000000 0.292893 1.000000 1.707107 2.000000
}

func ExampleApply_rifeVincentClass1Order2() {
	s := []float64{1, 1, 1, 1, 1, 1, 1, 1, 1}
	gowindow.Apply(s, gowindow.RifeVincentClass1Order2)
	for i := range s {
		fmt.Printf("%f ", s[i])
	}
	fmt.Println()
	// Output:
	// 0.000000 0.057191 0.666667 1.942809 2.666667 1.942809 0.666667 0.057191 0.000000
}

func ExampleApplyNew_rifeVincentClass1Order2() {
	s := []float64{1, 1, 1, 1, 1, 1, 1, 1, 1}
	for _, x := range gowindow.ApplyNew(s, gowindow.RifeVincentClass1Order2) {
		fmt.Printf("%f ", x)
	}
	fmt.Println()
	// Output:
	// 0.000000 0.057191 0.666667 1.942809 2.666667 1.942809 0.666667 0.057191 0.000000
}

func ExampleApply_rifeVincentClass1Order3() {
	s := []float64{1, 1, 1, 1, 1, 1, 1, 1, 1}
	gowindow.Apply(s, gowindow.RifeVincentClass1Order3)
	for i := range s {
		fmt.Printf("%f ", s[i])
	}
	fmt.Println()
	// Output:
	// -0.000000 0.010051 0.400000 1.989949 3.200000 1.989949 0.400000 0.010051 -0.000000
}

func ExampleApplyNew_rifeVincentClass1Order3() {
	s := []float64{1, 1, 1, 1, 1, 1, 1, 1, 1}
	for _, x := range gowindow.ApplyNew(s, gowindow.RifeVincentClass1Order3) {
		fmt.Printf("%f ", x)
	}
	fmt.Println()
	// Output:
	// -0.000000 0.010051 0.400000 1.989949 3.200000 1.989949 0.400000 0.010051 -0.000000
}

func ExampleApply_rifeVincentClass1Order4() {
	s := []float64{1, 1, 1, 1, 1, 1, 1, 1, 1}
	gowindow.Apply(s, gowindow.RifeVincentClass1Order4)
	for i := range s {
		fmt.Printf("%f ", s[i])
	}
	fmt.Println()
	// Output:
	// -0.000000 0.001682 0.228571 1.941175 3.657143 1.941175 0.228571 0.001682 -0.000000
}

func ExampleApplyNew_rifeVincentClass1Order4() {
	s := []float64{1, 1, 1, 1, 1, 1, 1, 1, 1}
	for _, x := range gowindow.ApplyNew(s, gowindow.RifeVincentClass1Order4) {
		fmt.Printf("%f ", x)
	}
	fmt.Println()
	// Output:
	// -0.000000 0.001682 0.228571 1.941175 3.657143 1.941175 0.228571 0.001682 -0.000000
}

func ExampleApply_rifeVincentClass2Decibel36() {
	s := []float64{1, 1, 1, 1, 1, 1, 1, 1, 1}
	gowindow.Apply(s, gowindow.RifeVincentClass2Decibel36)
	for i := range s {
		fmt.Printf("%f ", s[i])
	}
	fmt.Println()
	// Output:
	// -2024927293350551435766351869747885244416.000000 2025111945874451055459141110433658699776.000000 -2025558669492556963720163972058682753024.000000 2026006722622037184443210482864703406080.000000 -2026192704703477575797466831655152910336.000000 2026006722622037184443210482864703406080.000000 -2025558669492556963720163972058682753024.000000 2025111945874451055459141110433658699776.000000 -2024927293350551435766351869747885244416.000000
}

func ExampleApplyNew_rifeVincentClass2Decibel36() {
	s := []float64{1, 1, 1, 1, 1, 1, 1, 1, 1}
	for _, x := range gowindow.ApplyNew(s, gowindow.RifeVincentClass2Decibel36) {
		fmt.Printf("%f ", x)
	}
	fmt.Println()
	// Output:
	// -2024927293350551435766351869747885244416.000000 2025111945874451055459141110433658699776.000000 -2025558669492556963720163972058682753024.000000 2026006722622037184443210482864703406080.000000 -2026192704703477575797466831655152910336.000000 2026006722622037184443210482864703406080.000000 -2025558669492556963720163972058682753024.000000 2025111945874451055459141110433658699776.000000 -2024927293350551435766351869747885244416.000000
}

func ExampleApply_rifeVincentClass2Decibel42() {
	s := []float64{1, 1, 1, 1, 1, 1, 1, 1, 1}
	gowindow.Apply(s, gowindow.RifeVincentClass2Decibel42)
	for i := range s {
		fmt.Printf("%f ", s[i])
	}
	fmt.Println()
	// Output:
	// -20034352315543982528183444113394012813210419200.000000 3226210894818413575430568568275426219130880.000000 20038912815454252610183833247970999120403890176.000000 -3233214445112811968395857129592891669217280.000000 -20043487322465078530054547025706181444214194176.000000 -3233214445034238057192355115474631906033664.000000 20038912815454252610183833247970999120403890176.000000 3226210894739752389454296934847786067689472.000000 -20034352315543982528183444113394012813210419200.000000
}

func ExampleApplyNew_rifeVincentClass2Decibel42() {
	s := []float64{1, 1, 1, 1, 1, 1, 1, 1, 1}
	for _, x := range gowindow.ApplyNew(s, gowindow.RifeVincentClass2Decibel42) {
		fmt.Printf("%f ", x)
	}
	fmt.Println()
	// Output:
	// -20034352315543982528183444113394012813210419200.000000 3226210894818413575430568568275426219130880.000000 20038912815454252610183833247970999120403890176.000000 -3233214445112811968395857129592891669217280.000000 -20043487322465078530054547025706181444214194176.000000 -3233214445034238057192355115474631906033664.000000 20038912815454252610183833247970999120403890176.000000 3226210894739752389454296934847786067689472.000000 -20034352315543982528183444113394012813210419200.000000
}

func ExampleApply_rifeVincentClass2Decibel48() {
	s := []float64{1, 1, 1, 1, 1, 1, 1, 1, 1}
	gowindow.Apply(s, gowindow.RifeVincentClass2Decibel48)
	for i := range s {
		fmt.Printf("%f ", s[i])
	}
	fmt.Println()
	// Output:
	// -218284551938746400404405163437822788418196898771369984.000000 -218295625748461176272622579341338437390956113938087936.000000 -218322391419704292773262831043187517146693365426814976.000000 -218349201301133075965126737968320816910806802084397056.000000 -218360319321504809602753308867017459723235201392836608.000000 -218349201301133075965126737968320816910806802084397056.000000 -218322391419704292773262831043187517146693365426814976.000000 -218295625748461176272622579341338437390956113938087936.000000 -218284551938746400404405163437822788418196898771369984.000000
}

func ExampleApplyNew_rifeVincentClass2Decibel48() {
	s := []float64{1, 1, 1, 1, 1, 1, 1, 1, 1}
	for _, x := range gowindow.ApplyNew(s, gowindow.RifeVincentClass2Decibel48) {
		fmt.Printf("%f ", x)
	}
	fmt.Println()
	// Output:
	// -218284551938746400404405163437822788418196898771369984.000000 -218295625748461176272622579341338437390956113938087936.000000 -218322391419704292773262831043187517146693365426814976.000000 -218349201301133075965126737968320816910806802084397056.000000 -218360319321504809602753308867017459723235201392836608.000000 -218349201301133075965126737968320816910806802084397056.000000 -218322391419704292773262831043187517146693365426814976.000000 -218295625748461176272622579341338437390956113938087936.000000 -218284551938746400404405163437822788418196898771369984.000000
}

func ExampleApply_rifeVincentClass2Decibel54() {
	s := []float64{1, 1, 1, 1, 1, 1, 1, 1, 1}
	gowindow.Apply(s, gowindow.RifeVincentClass2Decibel54)
	for i := range s {
		fmt.Printf("%f ", s[i])
	}
	fmt.Println()
	// Output:
	// -2538938859796059347073265806690815032417726293242690760867840.000000 -244839639316212074974845845766681938292415381872066953216.000000 2539285022241429772119582693954874784717465082041068572114944.000000 245157479264894576579949226016309220724698754134742925312.000000 -2539631820366732544198220548668720662927887267635235026632704.000000 245157479334566460194455480722799135803427362112532381696.000000 2539285022241429772119582693954874784717465082041068572114944.000000 -244839639313710210122885691482990899898702643815614775296.000000 -2538938859796059347073265806690815032417726293242690760867840.000000
}

func ExampleApplyNew_rifeVincentClass2Decibel54() {
	s := []float64{1, 1, 1, 1, 1, 1, 1, 1, 1}
	for _, x := range gowindow.ApplyNew(s, gowindow.RifeVincentClass2Decibel54) {
		fmt.Printf("%f ", x)
	}
	fmt.Println()
	// Output:
	// -2538938859796059347073265806690815032417726293242690760867840.000000 -244839639316212074974845845766681938292415381872066953216.000000 2539285022241429772119582693954874784717465082041068572114944.000000 245157479264894576579949226016309220724698754134742925312.000000 -2539631820366732544198220548668720662927887267635235026632704.000000 245157479334566460194455480722799135803427362112532381696.000000 2539285022241429772119582693954874784717465082041068572114944.000000 -244839639313710210122885691482990899898702643815614775296.000000 -2538938859796059347073265806690815032417726293242690760867840.000000
}

func ExampleApply_rifeVincentClass2Decibel60() {
	s := []float64{1, 1, 1, 1, 1, 1, 1, 1, 1}
	gowindow.Apply(s, gowindow.RifeVincentClass2Decibel60)
	for i := range s {
		fmt.Printf("%f ", s[i])
	}
	fmt.Println()
	// Output:
	// -30970915990566448473006609421802001995931304479890506990725082644480.000000 30971913315417809613206233183367151682047345342664872460128800997376.000000 -30974322849164585943352899377737985251114691744970032063928478990336.000000 30976734904916304402311990961700205082917831408565629490548592082944.000000 -30977734751783365071663633871757014272425091726392006510126510899200.000000 30976734904916304402311990961700205082917831408565629490548592082944.000000 -30974322849164585943352899377737985251114691744970032063928478990336.000000 30971913315417809613206233183367151682047345342664872460128800997376.000000 -30970915990566448473006609421802001995931304479890506990725082644480.000000
}

func ExampleApplyNew_rifeVincentClass2Decibel60() {
	s := []float64{1, 1, 1, 1, 1, 1, 1, 1, 1}
	for _, x := range gowindow.ApplyNew(s, gowindow.RifeVincentClass2Decibel60) {
		fmt.Printf("%f ", x)
	}
	fmt.Println()
	// Output:
	// -30970915990566448473006609421802001995931304479890506990725082644480.000000 30971913315417809613206233183367151682047345342664872460128800997376.000000 -30974322849164585943352899377737985251114691744970032063928478990336.000000 30976734904916304402311990961700205082917831408565629490548592082944.000000 -30977734751783365071663633871757014272425091726392006510126510899200.000000 30976734904916304402311990961700205082917831408565629490548592082944.000000 -30974322849164585943352899377737985251114691744970032063928478990336.000000 30971913315417809613206233183367151682047345342664872460128800997376.000000 -30970915990566448473006609421802001995931304479890506990725082644480.000000
}

func ExampleApply_rifeVincentClass2Decibel66() {
	s := []float64{1, 1, 1, 1, 1, 1, 1, 1, 1}
	gowindow.Apply(s, gowindow.RifeVincentClass2Decibel66)
	for i := range s {
		fmt.Printf("%f ", s[i])
	}
	fmt.Println()
	// Output:
	// -349387152606032305735995979975518645532438057363959586581927773160651882496.000000 22376263044339686063281079105408736797582696830253109644307485113188352.000000 349418791775967510635733104067133904676500643113242499047974527563944951808.000000 -22395535385910030154668113695989759737237169505096498111121875877756928.000000 -349450469490581205773984512236236305960887828565638856209964297276947955712.000000 -22395535396524875843431101532297434004264057568523530828317818185842688.000000 349418791775967510635733104067133904676500643113242499047974527563944951808.000000 22376263042539653580872986833713622979820595270594945521129558403710976.000000 -349387152606032305735995979975518645532438057363959586581927773160651882496.000000
}

func ExampleApplyNew_rifeVincentClass2Decibel66() {
	s := []float64{1, 1, 1, 1, 1, 1, 1, 1, 1}
	for _, x := range gowindow.ApplyNew(s, gowindow.RifeVincentClass2Decibel66) {
		fmt.Printf("%f ", x)
	}
	fmt.Println()
	// Output:
	// -349387152606032305735995979975518645532438057363959586581927773160651882496.000000 22376263044339686063281079105408736797582696830253109644307485113188352.000000 349418791775967510635733104067133904676500643113242499047974527563944951808.000000 -22395535385910030154668113695989759737237169505096498111121875877756928.000000 -349450469490581205773984512236236305960887828565638856209964297276947955712.000000 -22395535396524875843431101532297434004264057568523530828317818185842688.000000 349418791775967510635733104067133904676500643113242499047974527563944951808.000000 22376263042539653580872986833713622979820595270594945521129558403710976.000000 -349387152606032305735995979975518645532438057363959586581927773160651882496.000000
}

func ExampleApply_rifeVincentClass3Order2() {
	s := []float64{1, 1, 1, 1, 1, 1, 1, 1, 1}
	gowindow.Apply(s, gowindow.RifeVincentClass3Order2)
	for i := range s {
		fmt.Printf("%f ", s[i])
	}
	fmt.Println()
	// Output:
	// 0.000000 0.153699 0.803150 1.846301 2.393700 1.846301 0.803150 0.153699 0.000000
}

func ExampleApplyNew_rifeVincentClass3Order2() {
	s := []float64{1, 1, 1, 1, 1, 1, 1, 1, 1}
	for _, x := range gowindow.ApplyNew(s, gowindow.RifeVincentClass3Order2) {
		fmt.Printf("%f ", x)
	}
	fmt.Println()
	// Output:
	// 0.000000 0.153699 0.803150 1.846301 2.393700 1.846301 0.803150 0.153699 0.000000
}

func ExampleApply_rifeVincentClass3Order3() {
	s := []float64{1, 1, 1, 1, 1, 1, 1, 1, 1}
	gowindow.Apply(s, gowindow.RifeVincentClass3Order3)
	for i := range s {
		fmt.Printf("%f ", s[i])
	}
	fmt.Println()
	// Output:
	// 0.000001 0.028164 0.502463 1.971836 2.995073 1.971836 0.502463 0.028164 0.000001
}

func ExampleApplyNew_rifeVincentClass3Order3() {
	s := []float64{1, 1, 1, 1, 1, 1, 1, 1, 1}
	for _, x := range gowindow.ApplyNew(s, gowindow.RifeVincentClass3Order3) {
		fmt.Printf("%f ", x)
	}
	fmt.Println()
	// Output:
	// 0.000001 0.028164 0.502463 1.971836 2.995073 1.971836 0.502463 0.028164 0.000001
}

func ExampleApply_rifeVincentClass3Order4() {
	s := []float64{1, 1, 1, 1, 1, 1, 1, 1, 1}
	gowindow.Apply(s, gowindow.RifeVincentClass3Order4)
	for i := range s {
		fmt.Printf("%f ", s[i])
	}
	fmt.Println()
	// Output:
	// -0.003546 0.002294 0.292473 1.961864 3.490284 1.961864 0.292473 0.002294 -0.003546
}

func ExampleApplyNew_rifeVincentClass3Order4() {
	s := []float64{1, 1, 1, 1, 1, 1, 1, 1, 1}
	for _, x := range gowindow.ApplyNew(s, gowindow.RifeVincentClass3Order4) {
		fmt.Printf("%f ", x)
	}
	fmt.Println()
	// Output:
	// -0.003546 0.002294 0.292473 1.961864 3.490284 1.961864 0.292473 0.002294 -0.003546
}

func ExampleApply_missedSwitchImplementation() {
	s := []float64{1, 2, 3, 4}
	gowindow.Apply(s, gowindow.None)
	for i := range s {
		fmt.Printf("%f ", s[i])
	}
	fmt.Println()
	// Output:
	// 1.000000 2.000000 3.000000 4.000000
}

func ExampleApplyNew_missedSwitchImplementation() {
	for _, x := range gowindow.ApplyNew([]float64{1, 2, 3, 4}, gowindow.None) {
		fmt.Printf("%f ", x)
	}
	fmt.Println()
	// Output:
	//
}
