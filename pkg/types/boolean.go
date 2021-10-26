// Package types contains structs/interfaces representing TypeScript types
package types

// Boolean - boolean in TypeScript
type Boolean struct {
	Common
}

var _ Type = &Boolean{}

// UsedAsMapKey returns whether this type can be used as the key for map
func (b *Boolean) UsedAsMapKey() bool {
	return false
}

// String returns this type in string representation
func (b *Boolean) String() string {
	return "Boolean"
}
