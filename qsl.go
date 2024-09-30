package adif

func NewQsl() *Qsl {
	return &Qsl{
		QslRcvd: "N",
		QslSent: "N",
	}
}
