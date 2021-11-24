// Package parser provides a Go module parser for TypeScript AST
package parser

import "github.com/go-generalize/go-easyparser"

// FilterOpt is options to filter exported interfaces
type FilterOpt = easyparser.FilterOpt

// Default exports all other than neither in base-package nor exported
var Default = easyparser.Default

// All exports all interfaces
var All = easyparser.All
