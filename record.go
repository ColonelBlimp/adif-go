package adif

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"reflect"
	"strings"
)

const (
	adifFormat    = "<%s:%d>%s"
	errMsgTag     = "errormsg"
	jsonStructTag = "json"
	emptyStr      = ""
	dotStr        = "."
)

// NewRecord creates a new ADI Record object
func NewRecord(qso *Qso) (*Record, error) {
	if qso == nil {
		return nil, ErrorNilQso
	}

	validate := validator.New()
	if err := registerValidators(validate); err != nil {
		return nil, err
	}

	return &Record{
		validate: validate,
		QSO:      qso,
	}, nil
}

func (r *Record) SetHeader(ptr *Header) error {
	if ptr == nil {
		return ErrorNilHeader
	}
	r.HEADER = ptr
	return nil
}

func (r *Record) SetQsl(ptr *Qsl) error {
	if ptr == nil {
		return ErrorNilQsl
	}
	r.QSL = ptr
	return nil
}

// Validate validates the Record object. The fields which are checked are:
// Qso.Band, Qso.Freq, Qso.Mode, Qso.QsoDate, Qso.RstRcvd, Qso.RstSent, Qso.TimeOn
// ContactedStation.Call, LoggingStation.Name, LoggingStation.StationCallsign
// If any of the fields are invalid, the function returns an error with an appropriate error message.
func (r *Record) Validate() error {
	return ValidateFunc[Record](*r, r.validate)
}

func (r *Record) ADIString() string {
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

	parseStruct(reflect.ValueOf(r).Elem())
	return result + "<EOR>"
}

func formatField(tagName string, value string) string {
	return fmt.Sprintf(adifFormat, strings.ToUpper(tagName), len(value), value)
}
