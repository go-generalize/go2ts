// Package types contains structs/interfaces representing TypeScript types
package types

// Date - RFC3399 string in TypeScript
type Date struct {
	Common
}

var _ Type = &Date{}

// UsedAsMapKey returns whether this type can be used as the key for map
func (*Date) UsedAsMapKey() bool {
	return false
}

// String returns this type in string representation
func (*Date) String() string {
	return "Date"
}
