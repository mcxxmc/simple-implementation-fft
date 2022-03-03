package test

import (
	"fmt"
	"github.com/mcxxmc/simple-implementation-fft/fft"
	"testing"
)

func TestFFT(t *testing.T) {
	a := fft.Convert2Complex([]float64{0.0, 1.0, 2.0, 3.0})
	b := fft.Fft(a)
	c := fft.FftV2(a)
	d := fft.BruteForceFFT(a)
	b = fft.RoundUpComplexNums(b)
	c = fft.RoundUpComplexNums(c)
	d = fft.RoundUpComplexNums(d)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)

	a = fft.Convert2Complex([]float64{1.0, 2.0, 3.0, 4.0})
	b = fft.Fft(a)
	c = fft.FftV2(a)
	d = fft.BruteForceFFT(a)
	b = fft.RoundUpComplexNums(b)
	c = fft.RoundUpComplexNums(c)
	d = fft.RoundUpComplexNums(d)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)

	a = fft.Convert2Complex([]float64{5.0, 1.0, 2.0, 3.0, 4.0, 6.0, 7.0, 8.0})
	b = fft.Fft(a)
	c = fft.FftV2(a)
	d = fft.BruteForceFFT(a)
	b = fft.RoundUpComplexNums(b)
	c = fft.RoundUpComplexNums(c)
	d = fft.RoundUpComplexNums(d)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
}

func TestBitReverse(t *testing.T) {
	a := fft.Convert2Complex([]float64{0, 4, 2, 6, 1, 5, 3, 7})
	b := fft.BitReverse(a)
	fmt.Println(b)
}
