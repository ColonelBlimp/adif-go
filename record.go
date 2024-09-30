//go:build windows

package adif

import (
	"github.com/go-playground/validator/v10"
)

const errMsgTag = "errormsg"

func NewRecord(qso *Qso) *Record {
	validate := validator.New()
	validate.RegisterValidation("freqency-check", validateFrequency)
	validate.RegisterValidation("band-check", validateBand)
	validate.RegisterValidation("mode-check", validateMode)

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
