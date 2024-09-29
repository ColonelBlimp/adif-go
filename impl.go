//go:build windows

package adif

func New() *Record {
	return &Record{
		HEADER: Header{},
		QSO:    Qso{},
		QSL:    Qsl{},
	}
}

func (r *Record) SetHeader(header Header) {
	r.HEADER = header
}

func (r *Record) SetContactedStationData(v ContactedStation) {
	r.QSO.ContactedStation = v
}

func (r *Record) SetLoggingStationData(v LoggingStation) {
	r.QSO.LoggingStation = v
}

func (r *Record) SetQSLData(v Qsl) {
	r.QSL = v
}
