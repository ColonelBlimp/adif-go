//go:build windows

package adif

import (
	"encoding/json"
	"os"
	"testing"
)

//var data = []byte("")

func TestUnmarshal(t *testing.T) {
	data, err := os.ReadFile("testdata/378.adi")
	if err != nil {
		t.Fatal(err)
	}

	var record Record
	err = UnmarshalADI(data, &record)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(len(record.QsoSlice))

	var b []byte
	b, err = json.Marshal(&record)
	if err != nil {
		t.Fatal(err)
	}

	var record2 Record
	err = json.Unmarshal(b, &record2)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(len(record2.QsoSlice))
}
