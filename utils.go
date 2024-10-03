package adif

import (
	"fmt"
	"reflect"
	"strings"
	"unicode"
)

func isNthRuneFromRightEqual(s string, n int, char rune) bool {
	runes := []rune(s) // Convert the string to a slice of runes
	if n > len(runes) || n <= 0 {
		return false // Out of bounds, return false
	}
	return runes[len(runes)-n] == char
}

func isAllDigits(s string) bool {
	if len(s) == 0 {
		return false // Consider empty string as not all digits
	}
	for _, char := range s {
		if !unicode.IsDigit(char) {
			return false
		}
	}
	return true
}

func parseStructToADIString(obj interface{}) string {
	var result string

	var parseStruct func(v reflect.Value)
	parseStruct = func(v reflect.Value) {
		for i := 0; i < v.NumField(); i++ {
			fieldName := v.Type().Field(i).Name
			if fieldName == "validate" {
				continue
			}

			field := v.Field(i)
			if field.Kind() == reflect.Ptr {
				field = field.Elem()
			}

			if field.Kind() == reflect.Struct {
				parseStruct(field)
				continue
			}

			if field.Kind() == reflect.String && field.String() != emptyStr {
				tag := v.Type().Field(i).Tag.Get(jsonStructTag)
				result += formatField(tag, field.String())
			}
		}
	}

	parseStruct(reflect.ValueOf(obj).Elem())

	return result
}

func formatField(tagName string, value string) string {
	return fmt.Sprintf(adifFormat, strings.ToUpper(tagName), len(value), value)
}
