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

func Unmarshal(data []byte, r *Record) error {
	if r == nil {
		return errors.New("nil pointer passed to Unmarshal")
	}
	if r.QsoSlice == nil {
		r.QsoSlice = make(QsoSlice, 0)
	}
	var qso = new(Qso)
	qso.ContactedStation = new(ContactedStation)
	qso.LoggingStation = new(LoggingStation)

	lines := strings.Split(string(data), newLineStr)

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if strings.ToLower(line) == eohStr {
			continue
		}
		if strings.ToLower(line) == eorStr {
			r.QsoSlice = append(r.QsoSlice, qso)
			qso = new(Qso)
			qso.ContactedStation = new(ContactedStation)
			qso.LoggingStation = new(LoggingStation)
		}
		if strings.HasPrefix(line, chevronLeft) && strings.Contains(line, colonStr) && strings.Contains(line, chevronRight) {
			parts := strings.SplitN(line, colonStr, 2)
			key := strings.ToLower(strings.Trim(parts[0], chevronLeft))
			sub := strings.SplitN(parts[1], chevronRight, 2)
			value := strings.TrimSpace(sub[1])

			fieldName, found := findJSONTagByName(r, key)
			if found {
				field := reflect.ValueOf(r).Elem().FieldByName(fieldName)
				if field.IsValid() && field.CanSet() {
					field.SetString(value)
				}
			} else {
				fieldName, found = findJSONTagByName(qso, key)
				if found {
					field := reflect.ValueOf(qso).Elem().FieldByName(fieldName)
					if field.IsValid() && field.CanSet() {
						field.SetString(value)
					}
				} else {
					fieldName, found = findJSONTagByName(qso.ContactedStation, key)
					if found {
						field := reflect.ValueOf(qso.ContactedStation).Elem().FieldByName(fieldName)
						if field.IsValid() && field.CanSet() {
							field.SetString(value)
						}
					} else {
						fieldName, found = findJSONTagByName(qso.LoggingStation, key)
						if found {
							field := reflect.ValueOf(qso.LoggingStation).Elem().FieldByName(fieldName)
							if field.IsValid() && field.CanSet() {
								field.SetString(value)
							}
						}
					}
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
