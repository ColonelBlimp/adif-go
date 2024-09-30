package adif

import "errors"

var (
	ErrorNilHeader = errors.New("header object is nil")
	ErrorNilQsl    = errors.New("qsl object is nil")
	ErrorNilQso    = errors.New("qso object is nil")
)
