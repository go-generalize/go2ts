// Package types contains structs/interfaces representing TypeScript types
package types

import (
	"bytes"
	"fmt"
	"reflect"
	"strings"
)

// RawStringEnumCandidate represents a raw candidate for string enum
type RawStringEnumCandidate struct {
	Key   string
	Value string
}

// String - string in TypeScript
type String struct {
	Common
	Name string

	Enum    []string
	RawEnum []RawStringEnumCandidate
}

var _ Type = &String{}
var _ NamedType = &String{}
var _ Enumerable = &String{}

// UsedAsMapKey returns whether this type can be used as the key for map
func (n *String) UsedAsMapKey() bool {
	return len(n.Enum) == 0
}

// AddCandidates adds an candidate for enum
func (n *String) AddCandidates(key string, v interface{}) {
	switch v := v.(type) {
	case string:
		n.Enum = append(n.Enum, v)

		n.RawEnum = append(n.RawEnum, RawStringEnumCandidate{
			Key:   key,
			Value: v,
		})
	default:
		panic(fmt.Sprintf("unsupported type for string union type: %s", reflect.TypeOf(v)))
	}
}

// SetName sets a alternative name
func (n *String) SetName(name string) {
	n.Name = name
}

// String returns this type in string representation
func (n *String) String() string {
	buf := bytes.NewBuffer(nil)

	buf.WriteString("String")

	arr := make([]string, 0, 2)

	if n.Name != "" {
		arr = append(arr, n.Name)
	}
	if len(n.Enum) != 0 {
		arr = append(arr, "["+strings.Join(n.Enum, ",")+"]")
	}

	if len(arr) != 0 {
		buf.WriteString("(")
		buf.WriteString(strings.Join(arr, ", "))
		buf.WriteString(")")
	}

	return buf.String()
}
