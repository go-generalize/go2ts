// Package types contains structs/interfaces representing TypeScript types
package types

import (
	"bytes"
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

// Number - number in TypeScript
type Number struct {
	Name string

	Enum []int64
}

var _ Type = &Number{}
var _ NamedType = &Number{}
var _ Enumerable = &Number{}

// UsedAsMapKey returns whether this type can be used as the key for map
func (e *Number) UsedAsMapKey() bool {
	return len(e.Enum) == 0
}

// AddCandidates adds an candidate for enum
func (e *Number) AddCandidates(v interface{}) {
	switch v.(type) {
	case int, int8, int16, int32, int64,
		uint, uint8, uint16, uint32, uint64:

		v, _ := strconv.ParseInt(fmt.Sprint(v), 10, 64)

		e.Enum = append(e.Enum, v)
	default:
		panic(fmt.Sprintf("unsupported type for number union type: %s", reflect.TypeOf(v)))
	}
}

// SetName sets a alternative name
func (e *Number) SetName(name string) {
	e.Name = name
}

// String returns this type in string representation
func (e *Number) String() string {
	buf := bytes.NewBuffer(nil)

	buf.WriteString("Number")

	arr := make([]string, 0, 2)

	if e.Name != "" {
		arr = append(arr, e.Name)
	}
	if len(e.Enum) != 0 {
		enums := make([]string, len(e.Enum))
		for i := range e.Enum {
			enums = append(enums, strconv.FormatInt(e.Enum[i], 10))
		}

		arr = append(arr, "["+strings.Join(enums, ",")+"]")
	}

	if len(arr) != 0 {
		buf.WriteString("(")
		buf.WriteString(strings.Join(arr, ", "))
		buf.WriteString(")")
	}

	return buf.String()
}
