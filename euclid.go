// Package euclid implements the Euclidean algorithm in its simplest form. The implementation is based on "A Computational Introduction to Number Theory and Algebra" version 2, 2008, chapter 4 "Euclid's algorithm" by Victor Shoup [1].
package euclid

import "fmt"

// GCD computes the Greatest Common Divisor of non-negative integers a and b.
// On page 76 of [1] you will find a pseudocode.
func GCD(a, b uint64) (d uint64, err error) {
	if a == 0 && b == 0 {
		return 0, fmt.Errorf("Error: invalid arguments a = 0 and b = 0; GCD(0, 0) is undefined")
	}

	r0 := a
	r1 := b
	for r1 != 0 {
		r2 := r0 % r1
		r0 = r1
		r1 = r2
	}
	d = r0
	return d, nil
}
