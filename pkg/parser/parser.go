package parser

import (
	"go/types"
	"path/filepath"
	"reflect"

	tstypes "github.com/go-generalize/go2ts/pkg/types"
	"github.com/go-generalize/go2ts/pkg/util"
	"golang.org/x/tools/go/packages"
)

type Parser struct {
	pkgs []*packages.Package

	types map[string]tstypes.Type
}

func getPackagePath(dir string) (root string, pkg string, err error) {
	goModPath, err := util.GetGoModPath(dir)

	if err != nil {
		return "", "", err
	}
	goModDir := filepath.Dir(goModPath)

	mod, err := util.GetGoModule(goModPath)

	if err != nil {
		return "", "", err
	}

	abs, err := filepath.Abs(dir)

	if err != nil {
		return "", "", err
	}

	rel, err := filepath.Rel(goModDir, abs)

	if err != nil {
		return "", "", err
	}

	return goModDir, filepath.Join(mod, "/"+rel), nil
}

func NewParser(dir string) (*Parser, error) {
	root, pkg, err := getPackagePath(dir)

	if err != nil {
		return nil, err
	}

	cfg := &packages.Config{
		Mode: packages.NeedCompiledGoFiles | packages.NeedSyntax | packages.NeedTypes | packages.NeedTypesInfo,
		Dir:  root,
	}

	pkgs, err := packages.Load(cfg, pkg)

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
	if t.String() == "time.Time" {
		return &tstypes.Date{}
	}

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

func (p *Parser) parseSlice(u *types.Slice) tstypes.Type {
	return &tstypes.Nullable{
		Inner: &tstypes.Array{
			Inner: p.parseType(u.Elem()),
		},
	}
}

func (p *Parser) parseArray(u *types.Array) tstypes.Type {
	return &tstypes.Array{
		Inner: p.parseType(u.Elem()),
	}
}

func (p *Parser) parseInterface(u *types.Interface) tstypes.Type {
	return &tstypes.Any{}
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
	case *types.Slice:
		typ = p.parseSlice(u)
	case *types.Array:
		typ = p.parseArray(u)
	case *types.Interface:
		typ = p.parseInterface(u)
	default:
		panic("unsupported named type: " + reflect.TypeOf(u).String())
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

	p.sortConst()

	return p.types, nil
}
