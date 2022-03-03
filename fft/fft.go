package fft

import (
	"math"
	"math/cmplx"
)

// Fft returns FFT of a using recursive functions.
//
// assumes length of a to be a power of 2.
func Fft(a []complex128) []complex128 {
	n := len(a)
	if n == 1 {
		// the base case
		return a
	}
	// by euler's identity, e ^ ix = cos x + i sin x, where i is the complex symbol
	// mu = e ^ 2 pi i / n, where x = 2 pi / n
	mu := complex(math.Cos(2 * math.Pi / float64(n)), math.Sin(2 * math.Pi / float64(n)))
	m := 1.0 + 0.0i
	a0 := make([]complex128, n / 2)
	a1 := make([]complex128, n - (n / 2))
	index0, index1 := 0, 0
	for i, v := range a {
		if i % 2 == 0 {
			a0[index0] = v
			index0 ++
		} else {
			a1[index1] = v
			index1 ++
		}
	}
	y0 := Fft(a0)
	y1 := Fft(a1)

	y := make([]complex128, n)
	for k := 0; k < n / 2; k ++ {
		t := m * y1[k]
		y[k] = y0[k] + t
		y[k + (n / 2)] = y0[k] - t
		m = m * mu
	}
	return y
}

// FftV2 returns FFT of a using a loop.
//
// assumes length of a to be a power of 2.
func FftV2(a []complex128) []complex128 {
	n := len(a)
	y := BitReverse(a)
	for s := 1; s <= int(math.Log2(float64(n))); s ++ {
		m := int(math.Pow(float64(2), float64(s)))
		mu := complex(math.Cos(2 * math.Pi / float64(m)), math.Sin(2 * math.Pi / float64(m)))
		for k := 0; k <= n - 1; k += m {
			mul := 1.0 + 0.0i
			for j := 0; j <= m / 2 - 1; j ++ {
				t := mul * y[k + j + m / 2]
				u := y[k + j]
				y[k + j] = u + t
				y[k + j + m / 2] = u - t
				mul = mul * mu
			}
		}
	}
	return y
}

// BitReverse used by FftV2(); rearranges the sequences of elements in the slice to speed up later calculation.
func BitReverse(a []complex128) []complex128 {
	n := len(a)
	lgn := int(math.Log2(float64(n)))
	var rev func(k int) int		// reverse k as a new int with a length of lgn-bit
	rev = func(k int) int {
		r := 0
		step := 1
		for k > 0 {
			res := k % 2
			k = k / 2
			r = r + (res << (lgn - step))
			step ++
		}
		return r
	}
	r := make([]complex128, n)
	for k := 0; k < n; k ++ {
		r[rev(k)] = a[k]
	}
	return r
}

// BruteForceFFT returns the FFT of a using brute force.
//
// assumes length of a to be a power of 2.
//
// WARNING: this has a time complexity of O(n ^ 2) and is only used for testing.
func BruteForceFFT(a []complex128) []complex128 {
	n := len(a)
	y := make([]complex128, n)
	mu := complex(math.Cos(2 * math.Pi / float64(n)), math.Sin(2 * math.Pi / float64(n)))
	for k, _ := range y {
		muK := cmplx.Pow(mu, complex(float64(k), 0))
		for j := 0; j < n; j ++ {
			y[k] += a[j] * cmplx.Pow(muK, complex(float64(j), 0))
		}
	}
	return y
}

// Convert2Complex converts a float64 array to a complex128 array.
func Convert2Complex(a []float64) (b []complex128) {
	b = make([]complex128, len(a))
	for i, n := range a {
		b[i] = complex(n, 0)
	}
	return
}

// RoundUpComplexNums rounds up all the real and imaginary numbers in the slice of complex numbers.
func RoundUpComplexNums(a []complex128) (b []complex128) {
	b = make([]complex128, len(a))
	for i, n := range a {
		b[i] = complex(math.Round(real(n)), math.Round(imag(n)))
	}
	return
}
