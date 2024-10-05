package adif

// NewContactedStation creates a new ContactedStation struct with the provided call and name. Returns error if call or name is empty.
func NewContactedStation(call, name string) (*ContactedStation, error) {
	if call == emptyStr {
		return nil, ErrorCallEmpty
	}
	if name == emptyStr {
		return nil, ErrorNameEmpty
	}
	return &ContactedStation{
		Call: call,
		Name: name,
	}, nil
}
