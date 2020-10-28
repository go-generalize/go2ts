// Package types contains structs/interfaces representing TypeScript types
package types

// Type interface represents all TypeScript types
type Type interface {
	UsedAsMapKey() bool
	String() string
}

// Enumerable interface represents union types
type Enumerable interface {
	Type

	// AddCandidates adds an candidate for enum
	AddCandidates(v interface{})
}

// NamedType interface represents named types
type NamedType interface {
	Type

	// SetName sets a alternative name
	SetName(name string)
}
