package adif

// NewQsl creates a new Qsl instance with default values for QslRcvd and QslSent set to "N".
func NewQsl() *Qsl {
	return &Qsl{
		QslRcvd: "N",
		QslSent: "N",
	}
}
