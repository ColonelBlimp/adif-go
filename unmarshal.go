package adif

import (
	"errors"
	"reflect"
	"strings"
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

	lines := strings.Split(string(data), "\n")

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if strings.ToLower(line) == "<eoh>" {
			continue
		}
		if strings.ToLower(line) == "<eor>" {
			r.QsoSlice = append(r.QsoSlice, qso)
			qso = new(Qso)
			qso.ContactedStation = new(ContactedStation)
			qso.LoggingStation = new(LoggingStation)
		}
		if strings.HasPrefix(line, "<") && strings.Contains(line, ":") && strings.Contains(line, ">") {
			parts := strings.SplitN(line, ":", 2)
			key := strings.ToLower(strings.Trim(parts[0], "<"))
			sub := strings.SplitN(parts[1], ">", 2)
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
		return "?1", false
	}

	typ := val.Type()
	for i := 0; i < typ.NumField(); i++ {
		field := typ.Field(i)
		tag := field.Tag.Get("json")
		//		fmt.Printf("Tag: %s looking: %s\n", tag, tagName)
		if tag == tagName {
			return field.Name, true
		}
	}
	return "?2", false
}
