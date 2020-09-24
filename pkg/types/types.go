package types

import (
	"fmt"
	"reflect"
	"strconv"
)

type Type interface {
}

type Enumerable interface {
	AddCandidates(v interface{})
}

type NamedType interface {
	SetName(name string)
}

type ObjectEntry struct {
	Type     Type
	Optional bool
}

type Object struct {
	Name string

	Entries map[string]ObjectEntry
}

func (n *Object) SetName(name string) {
	n.Name = name
}

type Array struct {
	Inner Type
}

type Date struct {
}

type Number struct {
	Name string

	Enum []int64
}

func (n *Number) AddCandidates(v interface{}) {
	switch v.(type) {
	case int, int8, int16, int32, int64,
		uint, uint8, uint16, uint32, uint64:

		v, _ := strconv.ParseInt(fmt.Sprint(v), 10, 64)

		n.Enum = append(n.Enum, v)
	default:
		panic(fmt.Sprintf("unsupported type: %s", reflect.TypeOf(v)))
	}
}

func (n *Number) SetName(name string) {
	n.Name = name
}

type String struct {
	Name string

	Enum []string
}

func (n *String) AddCandidates(v interface{}) {
	switch v := v.(type) {
	case string:
		n.Enum = append(n.Enum, v)
	default:
		panic(fmt.Sprintf("unsupported type: %s", reflect.TypeOf(v)))
	}
}

func (n *String) SetName(name string) {
	n.Name = name
}

type Boolean struct {
}

type Nullable struct {
	Inner Type
}

type Any struct{}
