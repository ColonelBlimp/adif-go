//go:build windows

package adif

import "testing"

func TestNewQso(t *testing.T) {
	qso, err := NewQso("20m", "14.200", "USB", "20210606", "1200", "59", "59")
	if err != nil {
		t.Fatal(err)
	}
	if qso.Band != "20m" {
		t.Log("Band not set correctly")
		t.FailNow()
	}
	if qso.Freq != "14.200" {
		t.Log("Frequency not set correctly")
		t.FailNow()
	}
	if qso.Mode != "USB" {
		t.Log("Mode not set correctly")
		t.FailNow()
	}
	if qso.QsoDate != "20210606" {
		t.Log("QSO date not set correctly")
		t.FailNow()
	}
	if qso.TimeOn != "1200" {
		t.Log("Time on not set correctly")
		t.FailNow()
	}
	if qso.RstRcvd != "59" {
		t.Log("RST received not set correctly")
		t.FailNow()
	}
	if qso.RstSent != "59" {
		t.Log("RST sent not set correctly")
		t.FailNow()
	}
	if err = qso.Validate(); err == nil {
		t.Log("Expected a validation error")
		t.FailNow()
	}

	cs, err := NewContactedStation("XX1XX", "Test")
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	if err = qso.SetContactedStation(cs); err != nil {
		t.Error(err)
		t.FailNow()
	}

	ls, err := NewLoggingStation("YY2YY")
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	if err = qso.SetLoggingStation(ls); err != nil {
		t.Error(err)
		t.FailNow()
	}
	if err = qso.SetQsl(NewQsl()); err != nil {
		t.Error(err)
		t.FailNow()
	}
	if err = qso.Validate(); err != nil {
		t.Log(err)
		t.FailNow()
	}

	record, err := NewRecord("3.1.4", "YYYYMMDD HHMMSS", "adif-go", "0.1.0")
	if err != nil {
		t.Fatal(err)
	}
	if err = record.AddQso(qso); err != nil {
		t.Fatal(err)
	}
	if err = record.Validate(); err != nil {
		t.Log(err)
		t.FailNow()
	}

	t.Log(record.ADIHeaderString())
	t.Log(record.QsoSlice[0].ADIQsoString())
}
