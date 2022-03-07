package fft

import (
	"math"
	"math/cmplx"
)

// Interpolation calculates the interpolation, which is the reverse of DFT (or DFT ^ -1).
//
// This function is a modification of the recursive Fft().
func Interpolation(ys []complex128) []complex128 {
	var recursion func(y []complex128) []complex128
	recursion = func(y []complex128) []complex128 {
		n := len(y)
		if n == 1 {
			return []complex128{y[0]}
		}
		tmp := complex(math.Cos(2 * math.Pi / float64(n)), math.Sin(2 * math.Pi / float64(n)))
		mu := cmplx.Pow(tmp, complex(-1, 0))
		m := 1.0 + 0.0i
		y0 := make([]complex128, n / 2)
		y1 := make([]complex128, n - (n / 2))
		index0, index1 := 0, 0
		for i, v := range y {
			if i % 2 == 0 {
				y0[index0] = v
				index0 ++
			} else {
				y1[index1] = v
				index1 ++
			}
		}
		t0 := recursion(y0)
		t1 := recursion(y1)

		a := make([]complex128, n)
		for k := 0; k < n / 2; k ++ {
			t := m * t1[k]
			a[k] = t0[k] + t
			a[k + (n / 2)] = t0[k] - t
			m = m * mu
		}
		return a
	}

	a := recursion(ys)
	n := len(ys)
	nc := complex(float64(n), 0)
	for i := 0; i < n; i ++ {
		a[i] = a[i] / nc
	}
	return a
}

// InterpolationV2 calculates the interpolation, which is the reverse of DFT (or DFT ^ -1).
//
// This function is a modification of the iterative FftV2().
func InterpolationV2(ys []complex128) []complex128 {
	n := len(ys)
	a := BitReverse(ys)
	for s := 1; s <= int(math.Log2(float64(n))); s ++ {
		m := int(math.Pow(float64(2), float64(s)))
		tmp := complex(math.Cos(2 * math.Pi / float64(m)), math.Sin(2 * math.Pi / float64(m)))
		mu := cmplx.Pow(tmp, complex(-1, 0))
		for k := 0; k <= n - 1; k += m {
			mul := 1.0 + 0.0i
			for j := 0; j <= m / 2 - 1; j ++ {
				t := mul * a[k + j + m / 2]
				u := a[k + j]
				a[k + j] = u + t
				a[k + j + m / 2] = u - t
				mul = mul * mu
			}
		}
	}
	nc := complex(float64(n), 0)
	for i := 0; i < n; i ++ {
		a[i] = a[i] / nc
	}
	return a
}

// ExtractReals extract the real parts from []complex128.
func ExtractReals(c []complex128) []float64 {
	r := make([]float64, len(c))
	for i := 0; i < len(c); i ++ {
		r[i] = real(c[i])
	}
	return r
}
