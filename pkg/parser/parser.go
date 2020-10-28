// Package parser provides a Go module parser for TypeScript AST
package parser

import (
	"fmt"
	"go/types"
	"path/filepath"
	"reflect"

	tstypes "github.com/go-generalize/go2ts/pkg/types"
	"github.com/go-generalize/go2ts/pkg/util"
	"golang.org/x/tools/go/packages"
)

// Parser is a Go module parser for TypeScript AST
type Parser struct {
	pkgs []*packages.Package

	types       map[string]tstypes.Type
	consts      map[string][]interface{}
	basePackage string

	Filter func(opt *FilterOpt) bool
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

// NewParser initializes a new Parser
func NewParser(dir string, filter func(*FilterOpt) bool) (*Parser, error) {
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
		pkgs:        pkgs,
		basePackage: pkg,
		Filter:      filter,
	}, nil
}

func (p *Parser) exported(t *types.Named, dep bool) bool {
	opt := &FilterOpt{
		BasePackage: false,
		Package:     t.Obj().Pkg().String(),
		Name:        t.Obj().Name(),
		Exported:    t.Obj().Exported(),
		Dependency:  dep,
	}
	packages.Visit(p.pkgs, nil, func(pkg *packages.Package) {
		if pkg.Types.Scope() == t.Obj().Parent() {
			opt.BasePackage = true
		}
	})

	return p.Filter(opt)
}

func (p *Parser) parseNamed(t *types.Named, dep bool) tstypes.Type {
	if t.String() == "time.Time" {
		return &tstypes.Date{}
	}

	exported := p.exported(t, dep)

	if exported {
		tt, ok := p.types[t.String()]

		if ok {
			return tt
		}
	} else if !dep {
		return nil
	}

	typ := p.parseType(t.Underlying())

	if exported {
		if typ, ok := typ.(tstypes.Enumerable); ok {
			consts := p.consts[t.String()]

			for i := range consts {
				typ.AddCandidates(consts[i])
			}
		}

		if named, ok := typ.(tstypes.NamedType); ok {
			named.SetName(t.String())
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

func (p *Parser) parseMap(u *types.Map) tstypes.Type {
	keyType := p.parseType(u.Key())

	if !keyType.UsedAsMapKey() {
		panic(keyType.String() + " cannot be used as key")
	}

	return &tstypes.Map{
		Key:   keyType,
		Value: p.parseType(u.Elem()),
	}
}

func (p *Parser) parseInterface(u *types.Interface) tstypes.Type {
	return &tstypes.Any{}
}

func (p *Parser) parseType(u types.Type) tstypes.Type {
	var typ tstypes.Type
	switch u := u.(type) {
	case *types.Named:
		typ = p.parseNamed(u, true)
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
	case *types.Map:
		typ = p.parseMap(u)
	case *types.Interface:
		typ = p.parseInterface(u)
	default:
		panic("unsupported named type: " + reflect.TypeOf(u).String())
	}

	return typ
}

// Parse parses the Go module and returns ASTs
func (p *Parser) Parse() (res map[string]tstypes.Type, err error) {
	defer func() {
		if e := recover(); e != nil {
			var ok bool
			err, ok = e.(error)

			if !ok {
				err = fmt.Errorf("%+v", e)
			}
		}
	}()

	p.types = make(map[string]tstypes.Type)
	p.consts = make(map[string][]interface{})

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

	// parse types
	packages.Visit(p.pkgs, nil, func(pkg *packages.Package) {
		for _, obj := range pkg.TypesInfo.Defs {
			if obj == nil {
				continue
			}

			if obj.Parent() != pkg.Types.Scope() {
				continue
			}

			v, ok := obj.(*types.TypeName)
			if !(ok && !v.IsAlias()) {
				continue
			}

			t, ok := v.Type().(*types.Named)

			if !ok {
				continue
			}

			p.parseNamed(t, false)
		}
	})

	p.sortConst()

	return p.types, nil
}

// GetBasePackage returns a base module for the root package
func (p *Parser) GetBasePackage() string {
	return p.basePackage
}
