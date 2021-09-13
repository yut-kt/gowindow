# gowindow

[![v0.1.3](https://img.shields.io/github/v/release/yut-kt/gowindow?logoColor=ff69b4&style=social)]()
[![Test](https://github.com/yut-kt/gowindow/actions/workflows/default_branch_test.yaml/badge.svg)](https://github.com/yut-kt/gowindow/actions/workflows/default_branch_test.yaml)
[![coverage](https://img.shields.io/badge/coverage-96.4%25-green)]()
[![Go Report Card](https://goreportcard.com/badge/github.com/yut-kt/gowindow)](https://goreportcard.com/report/github.com/yut-kt/gowindow)  
[![Go Reference](https://pkg.go.dev/badge/github.com/yut-kt/gowindow.svg)](https://pkg.go.dev/github.com/yut-kt/gowindow)
[![license](http://img.shields.io/badge/license-MIT-red.svg?style=flat)](https://raw.githubusercontent.com/yut-kt/gowindow/main/LICENSE)


**Window Function support for Go language**

## Install
```bash
$ go get github.com/yut-kt/gowindow
```

## Usage
See [gowindow_test.go](https://github.com/yut-kt/gowindow/blob/main/gowindow_test.go) for detailed Usage

## Supported window function
*[a list of window function](https://en.wikipedia.org/wiki/Window_function#A_list_of_window_functions)*
- Rectangular window
- B-spline windows
  - Triangular window
  - Parzen window
- Other polynomial windows 
  - Welch window
- Sine window
  - Power-of-sine/cosine windows
- Cosine-sum windows
  - Hann and Hamming windows
  - Blackman window
  - Nuttall window, continuous first derivative
  - Blackman–Nuttall window
  - Blackman–Harris window
  - Flat top window
  - Rife–Vincent windows
- Adjustable windows
  - Gaussian window
  - Confined Gaussian window
  - Approximate confined Gaussian window
  - Generalized normal window
  - Tukey window
  - Planck-taper window
  - ~~DPSS or Slepian window~~
    - Difficult to implement
  - Kaiser window
  - Dolph–Chebyshev window
  - Ultraspherical window
  - Exponential or Poisson window
- Hybrid windows
  - Bartlett–Hann window
  - ~~Planck–Bessel window~~
  - ~~Hann–Poisson window~~
- Other windows
  - ~~Generalized adaptive polynomial (GAP) window~~
  - ~~Lanczos window~~

## Benchmark
https://github.com/yut-kt/gowindow/wiki/Benchmark

## License
gowindow is released under the [MIT License](https://raw.githubusercontent.com/yut-kt/gowindow/main/LICENSE).
