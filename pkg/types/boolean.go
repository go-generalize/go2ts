package types

type Boolean struct {
}

var _ Type = &Boolean{}

func (e *Boolean) UsedAsMapKey() bool {
	return false
}

func (n *Boolean) String() string {
	return "Boolean"
}
