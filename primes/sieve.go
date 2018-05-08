package primes

import (
	"fmt"
	"math"
)


// Sieve is a prime sieve.
func Sieve(n uint64) {
	A := make([]bool, n)

	var k uint64
	
	// Set all to true initially
	for k = 2; k < n; k++ {
		A[k] = true
	}

	// Limit of how far to look for primes
	limit := uint64(math.Floor(math.Sqrt(float64(n))))
	
	for k = 2; k < limit; k++ {
		if A[k] {
			i := 2*k
			for i < n {
				fmt.Println(i)
				A[i] = false
				i = i + k
			}
		}
	}
	fmt.Println(A)
}