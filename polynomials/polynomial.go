package polynomial


// Polynomial does some thing.
type Polynomial struct {
	base uint32
	coefficients []int64
	degree uint64
}

// NewPolynomial creates a new polynomial
func NewPolynomial(degree uint64) (poly *Polynomial, err error) {
	poly = &Polynomial{}
	poly.coefficients = make([]int64, degree)
	poly.degree = degree
	return poly, err
}

// AddCoefficient adds a coefficient to the array.
func (poly *Polynomial) AddCoefficient(coefficient int64, index uint64) {
	poly.coefficients[index] += coefficient
}