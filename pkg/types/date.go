package types

type Date struct {
}

var _ Type = &Date{}

func (e *Date) UsedAsMapKey() bool {
	return false
}

func (n *Date) String() string {
	return "Date"
}
