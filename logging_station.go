package adif

func NewLoggingStation(stationCallsign string) *LoggingStation {
	return &LoggingStation{
		StationCallsign: stationCallsign,
	}
}
