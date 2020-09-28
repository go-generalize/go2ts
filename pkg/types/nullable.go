package types

type Nullable struct {
	Inner Type
}

var _ Type = &Nullable{}

func (e *Nullable) UsedAsMapKey() bool {
	return false
}

func (n *Nullable) String() string {
	return "Nullable(" + n.Inner.String() + ")"
}
