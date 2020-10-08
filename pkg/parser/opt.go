package parser

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

var Default = func(opt *FilterOpt) bool {
	if !opt.BasePackage {
		return false
	}

	if !opt.Exported {
		return false
	}

	return true
}

var All = func(opt *FilterOpt) bool {
	return true
}
