package types

type Any struct{}

var _ Type = &Any{}

func (e *Any) UsedAsMapKey() bool {
	return false
}

func (e *Any) String() string {
	return "Any"
}
