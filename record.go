//go:build windows

package adif

import (
	"github.com/go-playground/validator/v10"
)

const errMsgTag = "errormsg"

func NewRecord(qso *Qso) (*Record, error) {
	validate := validator.New()
	if err := validate.RegisterValidation("freqency-check", validateFrequency); err != nil {
		return nil, err
	}
	if err := validate.RegisterValidation("band-check", validateBand); err != nil {
		return nil, err
	}
	if err := validate.RegisterValidation("mode-check", validateMode); err != nil {
		return nil, err
	}

	return &Record{
		validate: validate,
		QSO:      qso,
	}, nil
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
