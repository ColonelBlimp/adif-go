package adif

import (
	"testing"
)

func TestRecord(t *testing.T) {
	record, err := NewRecord("3.1.4", "YYYYMMDD HHMMSS", "adif-go", "0.1.0")
	if err != nil {
		t.Fatal(err)
	}
	if record.ADIFVer != "3.1.4" {
		t.Log("Version not set correctly")
		t.FailNow()
	}
	if record.CreatedTimestamp != "YYYYMMDD HHMMSS" {
		t.Log("Timestamp not set correctly")
		t.FailNow()
	}
	if record.ProgramID != "adif-go" {
		t.Log("Program ID not set correctly")
		t.FailNow()
	}
	if record.ProgramVersion != "0.1.0" {
		t.Log("Program version not set correctly")
		t.FailNow()
	}
	if err = record.Validate(); err != nil {
		t.Log(err)
		t.FailNow()
	}
}
