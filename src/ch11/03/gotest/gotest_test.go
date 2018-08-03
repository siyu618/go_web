package gotest

import "testing"

// run : go test
func TestDivision_1(t *testing.T)  {
	if i, e := Division(6, 2); i != 3 || e != nil {
		t.Error("failed")
	} else {
		t.Log("pass")
	}
}

func TestDivision_2(t *testing.T)  {
	//t.Error("failed")
}

func TestDivision_3(t *testing.T) {
	if _, e := Division(6, 0); e == nil {
		t.Error("failed")
	} else {
		t.Log("passed")
	}
}

