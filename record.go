package adif

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"reflect"
)

const errMsgTag = "errormsg"

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
	val := reflect.ValueOf(r).Elem()
	var result string

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		fieldType := val.Type().Field(i)

		if fieldType.Name == "validate" {
			continue
		}

		if field.Kind() == reflect.Ptr {
			field = field.Elem()
		}

		if field.Kind() == reflect.Struct {
			for j := 0; j < field.NumField(); j++ {
				nestedField := field.Field(j)
				nestedFieldType := field.Type().Field(j)
				if nestedField.Kind() == reflect.Ptr {
					nestedField = nestedField.Elem()
				}
				if nestedField.Kind() == reflect.String && nestedField.String() != "" {
					result += formatField(nestedFieldType.Name, nestedField.String())
				}
			}
		} else if field.Kind() == reflect.String && field.String() != "" {
			result += formatField(fieldType.Name, field.String())
		}
	}

	return result + "<eor>"
}

func formatField(fieldName string, fieldValue string) string {
	return fmt.Sprintf("<%s:%d>%s", fieldName, len(fieldValue), fieldValue)
}
