package adif

func NewContactedStation(call, name string) *ContactedStation {
	return &ContactedStation{
		Call: call,
		Name: name,
	}
}
