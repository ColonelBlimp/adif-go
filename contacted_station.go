//go:build windows

package adif

func NewContactedStation(call string) *ContactedStation {
	return &ContactedStation{
		Call: call,
	}
}
