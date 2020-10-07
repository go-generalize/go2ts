package expander

import (
	"strings"

	tstypes "github.com/go-generalize/go2ts/pkg/types"
)

func (g *expander) expandArray(typ *tstypes.Array) *tstypes.Array {
	typ.Inner = g.expandType(typ.Inner)

	return typ
}

func (g *expander) expandMap(obj *tstypes.Map) *tstypes.Map {
	obj.Key = g.expandType(obj.Key)
	obj.Value = g.expandType(obj.Value)

	return obj
}

func (g *expander) expandObject(obj *tstypes.Object) *tstypes.Object {
	for _, v := range obj.Entries {
		v.Type = g.expandType(v.Type)
	}

	if obj.Name != "" {
		return obj
	}

	if strings.HasPrefix(obj.Package, g.basePackage+".") {
		if !g.rootMode {
			return obj
		}

		obj.Name = strings.TrimPrefix(obj.Package, g.basePackage+".")
		g.data[obj.Name] = obj
	} else {
		if g.rootMode {
			return obj
		}

		idx := strings.LastIndex(obj.Package, ".")

		if idx == -1 {
			return obj
		}

		name := obj.Package[idx+1:]

		obj.Name = name
	}

	return obj
}
