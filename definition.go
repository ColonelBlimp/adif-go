package adif

type QsoSlice []*Qso

type Qso struct {
	*ContactedStation
	*LoggingStation
	*Qsl
	AntPath     string `json:"ant_path,omitempty"`
	AntSect     string `json:"ant_sect,omitempty"`
	AIndex      string `json:"a_index,omitempty"`
	Band        string `json:"band" validate:"band-check" errormsg:"Invalid band"` // Band: the QSO band
	Comment     string `json:"comment,omitempty"`
	CommentIntl string `json:"comment_intl,omitempty"`
	Distance    string `json:"distance,omitempty"`
	Freq        string `json:"freq" validate:"freqency-check" errormsg:"Invalid frequency"` // Freq: QSO frequency in Megahertz
	KIndex      string `json:"k_index,omitempty"`
	Mode        string `json:"mode" validate:"mode-check" errormsg:"Invalid mode"` // Mode: the QSO mode
	QsoComplete string `json:"qso_complete,omitempty"`
	QsoDate     string `json:"qso_date" validate:"required,len=8,numeric"` // QsoDate: the QSO date in the format YYYYMMDD
	QsoDateOff  string `json:"qso_date_off,omitempty"`
	QsoRandom   string `json:"qso_random,omitempty"`
	RstRcvd     string `json:"rst_rcvd" validate:"required,min=2,max=3,numeric"` // RstRcvd: the RST code received
	RstSent     string `json:"rst_sent" validate:"required,min=2,max=3,numeric"` // RstSent: the RST code sent
	TimeOff     string `json:"time_off,omitempty"`
	TimeOn      string `json:"time_on" validate:"required,len=4,numeric"` // TimeOn: the QSO time in the format HHMM
}

type ContactedStation struct {
	Address       string `json:"address,omitempty"`
	AddressIntl   string `json:"address_intl,omitempty"`
	Age           string `json:"age,omitempty"`
	Altitude      string `json:"altitude,omitempty"`
	Call          string `json:"call" validate:"required,min=3"` // Call: the contacted station's callsign
	Cont          string `json:"cont,omitempty"`                 // Cont: the contacted station's Continent
	ContactedOP   string `json:"contacted_op,omitempty"`         // ContactedOP: the callsign of the individual operating the contacted station
	Country       string `json:"country,omitempty"`              // Country: the contacted station's DXCC entity name
	CountryIntl   string `json:"country_intl,omitempty"`         // CountryIntl: the contacted station's DXCC entity name
	CQZ           string `json:"cqz,omitempty"`                  // CQZ: the contacted station's CQ Zone in the range 1 to 40 (inclusive)
	DarcDok       string `json:"darc_dok,omitempty"`             // DarcDok: the contacted station's DARC DOK (District Location Code)
	DXCC          string `json:"dxcc,omitempty"`                 // DXCC: the contacted station's DXCC Entity Code
	Email         string `json:"email,omitempty"`                // Email: the contacted station's email address
	EQCall        string `json:"eq_call,omitempty"`              // EQCall: the contacted station's owner's callsign
	Gridsquare    string `json:"gridsquare,omitempty"`           // Gridsquare: the contacted station's 2-character, 4-character, 6-character, or 8-character Maidenhead Grid Square
	GridsquareExt string `json:"gridsquare_ext,omitempty"`       // GridsquareExt:
	IOTA          string `json:"iota,omitempty"`                 // IOTA: the contacted station's IOTA designator, in format CC-XXX
	IotaIslandID  string `json:"iota_island_id,omitempty"`       // IotaIslandID: the contacted station's IOTA Island Identifier, an 8-digit integer in the range 1 to 99999999
	ITUZ          string `json:"ituz,omitempty"`                 // ITUZ: the contacted station's ITU zone in the range 1 to 90 (inclusive)
	Lat           string `json:"lat,omitempty"`                  // Lat: the contacted station's latitude
	Lon           string `json:"lon,omitempty"`                  // Lon: the contacted station's longitude
	Name          string `json:"name" validate:"required,min=2"` // Name: the contacted station's operator's name
	NameIntl      string `json:"name_intl,omitempty"`
	PFX           string `json:"pfx,omitempty"`
	PotaRef       string `json:"pota_ref,omitempty"`
	QTH           string `json:"qth,omitempty"`
	QTHIntl       string `json:"qth_intl,omitempty"`
	Region        string `json:"region,omitempty"`
	Rig           string `json:"rig,omitempty"`
	RigIntl       string `json:"rig_intl,omitempty"`
	RxPwr         string `json:"rx_pwr,omitempty"`
	Sig           string `json:"sig,omitempty"`
	SigIntl       string `json:"sig_intl,omitempty"`
	SigInfo       string `json:"sig_info,omitempty"`
	SigInfoIntl   string `json:"sig_info_intl,omitempty"`
	SotaRef       string `json:"sota_ref,omitempty"`
	State         string `json:"state,omitempty"`
	Web           string `json:"web,omitempty"`
}

type LoggingStation struct {
	AntAZ            string `json:"ant_az,omitempty"`
	AntEL            string `json:"ant_el,omitempty"`
	BandRX           string `json:"band_rx,omitempty"` // BandRX in a split frequency QSO, the logging station's receiving band
	FreqRX           string `json:"freq_rx,omitempty"` // FreqRX in a split frequency QSO, the logging station's receiving frequency in Megahertz
	MyAltitude       string `json:"my_altitude,omitempty"`
	MyAntenna        string `json:"my_antenna,omitempty"`
	MyAntennaIntl    string `json:"my_antenna_intl,omitempty"`
	MyCity           string `json:"my_city,omitempty"`
	MyCountry        string `json:"my_country,omitempty"`
	MyCountryIntl    string `json:"my_country_intl,omitempty"`
	MyCQZone         string `json:"my_cq_zone,omitempty"`
	MyDXCC           string `json:"my_dxcc,omitempty"`
	MyGridsquare     string `json:"my_gridsquare,omitempty"`
	MyGridsquareExt  string `json:"my_gridsquare_ext,omitempty"`
	MyIOTA           string `json:"my_iota,omitempty"`
	MyIotaIslandID   string `json:"my_iota_island_id,omitempty"`
	MyItuZone        string `json:"my_itu_zone,omitempty"`
	MyLat            string `json:"my_lat,omitempty"`
	MyLon            string `json:"my_lon,omitempty"`
	MyName           string `json:"my_name"`                // Name: the logging operator's name
	MyNameIntl       string `json:"my_name_intl,omitempty"` // NameIntl: the logging operator's name
	MyPostalCode     string `json:"my_postal_code,omitempty"`
	MyPostalCodeIntl string `json:"my_postal_code_intl,omitempty"`
	MyPotaRef        string `json:"my_pota_ref,omitempty"`
	MyRig            string `json:"my_rig,omitempty"`
	MyRigIntl        string `json:"my_rig_intl,omitempty"`
	MySig            string `json:"my_sig,omitempty"`
	MySigIntl        string `json:"my_sig_intl,omitempty"`
	MySigInfo        string `json:"my_sig_info,omitempty"`
	MySigInfoIntl    string `json:"my_sig_info_intl,omitempty"`
	MySotaRef        string `json:"my_sota_ref,omitempty"`
	MyState          string `json:"my_state,omitempty"`
	MyStreet         string `json:"my_street,omitempty"`
	MyStreetIntl     string `json:"my_street_intl,omitempty"`
	Notes            string `json:"notes,omitempty"`
	NotesIntl        string `json:"notes_intl,omitempty"`
	Operator         string `json:"operator,omitempty"` // Operator: the logging operator's callsign if StationCallsign is absent
	OwnerCallsign    string `json:"owner_callsign,omitempty"`
	StationCallsign  string `json:"station_callsign" validate:"required"` // StationCallsign: the logging station's callsign (the callsign used over the air). If StationCallsign is absent, Operator shall be treated as both the logging station's callsign and the logging operator's callsign
	TxPwr            string `json:"tx_pwr,omitempty"`
}

type Qsl struct {
	QslMsg     string `json:"qslmsg,omitempty"`
	QslMsgIntl string `json:"qslmsg_intl,omitempty"`
	QslRDate   string `json:"qslrdate,omitempty"`
	QslSDate   string `json:"qslsdate,omitempty"`
	QslRcvd    string `json:"qsl_rcvd" validate:"required,len=1"` // QslRcvd: the QSL received status
	QslSent    string `json:"qsl_sent" validate:"required,len=1"` // QslSent: the QSL sent status
	QslSentVia string `json:"qsl_sent_via,omitempty"`
	QslVia     string `json:"qsl_via,omitempty"`
}

type Record struct {
	ADIFVer          string `json:"adif_ver"`          // ADIFVer: the ADIF version number
	CreatedTimestamp string `json:"created_timestamp"` // CreatedTimestamp: the timestamp when the ADIF file was created
	ProgramID        string `json:"programid"`         // ProgramID: the name of the logging program
	ProgramVersion   string `json:"programversion"`    // ProgramVersion: the version of the logging program
	QsoSlice         `validate:"required" errormsg:"QSOS slice is required"`
}
