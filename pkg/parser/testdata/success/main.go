package main

import (
	tm "time"

	"github.com/go-generalize/go2ts/pkg/parser/testdata/success/pkg"
)

type Embedded struct {
	Foo int `json:"foo,omitempty"`
}

type Status string

const (
	OK      Status = "OK"
	Failure Status = "Failure"
)

type Data struct {
	Time    tm.Time
	Package *pkg.Package

	Embedded

	A int
	B *int `json:"b,omitempty"`
	C string
	D *float32

	Array []int

	Status Status

	Foo *foo `json:",omitempty"`

	U Unexported
}

type foo struct {
	V int
}

type Unexported struct {
	Data int
}
