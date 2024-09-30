# ADI, ADIF, ADX Library

# Introduction
This library is *not* intended to be a complete implementation of the ADI spec, but was written as part of
the **7Q-App**. The module is intended to be used in conjunction with the **7Q-App** and is not intended to be
a standalone library. It can be used as a standalone, but it has its limitations. The module is **opinionated
concerning what fields are and are not required** for a QSO to be considered valid; this is heavily influenced
by QRZ.com and ClubLog.org

# Development

## Requirements

```
go get github.com/go-playground/validator/v10
```

# Validation
The required fields are:

- Qso.Band
- Qso.Freq
- Qso.Mode
- Qso.QsoDate
- Qso.RstRcvd
- Qso.RstSent
- Qso.TimeOn
- ContactStation.Call
- LoggingStation.Name
- LoggingStation.StationCallsign
