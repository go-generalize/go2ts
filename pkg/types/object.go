// Package types contains structs/interfaces representing TypeScript types
package types

import (
	"bytes"
	"fmt"
	"strings"
)

// ObjectEntry is an field in objects
type ObjectEntry struct {
	Type     Type
	Optional bool
}

// Object - {field1: ..., field2: ...} in TypeScript
type Object struct {
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
	buf := bytes.NewBuffer(nil)

	buf.WriteString("Object")

	arr := make([]string, 0, len(n.Entries))

	for key, val := range n.Entries {
		arr = append(arr, fmt.Sprintf("%s:%s", key, val.Type.String()))
	}

	buf.WriteString("{")
	buf.WriteString(strings.Join(arr, ","))
	buf.WriteString("}")

	return buf.String()
}
