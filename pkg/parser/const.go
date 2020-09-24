package parser

import (
	"go/constant"
	"go/types"
	"sort"
	"strconv"

	tstypes "github.com/go-generalize/go2ts/pkg/types"
)

func (p *Parser) parseConst(c *types.Const) {
	if !c.Exported() {
		return
	}

	typ, ok := p.types[c.Type().String()]

	if !ok {
		return
	}

	t, ok := typ.(tstypes.Enumerable)

	if !ok {
		return
	}

	switch c.Val().Kind() {
	case constant.Int:
		v, err := strconv.ParseInt(c.Val().ExactString(), 10, 64)

		if err != nil {
			panic(err)
		}
		t.AddCandidates(v)
	case constant.Float:
		v, err := strconv.ParseFloat(c.Val().ExactString(), 64)

		if err != nil {
			panic(err)
		}
		t.AddCandidates(v)
	case constant.String:
		unquoted, err := strconv.Unquote(c.Val().ExactString())

		if err != nil {
			panic(err)
		}

		t.AddCandidates(unquoted)
	default:
		panic("unsupported enum type")
	}
}

func (p *Parser) sortConst() {
	for _, v := range p.types {
		switch v := v.(type) {
		case *tstypes.String:
			sort.Strings(v.Enum)
		case *tstypes.Number:
			sort.Slice(v.Enum, func(i int, j int) bool {
				return v.Enum[i] < v.Enum[j]
			})
		}
	}
}
