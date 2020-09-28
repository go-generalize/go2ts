package types

type Array struct {
	Inner Type
}

var _ Type = &Array{}

func (e *Array) UsedAsMapKey() bool {
	return false
}

func (n *Array) String() string {
	return "Array(" + n.Inner.String() + ")"
}
