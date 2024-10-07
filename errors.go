package adif

import "errors"

var (
	ErrorNilQsl              = errors.New("qsl object is nil")
	ErrorNilQso              = errors.New("qso slice object is nil")
	ErrorNilLoggingStation   = errors.New("logging station object is nil")
	ErrorNilContactedStation = errors.New("contacted station object is nil")
	ErrorCallEmpty           = errors.New("contacted station's call(sign) parameter is empty")
	ErrorNameEmpty           = errors.New("contacted station's name parameter is empty")
	ErrorCallsignEmpty       = errors.New("logging station's stationCallsign parameter is empty")
)
