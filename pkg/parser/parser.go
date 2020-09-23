package parser

import (
	"go/types"

	tstypes "github.com/go-generalize/go2ts/pkg/types"
	"golang.org/x/tools/go/packages"
)

type Parser struct {
	pkgs []*packages.Package

	types map[string]tstypes.Type
}

func NewParser(dir string) (*Parser, error) {
	cfg := &packages.Config{
		Mode: packages.NeedCompiledGoFiles | packages.NeedSyntax | packages.NeedTypes | packages.NeedTypesInfo,
	}

	pkgs, err := packages.Load(cfg, dir)

	if err != nil {
		return nil, err
	}

	if err := visitErrors(pkgs); err != nil {
		return nil, err
	}

	return &Parser{
		pkgs: pkgs,
	}, nil
}

func (p *Parser) exported(t *types.Named) bool {
	if !t.Obj().Exported() {
		return false
	}

	flag := false
	packages.Visit(p.pkgs, nil, func(pkg *packages.Package) {
		if pkg.Types.Scope() == t.Obj().Parent() {
			flag = true
		}
	})

	return flag
}

func (p *Parser) parseNamed(t *types.Named) tstypes.Type {
	exported := p.exported(t)

	if exported {
		tt, ok := p.types[t.String()]

		if ok {
			return tt
		}
	}

	typ := p.parseType(t.Underlying())

	if exported {
		if named, ok := typ.(tstypes.NamedType); ok {
			named.SetName(t.Obj().Name())
		}

		p.types[t.String()] = typ
	}

	return typ
}

func (p *Parser) parsePointer(u *types.Pointer) tstypes.Type {
	return &tstypes.Nullable{
		Inner: p.parseType(u.Elem()),
	}
}

func (p *Parser) parseType(u types.Type) tstypes.Type {
	var typ tstypes.Type
	switch u := u.(type) {
	case *types.Named:
		typ = p.parseNamed(u)
	case *types.Struct:
		typ = p.parseStruct(u)
	case *types.Basic:
		typ = p.parseBasic(u)
	case *types.Pointer:
		typ = p.parsePointer(u)
	default:
		panic("unsupported named type")
	}

	return typ
}

func (p *Parser) Parse() (map[string]tstypes.Type, error) {
	p.types = make(map[string]tstypes.Type)

	// parse types
	packages.Visit(p.pkgs, nil, func(pkg *packages.Package) {
		for _, obj := range pkg.TypesInfo.Defs {
			if obj == nil {
				continue
			}

			if obj.Parent() != pkg.Types.Scope() {
				continue
			}

			if v, ok := obj.(*types.TypeName); ok && !v.IsAlias() {
				p.parseType(v.Type())
			}
		}
	})

	// parse const
	packages.Visit(p.pkgs, nil, func(pkg *packages.Package) {
		for _, obj := range pkg.TypesInfo.Defs {
			if obj == nil {
				continue
			}

			if obj.Parent() != pkg.Types.Scope() {
				continue
			}

			switch v := obj.(type) {
			case *types.Const: // const a = 1
				p.parseConst(v)
			}

		}
	})

	return p.types, nil
}
