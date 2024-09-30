//go:build windows

package adif

import "github.com/go-playground/validator/v10"

type Header struct {
}

type Qso struct {
	ContactedStation *ContactedStation
	LoggingStation   *LoggingStation
	AntPath          string `json:"ant_path"`
	AntSect          string `json:"ant_sect"`
	AIndex           string `json:"a_index"`
	Band             string `json:"band"`
	Comment          string `json:"comment"`
	CommentIntl      string `json:"comment_intl"`
	Distance         string `json:"distance"`
	Freq             string `json:"freq" validate:"freqency-check" errormsg:"Invalid frequency"` // Freq: QSO frequency in Megahertz
	KIndex           string `json:"k_index"`
	Mode             string `json:"mode"`
	QsoDate          string `json:"qso_date"`
	QsoDateOff       string `json:"qso_date_off"`
	QsoRandom        string `json:"qso_random"`
	RstRcvd          string `json:"rst_rcvd"`
	RstSent          string `json:"rst_sent"`
	TimeOff          string `json:"time_off"`
	TimeOn           string `json:"time_on"`
}

type ContactedStation struct {
	Address       string `json:"address"`
	AddressIntl   string `json:"address_intl"`
	Age           string `json:"age"`
	Altitude      string `json:"altitude"`
	Call          string `json:"call" validate:"required"` // Call: the contacted station's callsign
	Cont          string `json:"cont"`                     // Cont: the contacted station's Continent
	ContactedOP   string `json:"contacted_op"`             // ContactedOP: the callsign of the individual operating the contacted station
	Country       string `json:"country"`                  // Country: the contacted station's DXCC entity name
	CountryIntl   string `json:"country_intl"`             // CountryIntl: the contacted station's DXCC entity name
	CQZ           string `json:"cqz"`                      // CQZ: the contacted station's CQ Zone in the range 1 to 40 (inclusive)
	DarcDok       string `json:"darc_dok"`                 // DarcDok: the contacted station's DARC DOK (District Location Code)
	DXCC          string `json:"dxcc"`                     // DXCC: the contacted station's DXCC Entity Code
	Email         string `json:"email"`                    // Email: the contacted station's email address
	EQCall        string `json:"eq_call"`                  // EQCall: the contacted station's owner's callsign
	Gridsquare    string `json:"gridsquare"`               // Gridsquare: the contacted station's 2-character, 4-character, 6-character, or 8-character Maidenhead Grid Square
	GridsquareExt string `json:"gridsquare_ext"`           // GridsquareExt:
	IOTA          string `json:"iota"`                     // IOTA: the contacted station's IOTA designator, in format CC-XXX
	IotaIslandID  string `json:"iota_island_id"`           // IotaIslandID: the contacted station's IOTA Island Identifier, an 8-digit integer in the range 1 to 99999999
	ITUZ          string `json:"ituz"`                     // ITUZ: the contacted station's ITU zone in the range 1 to 90 (inclusive)
	Lat           string `json:"lat"`                      // Lat: the contacted station's latitude
	Lon           string `json:"lon"`                      // Lon: the contacted station's longitude
	Name          string `json:"name"`                     // Name: the contacted station's operator's name
	NameIntl      string `json:"name_intl"`
	PFX           string `json:"pfx"`
	PotaRef       string `json:"pota_ref"`
	QTH           string `json:"qth"`
	QTHIntl       string `json:"qth_intl"`
	Region        string `json:"region"`
	Rig           string `json:"rig"`
	RigIntl       string `json:"rig_intl"`
	RxPwr         string `json:"rx_pwr"`
	Sig           string `json:"sig"`
	SigIntl       string `json:"sig_intl"`
	SigInfo       string `json:"sig_info"`
	SigInfoIntl   string `json:"sig_info_intl"`
	SotaRef       string `json:"sota_ref"`
	State         string `json:"state"`
	Web           string `json:"web"`
}

type LoggingStation struct {
	AntAZ           string `json:"ant_az"`
	AntEL           string `json:"ant_el"`
	BandRX          string `json:"band_rx"` // BandRX in a split frequency QSO, the logging station's receiving band
	FreqRX          string `json:"freq_rx"` // FreqRX in a split frequency QSO, the logging station's receiving frequency in Megahertz
	Altitude        string `json:"my_altitude"`
	Antenna         string `json:"my_antenna"`
	AntennaIntl     string `json:"my_antenna_intl"`
	City            string `json:"my_city"`
	Country         string `json:"my_country"`
	CountryIntl     string `json:"my_country_intl"`
	CQZone          string `json:"my_cq_zone"`
	DXCC            string `json:"my_dxcc"`
	Gridsquare      string `json:"my_gridsquare"`
	GridsquareExt   string `json:"my_gridsquare_ext"`
	IOTA            string `json:"my_iota"`
	IotaIslandID    string `json:"my_iota_island_id"`
	ItuZone         string `json:"my_itu_zone"`
	Lat             string `json:"my_lat"`
	Lon             string `json:"my_lon"`
	Name            string `json:"my_name" validate:"required"` // Name: the logging operator's name
	NameIntl        string `json:"my_name_intl"`                // NameIntl: the logging operator's name
	PostalCode      string `json:"my_postal_code"`
	PostalCodeIntl  string `json:"my_postal_code_intl"`
	PotaRef         string `json:"my_pota_ref"`
	Rig             string `json:"my_rig"`
	RigIntl         string `json:"my_rig_intl"`
	Sig             string `json:"my_sig"`
	SigIntl         string `json:"my_sig_intl"`
	SigInfo         string `json:"my_sig_info"`
	SigInfoIntl     string `json:"my_sig_info_intl"`
	SotaRef         string `json:"my_sota_ref"`
	State           string `json:"my_state"`
	Street          string `json:"my_street"`
	StreetIntl      string `json:"my_street_intl"`
	Notes           string `json:"notes"`
	NotesIntl       string `json:"notes_intl"`
	Operator        string `json:"operator"`
	OwnerCallsign   string `json:"owner_callsign"`
	StationCallsign string `json:"station_callsign" validate:"required"`
	TxPwr           string `json:"tx_pwr,required"`
}

type Qsl struct {
}

type Record struct {
	validate *validator.Validate
	HEADER   *Header
	QSO      *Qso `validate:"required"`
	QSL      *Qsl
}
