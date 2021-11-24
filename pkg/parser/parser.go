// Package parser provides a Go module parser for TypeScript AST
package parser

import (
	"github.com/go-generalize/go-easyparser"
)

// Parser is a Go module parser for TypeScript AST
// Deprecated: github.com/go-generalize/go-easyparser.Parser
type Parser = easyparser.Parser

// NewParser initializes a new Parser
// Deprecated: github.com/go-generalize/go-easyparser.NewParser
func NewParser(dir string, filter func(*FilterOpt) bool) (*Parser, error) {
	return easyparser.NewParser(dir, filter)
}
