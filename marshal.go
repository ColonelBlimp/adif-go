//go:build windows

package adif

import "encoding/json"

func Marshal(r *Record) ([]byte, error) {
	return json.Marshal(r)
}
