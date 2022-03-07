package fft

import "errors"

// MultipleYs multiplies 2 complex128 slices which should have equal length.
//
// e.g., y0 * y1 = y0[0] * y1[0], y0[1] * y1[1] ...
func MultipleYs(y0, y1 []complex128) ([]complex128, error) {
	if len(y0) != len(y1) {
		return []complex128{}, errors.New("y0 and y1 have different lengths")
	}
	r := make([]complex128, len(y0))
	for i := 0; i < len(y0); i ++ {
		r[i] = y0[i] * y1[i]
	}
	return r, nil
}
