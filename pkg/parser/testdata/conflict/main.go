package main

import (
	"github.com/go-generalize/go2ts/v2/pkg/parser/testdata/conflict/pkg"
)

type Data struct {
	Hoge    Hoge
	PkgHoge pkg.Hoge
}

type Hoge struct {
	Data int
}
