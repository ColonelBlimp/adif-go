package adif

// NewQso creates a new Qso object
func NewQso(band, frequency, mode, qsoDate, timeOn, rstRcvd, rstSent string) *Qso {
	return &Qso{
		Band:    band,
		Freq:    frequency,
		Mode:    mode,
		QsoDate: qsoDate,
		RstRcvd: rstRcvd,
		RstSent: rstSent,
		TimeOn:  timeOn,
	}
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
