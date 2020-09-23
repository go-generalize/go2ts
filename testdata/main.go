package main

type Embedded struct {
	Foo int `json:"foo,omitempty"`
}

type Status string

const (
	OK      Status = "OK"
	Failure Status = "Failure"
)

type Data struct {
	Embedded `json:"embedded"`

	A int
	B int `json:"b"`
	C string
	D float32

	Status Status
}
