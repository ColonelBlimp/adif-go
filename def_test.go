//go:build windows

package adif

import (
	"encoding/json"
	"testing"
)

func TestRecord(t *testing.T) {
	rec := New()

	data, err := json.Marshal(rec)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	t.Log(string(data))
}
