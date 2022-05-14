package util

import "testing"

func Test_SumInt(t *testing.T) {
	a := 1 + 2 + 3
	if a != SumInt([]int{1, 2, 3}...) {
		t.Errorf("SumInt error")
	}
}

func Test_FloatEqual(t *testing.T) {
	if !FloatEqual(6, 6, 0.000001) {
		t.Errorf("FloatEqual error")
	}
}

func Test_SumFloat64(t *testing.T) {
	a := 1 + 2 + 3.0
	if !FloatEqual(a, SumFloat64(1, 2, 3), 0.000001) {
		t.Errorf("SumFloat64 error")
	}
}
