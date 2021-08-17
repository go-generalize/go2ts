package main

import (
	tm "time"

	"github.com/go-generalize/go2ts/pkg/parser/testdata/success/base/pkg"
)

type Embedded struct {
	Foo int `json:"foo,omitempty"`
}

type Status string

const (
	OK      Status = "OK"
	Failure Status = "Failure"
)

type EmbeddedInt int

type Data struct {
	Time    tm.Time
	Package *pkg.Package

	Embedded
	EmbeddedInt

	A int
	B *int `json:"b,omitempty"`
	C string
	D *float32

	Array []int

	Status Status
	Map    map[string]Status

	Foo *foo `json:",omitempty"`

	U Unexported

	Hidden int `json:"-"`
}

type foo struct {
	V int
}

type Unexported struct {
	Data int
}
