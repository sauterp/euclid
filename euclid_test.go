package euclid

import (
	"testing"
	"fmt"
)

// TestZero tests GCD for correct behaviour with inputs set to zero.
func TestZero(t *testing.T) {
	d, err := GCD(0, 1)
	if err != nil || d != 1 {
		t.Fail()
	}
	d, err = GCD(1, 0)
	if err != nil || d != 1 {
		t.Fail()
	}
	d, err = GCD(0, 0)
	if err == nil {
		t.Fail()
	}
}

// TestPrimes tests GCD for all the combinations of the first ten primes.
func TestPrimes(t *testing.T) {
	primes := []uint64 {
		2,
		3,
		5,
		7,
		11,
		13,
		17,
		19,
		23,
		29,
	}
	for _, p := range primes {
		for _, q := range primes {
			t.Run(fmt.Sprintf("GCD(%d,%d)", p, q), func(t *testing.T) {
				d, err := GCD(p, q)
				if err != nil {
					t.Fail()
				}
				if p != q && d != 1 {
					t.Fail()
				}
				if p == q && d != p {
					t.Fail()
				}
			})
		}
	}
}

