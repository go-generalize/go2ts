package types

type Map struct {
	Key   Type
	Value Type
}

var _ Type = &Map{}

func (e *Map) UsedAsMapKey() bool {
	return false
}

func (n *Map) String() string {
	return "Map"
}
