package parser

import (
	"go/types"

	tstypes "github.com/go-generalize/go2ts/pkg/types"
)

func (p *Parser) parseBasic(t *types.Basic) tstypes.Type {
	switch {
	case t.Info()&(types.IsNumeric|types.IsFloat) != 0:
		return &tstypes.Number{}
	case t.Info()&types.IsBoolean != 0:
		return &tstypes.Boolean{}
	case t.Info()&types.IsString != 0:
		return &tstypes.String{}
	default:
		panic("unsupported type")
	}
}
