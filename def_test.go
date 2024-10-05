package adif

import (
	"testing"
)

func TestRecord(t *testing.T) {
	qso := createQSOObject(t)

	t.Log(qso.ADIString())

	rec, err := NewRecord(qso)
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

func createQSOObject(t *testing.T) *Qso {
	contactedStation, err := NewContactedStation("XX1XXX", "Their Name")
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	loggingStation, err := NewLoggingStation("Y1YY")
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

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

	return qso
}
