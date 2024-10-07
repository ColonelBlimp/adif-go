//go:build windows

package adif

import "testing"

var data = []byte("")

func TestUnmarshal(t *testing.T) {
	var record Record
	err := Unmarshal(data, &record)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(len(record.QsoSlice))
}
