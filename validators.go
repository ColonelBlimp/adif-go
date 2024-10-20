package adif

import (
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
	"reflect"
	"strings"
)

func registerValidators(validate *validator.Validate) error {
	if err := validate.RegisterValidation("freqency-check", validateFrequency); err != nil {
		return err
	}
	if err := validate.RegisterValidation("band-check", validateBand); err != nil {
		return err
	}
	if err := validate.RegisterValidation("mode-check", validateMode); err != nil {
		return err
	}
	return nil
}

func errorTagFunc[T interface{}](obj interface{}, snp string, fieldname, actualTag string) error {
	o := obj.(T)

	if !strings.Contains(snp, fieldname) {
		return nil
	}

	fieldArr := strings.Split(snp, dotStr)
	rsf := reflect.TypeOf(o)

	for i := 1; i < len(fieldArr); i++ {
		field, found := rsf.FieldByName(fieldArr[i])
		if found {
			if fieldArr[i] == fieldname {
				customMessage := field.Tag.Get(errMsgTag)
				if customMessage != emptyStr {
					return fmt.Errorf("%s: %s (%s)", fieldname, customMessage, actualTag)
				}
				return nil
			} else {
				if field.Type.Kind() == reflect.Ptr {
					// If the field type is a pointer, dereference it
					rsf = field.Type.Elem()
				} else {
					rsf = field.Type
				}
			}
		}
	}
	return nil
}

// ValidateFunc do not use this function use validates the object using the validator.Validate instance.
func validateFunc[T interface{}](obj interface{}, validate *validator.Validate) (errs error) {
	o := obj.(T)

	defer func() {
		if r := recover(); r != nil {
			errs = fmt.Errorf("unable to validate: %+v", r)
		}
	}()

	if err := validate.Struct(o); err != nil {
		var errorValid validator.ValidationErrors
		if ok := errors.As(err, &errorValid); ok {
			for _, e := range errorValid {
				snp := e.StructNamespace()
				errmgs := errorTagFunc[T](obj, snp, e.Field(), e.ActualTag())
				if errmgs != nil {
					errs = errors.Join(errs, fmt.Errorf("%w", errmgs))
				} else {
					errs = errors.Join(errs, fmt.Errorf("%w", e))
				}
			}
		}
	}

	if errs != nil {
		return errs
	}

	return nil
}

func validateFrequency(fl validator.FieldLevel) bool {
	freq := fl.Field().String()
	if len(freq) < 5 || len(freq) > 6 {
		return false
	}
	if !isNthRuneFromRightEqual(freq, 4, '.') {
		return false
	}
	parts := strings.Split(freq, emptyStr)
	if !isAllDigits(parts[0]) || !isAllDigits(parts[1]) {
		return false
	}
	return true
}

func validateBand(fl validator.FieldLevel) bool {
	band := fl.Field().String()
	switch band {
	case "160m", "80m", "40m", "30m", "20m", "17m", "15m", "12m", "10m", "6m", "2m", "70cm", "23cm", "13cm", "9cm", "6cm", "3cm", "1.25cm":
		return true
	default:
		return false
	}
}

func validateMode(fl validator.FieldLevel) bool {
	mode := fl.Field().String()
	switch mode {
	case "SSB", "LSB", "USB", "CW-L", "CW-U", "FM", "FM-N", "AM", "AM-N", "RTTY-L", "RTTY-U", "PSK", "DATA-L", "DATA-U", "DATA-FM", "DATA-FM-N":
		return true
	default:
		return false
	}
}

// Validate checks the provided object against the rules defined by the validators in validate.
// It registers custom validators if they are not already registered. This function is exported and can be used
// outside the module to validate structs which embed the adifgo structs. Returns an error if validation fails.
func Validate[T interface{}](obj interface{}) error {
	if validate == nil {
		validate = validator.New()
		if err := registerValidators(validate); err != nil {
			return err
		}
	}

	return validateFunc[T](obj, validate)
}
