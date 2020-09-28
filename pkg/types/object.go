package types

import (
	"bytes"
	"fmt"
	"strings"
)

type ObjectEntry struct {
	Type     Type
	Optional bool
}

type Object struct {
	Name string

	Entries map[string]ObjectEntry
}

var _ Type = &Object{}
var _ NamedType = &Object{}

func (e *Object) UsedAsMapKey() bool {
	return false
}

func (n *Object) SetName(name string) {
	n.Name = name
}

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
