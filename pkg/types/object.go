// Package types contains structs/interfaces representing TypeScript types
package types

import "go/token"

// ObjectEntry is an field in objects
type ObjectEntry struct {
	RawName    string
	RawTag     string
	FieldIndex int

	Type     Type
	Position *token.Position
	Optional bool
}

// Object - {field1: ..., field2: ...} in TypeScript
type Object struct {
	Common
	Name string

	Entries map[string]ObjectEntry
}

var _ Type = &Object{}
var _ NamedType = &Object{}

// UsedAsMapKey returns whether this type can be used as the key for map
func (n *Object) UsedAsMapKey() bool {
	return false
}

// SetName sets a alternative name
func (n *Object) SetName(name string) {
	n.Name = name
}

// String returns this type in string representation
func (n *Object) String() string {
	return "Object{...}"
}
