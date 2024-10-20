//go:build windows

package adif

import "testing"

type test struct {
	Qso
}

func TestValidate(t *testing.T) {
	model := test{}

	if err := Validate[test](model); err == nil {
		t.Log("Expected test to fail")
		t.FailNow()
	}
}
