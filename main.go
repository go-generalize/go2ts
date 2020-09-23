package main

import (
	"github.com/go-generalize/go2ts/pkg/parser"
	"github.com/k0kubun/pp"
)

func main() {
	parser, err := parser.NewParser("./testdata")

	if err != nil {
		panic(err)
	}

	types, err := parser.Parse()

	if err != nil {
		panic(err)
	}

	pp.Println(types["github.com/go-generalize/go2ts/testdata.Data"])
}
