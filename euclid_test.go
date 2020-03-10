package euclid

import "testing"

func TestZero(t *testing.T) {
	d, err := GCD(0, 1)
	if err != nil || d != 1 {
		t.Fail()
	}
	d, err = GCD(1, 0)
	if err != nil || d != 1 {
		t.Fail()
	}
}

//func Test(t *testing.T) {
