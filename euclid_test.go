package euclid

import "testing"

func ZeroTest(t *testing.T) {
	d, err := GCD(0, 1)
	if err != nil || d != 0 {
		t.Fail()
	}
	d, err = GCD(1, 0)
	if err != nil || d != 0 {
		t.Fail()
	}
}
