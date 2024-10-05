package adif

// NewLoggingStation creates a new LoggingStation with the specified station callsign. Returns an error if the callsign is empty.
func NewLoggingStation(stationCallsign string) (*LoggingStation, error) {
	if stationCallsign == emptyStr {
		return nil, ErrorCallsignEmpty
	}
	return &LoggingStation{
		StationCallsign: stationCallsign,
	}, nil
}
