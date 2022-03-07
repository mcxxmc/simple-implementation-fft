package test

import (
	"fmt"
	"github.com/mcxxmc/simple-implementation-fft/fft"
	"testing"
)

func TestFFT(t *testing.T) {
	a := fft.Convert2Complex([]float64{0.0, 1.0, 2.0, 3.0})
	fmt.Println("a: ", a)
	b := fft.Fft(a)
	c := fft.FftV2(a)
	d := fft.BruteForceFFT(a)
	b = fft.RoundUpComplexNums(b)
	c = fft.RoundUpComplexNums(c)
	d = fft.RoundUpComplexNums(d)
	aa := fft.Interpolation(b)
	ar := fft.ExtractReals(aa)
	aa2 := fft.InterpolationV2(b)
	ar2 := fft.ExtractReals(aa2)
	fmt.Println("y: ")
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println("after interpolation, a: ")
	fmt.Println(ar)
	fmt.Println(ar2)

	a = fft.Convert2Complex([]float64{1.0, 2.0, 3.0, 4.0})
	fmt.Println("a: ", a)
	b = fft.Fft(a)
	c = fft.FftV2(a)
	d = fft.BruteForceFFT(a)
	b = fft.RoundUpComplexNums(b)
	c = fft.RoundUpComplexNums(c)
	d = fft.RoundUpComplexNums(d)
	aa = fft.Interpolation(b)
	ar = fft.ExtractReals(aa)
	aa2 = fft.InterpolationV2(b)
	ar2 = fft.ExtractReals(aa2)
	fmt.Println("y: ")
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println("after interpolation, a: ")
	fmt.Println(ar)
	fmt.Println(ar2)

	a = fft.Convert2Complex([]float64{5.0, 1.0, 2.0, 3.0, 4.0, 6.0, 7.0, 8.0})
	fmt.Println("a: ", a)
	b = fft.Fft(a)
	c = fft.FftV2(a)
	d = fft.BruteForceFFT(a)
	b = fft.RoundUpComplexNums(b)
	c = fft.RoundUpComplexNums(c)
	d = fft.RoundUpComplexNums(d)
	aa = fft.Interpolation(b)
	ar = fft.ExtractReals(aa)
	aa2 = fft.InterpolationV2(b)
	ar2 = fft.ExtractReals(aa2)
	fmt.Println("y: ")
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println("after interpolation, a: ")
	fmt.Println(ar)
	fmt.Println(ar2)
}

func TestBitReverse(t *testing.T) {
	a := fft.Convert2Complex([]float64{0, 4, 2, 6, 1, 5, 3, 7})
	b := fft.BitReverse(a)
	fmt.Println(b)
}

// tests the multiplication of 2 polynomials
func TestFFTAdvanced(t *testing.T) {
	a1 := fft.Convert2Complex([]float64{0.0, 0.0, 2.0, 3.0})
	a2 := fft.Convert2Complex([]float64{0.0, 0.0, 3.0, 4.0})
	b1 := fft.FftV2(a1)
	b2 := fft.FftV2(a2)
	b3, err := fft.MultipleYs(b1, b2)
	if err != nil {
		t.Error(err)
		return
	}
	a3 := fft.Interpolation(b3)
	ar := fft.ExtractReals(a3)
	fmt.Println(ar)
}
