//go:build windows

package adif

import (
	"encoding/json"
	"testing"
)

func TestRecord(t *testing.T) {
	rec := Record{
		HEADER: Header{},
		QSO: Qso{
			StationContacted: ContactedStation{},
			StationLogging:   LoggingStation{},
			AntPath:          "",
			AntSect:          "",
			AIndex:           "",
			Band:             "",
			Comment:          "",
			CommentIntl:      "",
			Distance:         "",
			Freq:             "",
			KIndex:           "",
			Mode:             "",
			QsoDate:          "",
			QsoDateOff:       "",
			QsoRandom:        "",
			RstRcvd:          "",
			RstSent:          "",
			TimeOff:          "",
			TimeOn:           "",
		},
		QSL: Qsl{},
	}

	data, err := json.Marshal(rec)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	t.Log(string(data))
}
