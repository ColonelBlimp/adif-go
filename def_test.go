//go:build windows

package adif

import (
	"testing"
)

func TestRecord(t *testing.T) {
	contactedStation := NewContactedStation("XX1XXX")
	loggingStation := NewLoggingStation("Y1YY", "My Name")

	qso := NewQso("15m", "21.250", "USB", "20240929", "1621", "59", "59")
	qso.SetLoggingStation(loggingStation)
	qso.SetContactedStation(contactedStation)

	rec, err := NewRecord(qso)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	if err = rec.Validate(); err != nil {
		t.Error(err)
		t.FailNow()
	}
}
