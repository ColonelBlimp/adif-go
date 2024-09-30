//go:build windows

package adif

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

func (q *Qso) SetLoggingStation(ptr *LoggingStation) {
	q.LoggingStation = ptr
}

func (q *Qso) SetContactedStation(ptr *ContactedStation) {
	q.ContactedStation = ptr
}
