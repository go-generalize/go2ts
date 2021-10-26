// Package types contains structs/interfaces representing TypeScript types
package types

// Nullable - ... | null
type Nullable struct {
	Common
	Inner Type
}

var _ Type = &Nullable{}

// UsedAsMapKey returns whether this type can be used as the key for map
func (e *Nullable) UsedAsMapKey() bool {
	return false
}

// String returns this type in string representation
func (e *Nullable) String() string {
	return "Nullable(" + e.Inner.String() + ")"
}
