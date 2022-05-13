package util

import "testing"

func Test_IF(t *testing.T) {
	r := IF(1 > 2, 1, 2).(int)
	if r != 2 {
		t.Errorf("iff error")
	}
}

func Test_Any(t *testing.T) {
	any := Any(true, false, true)
	if any != true {
		t.Errorf("Any error")
	}
	any = Any()
	if any != false {
		t.Errorf("Any error")
	}
	any = Any(false, false, false)
	if any != false {
		t.Errorf("Any error")
	}
}

func Test_All(t *testing.T) {
	all := All(true, true, true)
	if all != true {
		t.Errorf("All error ")
	}
	all = All()
	if all != true {
		t.Errorf("All error")
	}
	all = All(true, false, false)
	if all != false {
		t.Errorf("All error")
	}
}
