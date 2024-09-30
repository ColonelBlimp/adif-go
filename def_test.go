package adif

import (
	"testing"
)

func TestRecord(t *testing.T) {
	contactedStation := NewContactedStation("XX1XXX")
	loggingStation := NewLoggingStation("Y1YY", "My Name")

	qso, err := NewQso("15m", "21.250", "USB", "20240929", "1621", "59", "59")
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	if err = qso.Validate(); err == nil {
		t.Log("Expected an error as the ContactedStation is not set")
		t.FailNow()
	}

	if err = qso.SetContactedStation(contactedStation); err != nil {
		t.Error(err)
		t.FailNow()
	}
	if err = qso.Validate(); err == nil {
		t.Log("Expected an error as the LoggingStation is not set")
		t.FailNow()
	}

	if err = qso.SetLoggingStation(loggingStation); err != nil {
		t.Error(err)
		t.FailNow()
	}

	if err = qso.Validate(); err != nil {
		t.Error(err)
		t.FailNow()
	}

	var rec *Record
	rec, err = NewRecord(qso)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	if err = rec.SetQsl(NewQsl()); err != nil {
		t.Error(err)
		t.FailNow()
	}

	if err = rec.Validate(); err != nil {
		t.Error(err)
		t.FailNow()
	}

	t.Log(rec.ADIString())
}
