// Package parser provides a Go module parser for TypeScript AST
package parser

import (
	"errors"
	"strings"

	"golang.org/x/tools/go/packages"
)

func visitErrors(pkgs []*packages.Package) error {
	errs := []string{}
	packages.Visit(pkgs, nil, func(pkg *packages.Package) {
		for i := range pkg.Errors {
			errs = append(errs, pkg.Errors[i].Error())
		}
	})

	if len(errs) == 0 {
		return nil
	}

	return errors.New(strings.Join(errs, "\n"))
}
