// Package parser provides a Go module parser for TypeScript AST
package parser

// FilterOpt is options to filter exported interfaces
type FilterOpt struct {
	BasePackage bool
	Package     string
	Name        string
	Exported    bool
	// Dependency - 他のstructから依存されている場合にtrueとなる
	// 同じstructに対して複数回呼ばれ、依存されていない状況ではfalseとして呼ばれる可能性がある
	// 一度でもtrueとして返せば出力される
	// dependencyがfalseの時にtrueを返す場合、trueでも常にtrueを返すべきである
	Dependency bool
}

// Default exports all other than neither in base-package nor exported
var Default = func(opt *FilterOpt) bool {
	if !opt.BasePackage {
		return false
	}

	if !opt.Exported {
		return false
	}

	return true
}

// All exports all interfaces
var All = func(opt *FilterOpt) bool {
	return true
}
