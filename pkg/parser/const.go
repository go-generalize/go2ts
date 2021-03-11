// Package parser provides a Go module parser for TypeScript AST
package parser

import (
	"go/constant"
	"go/types"
	"sort"
	"strconv"

	tstypes "github.com/go-generalize/go2ts/pkg/types"
)

type constCandidate struct {
	Key   string
	Value interface{}
}

func (p *Parser) addCandidates(typ, key string, val interface{}) {
	arr, ok := p.consts[typ]

	if !ok {
		arr = make([]constCandidate, 0, 10)
	}

	p.consts[typ] = append(arr, constCandidate{
		Key:   key,
		Value: val,
	})
}

func (p *Parser) parseConst(c *types.Const) {
	if !c.Exported() {
		return
	}

	key := c.Name()

	switch c.Val().Kind() {
	case constant.Int:
		v, err := strconv.ParseInt(c.Val().ExactString(), 10, 64)

		if err != nil {
			return
		}

		p.addCandidates(c.Type().String(), key, v)
	case constant.Float:
		v, err := strconv.ParseFloat(c.Val().ExactString(), 64)

		if err != nil {
			return
		}

		p.addCandidates(c.Type().String(), key, v)
	case constant.String:
		unquoted, err := strconv.Unquote(c.Val().ExactString())

		if err != nil {
			return
		}

		p.addCandidates(c.Type().String(), key, unquoted)
	}
}

func (p *Parser) sortConst() {
	for _, v := range p.types {
		switch v := v.(type) {
		case *tstypes.String:
			sort.Slice(v.RawEnum, func(i int, j int) bool {
				return v.RawEnum[i].Key < v.RawEnum[j].Key
			})
			sort.Strings(v.Enum)
		case *tstypes.Number:
			sort.Slice(v.RawEnum, func(i int, j int) bool {
				return v.RawEnum[i].Key < v.RawEnum[j].Key
			})
			sort.Slice(v.Enum, func(i int, j int) bool {
				return v.Enum[i] < v.Enum[j]
			})
		}
	}
}
