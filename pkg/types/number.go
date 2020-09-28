package types

import (
	"bytes"
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

type Number struct {
	Name string

	Enum []int64
}

var _ Type = &Number{}
var _ NamedType = &Object{}
var _ Enumerable = &Number{}

func (e *Number) UsedAsMapKey() bool {
	return false
}

func (n *Number) AddCandidates(v interface{}) {
	switch v.(type) {
	case int, int8, int16, int32, int64,
		uint, uint8, uint16, uint32, uint64:

		v, _ := strconv.ParseInt(fmt.Sprint(v), 10, 64)

		n.Enum = append(n.Enum, v)
	default:
		panic(fmt.Sprintf("unsupported type for number union type: %s", reflect.TypeOf(v)))
	}
}

func (n *Number) SetName(name string) {
	n.Name = name
}

func (n *Number) String() string {
	buf := bytes.NewBuffer(nil)

	buf.WriteString("Number")

	arr := make([]string, 0, 2)

	if n.Name != "" {
		arr = append(arr, n.Name)
	}
	if len(n.Enum) != 0 {
		enums := make([]string, len(n.Enum))
		for i := range n.Enum {
			enums = append(enums, strconv.FormatInt(n.Enum[i], 10))
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
