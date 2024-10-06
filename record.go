//go:build windows

package adif

import (
	"errors"
	"github.com/go-playground/validator/v10"
)

const (
	adifFormat    = "<%s:%d>%s"
	errMsgTag     = "errormsg"
	jsonStructTag = "json"
	emptyStr      = ""
	dotStr        = "."
	eorStr        = "<EOR>"
)

var validate *validator.Validate

// NewRecord creates a new ADI Record object
func NewRecord(adifVersion, createdTimestamp, programID, programVersion string) (*Record, error) {
	requiredFields := map[string]string{
		"adifVersion":      adifVersion,
		"createdTimestamp": createdTimestamp,
		"programID":        programID,
		"programVersion":   programVersion,
	}

	for field, value := range requiredFields {
		if value == "" {
			return nil, errors.New(field + " parameter is empty")
		}
	}
	if validate == nil {
		validate = validator.New()
		if err := registerValidators(validate); err != nil {
			return nil, err
		}
	}

	return &Record{
		ADIFVer:          adifVersion,
		CreatedTimestamp: createdTimestamp,
		ProgramID:        programID,
		ProgramVersion:   programVersion,
		QsoSlice:         make(QsoSlice, 0),
	}, nil
}

func (r *Record) AddQso(ptr *Qso) error {
	if ptr == nil {
		return ErrorNilQso
	}
	//if r.QSOS == nil {
	//	r.QSOS = make(QsoSlice, 0)
	//}
	//r.QSOS = append(r.QSOS, ptr)

	return nil
}

//func (r *Record) SetHeader(ptr *Header) error {
//	if ptr == nil {
//		return ErrorNilHeader
//	}
//	r.HEADER = ptr
//	return nil
//}

//func (r *Record) SetQsl(ptr *Qsl) error {
//	if ptr == nil {
//		return ErrorNilQsl
//	}
//	r.QSL = ptr
//	return nil
//}

// Validate validates the Record object. The fields which are checked are:
// Qso.Band, Qso.Freq, Qso.Mode, Qso.QsoDate, Qso.RstRcvd, Qso.RstSent, Qso.TimeOn
// ContactedStation.Call, LoggingStation.Name, LoggingStation.StationCallsign
// If any of the fields are invalid, the function returns an error with an appropriate error message.
func (r *Record) Validate() error {
	return ValidateFunc[Record](*r, validate)
}

func (r *Record) ADIString() string {
	return parseStructToADIString(r) + eorStr
}
