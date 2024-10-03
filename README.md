# ADI, ADIF, ADX Library

[![Maintainability](https://api.codeclimate.com/v1/badges/cccdec9522213a066d25/maintainability)](https://codeclimate.com/github/ColonelBlimp/adif/maintainability)

# Introduction
This library is *not* intended to be a complete implementation of the [ADIF spec](https://www.adif.org/), but was written
as part of the **7Q-Station-Master-Desktop** application. The module is intended to be used in conjunction with
**7Q-Station-Master-Desktop** and is not intended to bea standalone library. It can be used as a standalone,
but it has its limitations. The module is **highly opinionated concerning what fields are and are not required**
for a QSO to be considered valid; this is heavily influenced by QRZ.com and ClubLog.org

# Development

## Requirements

```
go get github.com/go-playground/validator/v10
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
- LoggingStation.Name
- LoggingStation.StationCallsign
