package adif

import (
	"errors"
	"github.com/go-playground/validator/v10"
)

// NewQso creates a new Qso object populated with the given data, which is the minimum required for a valid QSO
// in the opinion of this module.
func NewQso(band, frequency, mode, qsoDate, timeOn, rstRcvd, rstSent string) (*Qso, error) {
	requiredFields := map[string]string{
		"band":      band,
		"frequency": frequency,
		"mode":      mode,
		"qsoDate":   qsoDate,
		"timeOn":    timeOn,
		"rstRcvd":   rstRcvd,
		"rstSent":   rstSent,
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

	return &Qso{
		Band:             band,
		Freq:             frequency,
		Mode:             mode,
		QsoDate:          qsoDate,
		RstRcvd:          rstRcvd,
		RstSent:          rstSent,
		TimeOn:           timeOn,
		LoggingStation:   new(LoggingStation),
		ContactedStation: new(ContactedStation),
		Qsl:              new(Qsl),
	}, nil
}

func (q *Qso) Validate() error {
	return ValidateFunc[Qso](*q, validate)
}

func (q *Qso) SetLoggingStation(ptr *LoggingStation) error {
	if ptr == nil {
		return ErrorNilLoggingStation
	}
	if q.LoggingStation == nil {
		q.LoggingStation = new(LoggingStation)
	}
	q.LoggingStation = ptr
	return nil
}

func (q *Qso) SetContactedStation(ptr *ContactedStation) error {
	if ptr == nil {
		return ErrorNilContactedStation
	}
	if q.ContactedStation == nil {
		q.ContactedStation = new(ContactedStation)
	}
	q.ContactedStation = ptr
	return nil
}

func (q *Qso) SetQsl(ptr *Qsl) error {
	if ptr == nil {
		return ErrorNilQsl
	}
	if q.Qsl == nil {
		q.Qsl = new(Qsl)
	}
	q.Qsl = ptr
	return nil
}

// ADIString returns the ADI string representation of the Qso object
func (q *Qso) ADIString() string {
	return parseStructToADIString(q) + eorStr
}
