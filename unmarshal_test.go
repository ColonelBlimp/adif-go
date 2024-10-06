//go:build windows

package adif

import "testing"

var data = []byte("<ADIF_VER:5>3.1.4\n<PROGRAMID:18>7Q-Station-Manager\n<PROGRAMVERSION:5>0.0.1\n<eoh><app_qrzlog_status:1>C\n<cont:2>EU\n<time_off:4>1253\n<my_itu_zone:2>53\n<qrzcom_qso_upload_status:1>Y\n<my_cq_zone:2>37\n<eqsl_qsl_sent:1>N\n<distance:4>7733\n<lon:11>W000 10.547\n<gridsquare:8>IO91VL88\n<freq_rx:5>24.97\n<lotw_qsl_rcvd:1>N\n<eqsl_qsl_rcvd:1>N\n<dxcc:3>223\n<qso_date:8>20240615\n<my_lat:11>S011 26.635\n<call:5>M0ZNK\n<app_qrzlog_qsldate:8>20240617\n<rst_rcvd:2>59\n<my_city:5>Mzuzu\n<band_rx:3>12m\n<qsl_sent:1>N\n<my_lon:11>E034 00.576\n<band:3>12m\n<station_callsign:6>7Q5MLV\n<lotw_qsl_sent:1>N\n<my_gridsquare:6>KH78an\n<rst_sent:2>46\n<mode:3>USB\n<cqz:2>14\n<ituz:2>27\n<my_country:6>Malawi\n<time_on:4>1253\n<qrzcom_qso_upload_date:8>20240617\n<email:22>Jamesdando50@gmail.com\n<qso_date_off:8>20240615\n<qsl_rcvd:1>N\n<lat:11>N051 29.704\n<country:7>England\n<name:11>JAMES DANDO\n<my_name:12>Marc L Veary\n<tx_pwr:3>500\n<qth:8>Somerset\n<freq:5>24.97\n<app_qrzlog_logid:10>1112032999\n<eor>")

func TestUnmarshal(t *testing.T) {
	var record Record
	err := Unmarshal(data, &record)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(len(record.QsoSlice))
}
