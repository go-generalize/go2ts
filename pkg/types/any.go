// Package types contains structs/interfaces representing TypeScript types
package types

// Any - any in TypeScript
type Any struct {
	Common
}

var _ Type = &Any{}

// UsedAsMapKey returns whether this type can be used as the key for map
func (e *Any) UsedAsMapKey() bool {
	return false
}

// String returns this type in string representation
func (e *Any) String() string {
	return "Any"
}
