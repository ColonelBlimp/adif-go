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

func TestFortyMeterBand(t *testing.T) {
	model := test{
		Qso: Qso{
			Band:    "40m",
			Freq:    "7.120",
			Mode:    "SSB",
			QsoDate: "20210101",
			RstRcvd: "59",
			RstSent: "59",
			TimeOn:  "1200",
		},
	}

	if err := Validate[test](model); err != nil {
		t.Logf("Expected test to pass: %v", err)
		t.FailNow()
	}
}
