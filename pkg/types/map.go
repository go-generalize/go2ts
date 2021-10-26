// Package types contains structs/interfaces representing TypeScript types
package types

import "fmt"

// Map - {[key: ...]: ...} in TypeScript
type Map struct {
	Common
	Key   Type
	Value Type
}

var _ Type = &Map{}

// UsedAsMapKey returns whether this type can be used as the key for map
func (e *Map) UsedAsMapKey() bool {
	return false
}

// String returns this type in string representation
func (e *Map) String() string {
	return fmt.Sprintf("Map{%s: %s}", e.Key.String(), e.Value.String())
}
