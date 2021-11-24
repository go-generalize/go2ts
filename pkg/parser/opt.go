// Package parser provides a Go module parser for TypeScript AST
package parser

import "github.com/go-generalize/go-easyparser"

// FilterOpt is options to filter exported interfaces
// Deprecated: github.com/go-generalize/go-easyparser.FilterOpt
type FilterOpt = easyparser.FilterOpt

// Default exports all other than neither in base-package nor exported
// Deprecated: github.com/go-generalize/go-easyparser.Default
var Default = easyparser.Default

// All exports all interfaces
// Deprecated: github.com/go-generalize/go-easyparser.All
var All = easyparser.All
