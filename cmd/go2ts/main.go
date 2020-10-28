// Package main is a CLI for go2ts
package main

import (
	"fmt"
	"os"

	"github.com/go-generalize/go2ts/pkg/generator"
	"github.com/go-generalize/go2ts/pkg/parser"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage:")
		fmt.Printf("%s [the target directory]\n", os.Args[0])

		os.Exit(1)
	}

	parser, err := parser.NewParser(os.Args[1], parser.Default)

	if err != nil {
		panic(err)
	}

	types, err := parser.Parse()

	if err != nil {
		panic(err)
	}

	fmt.Println(generator.NewGenerator(types).Generate())
}
