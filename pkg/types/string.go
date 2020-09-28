package types

import (
	"bytes"
	"fmt"
	"reflect"
	"strings"
)

type String struct {
	Name string

	Enum []string
}

var _ Type = &String{}
var _ NamedType = &Object{}
var _ Enumerable = &String{}

func (e *String) UsedAsMapKey() bool {
	return false
}

func (n *String) AddCandidates(v interface{}) {
	switch v := v.(type) {
	case string:
		n.Enum = append(n.Enum, v)
	default:
		panic(fmt.Sprintf("unsupported type for string union type: %s", reflect.TypeOf(v)))
	}
}

func (n *String) SetName(name string) {
	n.Name = name
}

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
