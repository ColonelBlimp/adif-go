//go:build windows

package adif

import (
	"github.com/go-playground/validator/v10"
)

const errMsgTag = "errormsg"

// NewRecord creates a new ADI Record object
func NewRecord(qso *Qso) (*Record, error) {
	if qso == nil {
		return nil, ErrorNilQso
	}
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
