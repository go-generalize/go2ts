package types

type Type interface {
	UsedAsMapKey() bool
	String() string
}

type Enumerable interface {
	Type
	AddCandidates(v interface{})
}

type NamedType interface {
	Type
	SetName(name string)
}
