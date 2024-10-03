package adif

import "github.com/go-playground/validator/v10"

// NewQso creates a new Qso object
func NewQso(band, frequency, mode, qsoDate, timeOn, rstRcvd, rstSent string) (*Qso, error) {
	if validate == nil {
		validate = validator.New()
		if err := registerValidators(validate); err != nil {
			return nil, err
		}
	}

	return &Qso{
		Band:    band,
		Freq:    frequency,
		Mode:    mode,
		QsoDate: qsoDate,
		RstRcvd: rstRcvd,
		RstSent: rstSent,
		TimeOn:  timeOn,
	}, nil
}

func (q *Qso) Validate() error {
	return ValidateFunc[Qso](*q, validate)
}

func (q *Qso) SetLoggingStation(ptr *LoggingStation) error {
	if ptr == nil {
		return ErrorNilLoggingStation
	}
	q.LoggingStation = ptr
	return nil
}

func (q *Qso) SetContactedStation(ptr *ContactedStation) error {
	if ptr == nil {
		return ErrorNilContactedStation
	}
	q.ContactedStation = ptr
	return nil
}

func (q *Qso) ADIString() string {
	return parseStructToADIString(q) + eorStr
}
