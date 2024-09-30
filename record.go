//go:build windows

package adif

import (
	"github.com/go-playground/validator/v10"
)

const tagCustom = "errormsg"

func NewRecord(qso *Qso) *Record {
	validate := validator.New()
	validate.RegisterValidation("freqency-check", frequencyCheck)

	return &Record{
		validate: validate,
		QSO:      qso,
	}
}

func (r *Record) SetHeader(ptr *Header) error {
	err := r.validate.Struct(ptr)
	if err != nil {
		return err
	}
	r.HEADER = ptr
	return nil
}

func (r *Record) Validate() (errs error) {
	return ValidateFunc[Record](*r, r.validate)
}
