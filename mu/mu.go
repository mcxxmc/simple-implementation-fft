package mu // Package mu implements the euler's identity which states that e^(ix) = cos(x) + i sin(x)
import "math"

// Mu with N being n, Mu equals to e ^ (2 * pi * i / n), where i is the symbol for complex number.
//
// Please use NewMu() as constructor.
type Mu struct {
	N		int
	exp 	int 		// the exponent
	c		int			// the coefficient
}

// NewMu with N being n, Mu equals to e ^ (2 * pi * i / n), where i is the symbol for complex number.
func NewMu(n int) *Mu {
	return &Mu{N: n, exp: 1, c: 1}
}

// Value returns the value of the Mu object in a complex form.
func (mu *Mu) Value() complex128 {
	x := 2 * math.Pi / float64(mu.N)
	return complex(math.Cos(x),math.Sin(x) )
}
