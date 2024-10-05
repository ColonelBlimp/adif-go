# ADI, ADIF, ADX Library

[![Maintainability](https://api.codeclimate.com/v1/badges/cccdec9522213a066d25/maintainability)](https://codeclimate.com/github/ColonelBlimp/adif/maintainability)

# Introduction
This library is *not* intended to be a complete implementation of the [ADIF spec](https://www.adif.org/), but was written
as part of the **7Q-Station-Manager-Desktop** application. The module is intended to be used in conjunction with
**7Q-Station-Manager-Desktop** and is not intended to bea standalone library. It can be used as a standalone,
but it has its limitations. The module is **highly opinionated concerning what fields are and are not required**
for a QSO to be considered valid; this is heavily influenced by QRZ.com and ClubLog.org

# Development

## Requirements

```
go get github.com/go-playground/validator/v10
```

# Usage

## Creating a QSO

```
	contactedStation, err := NewContactedStation("XX1XXX")
	if err != nil {
	    ...
	}
	
	loggingStation, err := NewLoggingStation("Y1YY", "My Name")
	if err != nil {
	    ...
	}
	
	qso, err := NewQso("15m", "21.250", "USB", "20240929", "1621", "59", "59")
	if err != nil {
	    ...
	}

	if err = qso.SetContactedStation(contactedStation); err != nil {
        ...
	}

	if err = qso.SetLoggingStation(loggingStation); err != nil {
	    ...
	}

	if err = qso.Validate(); err != nil {
        ...
	}
```


# Validation
The required fields for a valid QSO are:

- Qso.Band
- Qso.Freq
- Qso.Mode
- Qso.QsoDate
- Qso.RstRcvd
- Qso.RstSent
- Qso.TimeOn
- ContactStation.Call
- ContactStation.Name
- LoggingStation.StationCallsign
