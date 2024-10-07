package adif

import (
	"errors"
	"reflect"
	"strings"
)

const (
	newLineStr   = "\n"
	chevronRight = ">"
	chevronLeft  = "<"
	colonStr     = ":"
	eohStr       = "<eoh>"
)

// UnmarshalADI parses the provided ADI data and populates the given Record.
// data: the ADI data in byte slice format
// r: pointer to the Record struct to be populated
// Returns an error if the input Record pointer is nil or if there are issues during unmarshalling
// NOTES:
// As a Record object, use json.Marshal to convert the Record object to JSON format and
// json.Unmarshal to convert the JSON object back to a Record object or another struct with the same tags.
func UnmarshalADI(data []byte, r *Record) error {
	if r == nil {
		return errors.New("nil pointer passed to Unmarshal")
	}
	if r.QsoSlice == nil {
		r.QsoSlice = make(QsoSlice, 0)
	}
	var qso = newEmptyQso()

	lines := strings.Split(string(data), newLineStr)

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if strings.ToLower(line) == eohStr {
			continue
		}
		if strings.ToLower(line) == eorStr {
			r.QsoSlice = append(r.QsoSlice, qso)
			qso = newEmptyQso()
		}
		if strings.HasPrefix(line, chevronLeft) && strings.Contains(line, colonStr) && strings.Contains(line, chevronRight) {
			parts := strings.SplitN(line, colonStr, 2)
			key := strings.ToLower(strings.Trim(parts[0], chevronLeft))
			sub := strings.SplitN(parts[1], chevronRight, 2)
			value := strings.TrimSpace(sub[1])

			fieldName, found := findJSONTagByName(r, key)
			if !found {
				fieldName, found = findJSONTagByName(qso, key)
			}
			if !found {
				fieldName, found = findJSONTagByName(qso.ContactedStation, key)
			}
			if !found {
				fieldName, found = findJSONTagByName(qso.LoggingStation, key)
			}
			if found {
				field := reflect.ValueOf(r).Elem().FieldByName(fieldName)
				if !field.IsValid() || !field.CanSet() {
					field = reflect.ValueOf(qso).Elem().FieldByName(fieldName)
				}
				if !field.IsValid() || !field.CanSet() {
					field = reflect.ValueOf(qso.ContactedStation).Elem().FieldByName(fieldName)
				}
				if !field.IsValid() || !field.CanSet() {
					field = reflect.ValueOf(qso.LoggingStation).Elem().FieldByName(fieldName)
				}
				if field.IsValid() && field.CanSet() {
					field.SetString(value)
				}
			}
		}
	}
	return nil
}

func findJSONTagByName(v interface{}, tagName string) (string, bool) {
	val := reflect.ValueOf(v)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}
	if val.Kind() != reflect.Struct {
		return emptyStr, false
	}

	typ := val.Type()
	for i := 0; i < typ.NumField(); i++ {
		field := typ.Field(i)
		tag := field.Tag.Get(jsonStructTag)
		if tag == tagName {
			return field.Name, true
		}
	}
	return emptyStr, false
}

func newEmptyQso() *Qso {
	return &Qso{
		ContactedStation: &ContactedStation{},
		LoggingStation:   &LoggingStation{},
	}
}
