package adif

func NewLoggingStation(stationCallsign, name string) *LoggingStation {
	return &LoggingStation{
		Name:            name,
		StationCallsign: stationCallsign,
	}
}
