package generator

import (
	"bytes"
	"fmt"
	"sort"
	"strings"

	tstypes "github.com/go-generalize/go2ts/pkg/types"
	"github.com/go-generalize/go2ts/pkg/util"
	"github.com/iancoleman/strcase"
)

type Generator struct {
	types   map[string]tstypes.Type
	altPkgs map[string]string

	BasePackage string
}

func NewGenerator(types map[string]tstypes.Type) *Generator {
	return &Generator{
		types:   types,
		altPkgs: map[string]string{},
	}
}

func (g *Generator) indent(s string) string {
	arr := strings.Split(s, "\n")

	if len(arr) > 1 {
		a := arr[1:]
		for i := range a {
			a[i] = "\t" + a[i]
		}
	}

	return strings.Join(arr, "\n")
}

func (g *Generator) generateType(t tstypes.Type) string {
	switch v := t.(type) {
	case *tstypes.Array:
		return g.generateArray(v)
	case *tstypes.Object:
		return g.generateObject(v, false)
	case *tstypes.String:
		return g.generateString(v)
	case *tstypes.Number:
		return g.generateNumber(v)
	case *tstypes.Boolean:
		return g.generateBoolean(v)
	case *tstypes.Date:
		return g.generateDate(v)
	case *tstypes.Nullable:
		return g.generateNullable(v)
	case *tstypes.Any:
		return g.generateAny(v)
	case *tstypes.Map:
		return g.generateMap(v)
	default:
		panic("unsupported")
	}
}

func (g *Generator) Generate() string {
	buf := bytes.NewBuffer(nil)

	type typesType struct {
	}

	type entry struct {
		key string
		typ tstypes.Type
	}

	entries := make([]*entry, 0, len(g.types))
	for k, v := range g.types {
		entries = append(entries, &entry{
			key: k,
			typ: v,
		})
	}

	sort.Slice(entries, func(i, j int) bool {
		left := strings.HasPrefix(entries[i].key, g.BasePackage+".")
		right := strings.HasPrefix(entries[j].key, g.BasePackage+".")

		if left && !right {
			return true
		}
		if !left && right {
			return false
		}

		return entries[i].key < entries[j].key
	})

	used := map[string]struct{}{}
	for i, e := range entries {
		obj, ok := e.typ.(*tstypes.Object)

		if !ok {
			continue
		}

		pkg, strct := util.SplitPackegeStruct(obj.Name)
		if _, ok := used[strct]; !ok {
			g.altPkgs[obj.Name] = strct
			used[strct] = struct{}{}
			continue
		}

		p := util.GetPackageNameFromPath(pkg)
		name := strcase.ToCamel(p + "_" + strct)

		if _, ok := used[name]; !ok {
			g.altPkgs[obj.Name] = name
			used[name] = struct{}{}
			continue
		}

		name = fmt.Sprintf("%s%03d", name, i)
		g.altPkgs[obj.Name] = name
		used[name] = struct{}{}
	}

	for _, e := range entries {
		obj, ok := e.typ.(*tstypes.Object)

		if !ok {
			continue
		}

		buf.WriteString("export type " + g.altPkgs[obj.Name] + " = " + g.generateObject(obj, true) + "\n")
	}

	return buf.String()
}
