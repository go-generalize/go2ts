// Package types contains structs/interfaces representing TypeScript types
package types

// Array - array in TypeScript
type Array struct {
	Common
	Inner Type
}

var _ Type = &Array{}

// UsedAsMapKey returns whether this type can be used as the key for map
func (a *Array) UsedAsMapKey() bool {
	return false
}

// String returns this type in string representation
func (a *Array) String() string {
	return "Array(" + a.Inner.String() + ")"
}
