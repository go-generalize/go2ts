// Package types contains structs/interfaces representing TypeScript types
package types

// Type interface represents all TypeScript types handled by go2ts
type Type interface {
	UsedAsMapKey() bool
	String() string
}

// Enumerable interface represents union types
type Enumerable interface {
	Type

	// AddCandidates adds a candidate for enum
	AddCandidates(v interface{})
}

// NamedType interface represents named types
type NamedType interface {
	Type

	// SetName sets an alternative name
	SetName(name string)
}
