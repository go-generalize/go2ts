// Package generator is TypeScript generator from AST
package generator

import (
	"bytes"
	"fmt"
	"sort"

	tstypes "github.com/go-generalize/go2ts/pkg/types"
)

func (g *Generator) generateArray(typ *tstypes.Array) string {
	inner, union := g.generateType(typ.Inner)

	if union {
		return fmt.Sprintf("(%s)[]", inner)
	}

	return inner + "[]"
}

func (g *Generator) generateMap(obj *tstypes.Map) string {
	return fmt.Sprintf(
		"{[key: %s]: %s}",
		g.generateTypeSimple(obj.Key),
		g.generateTypeSimple(obj.Value),
	)
}

func (g *Generator) generateObject(obj *tstypes.Object, root bool) string {
	if !root && obj.Name != "" {
		return g.altPkgs[obj.Name]
	}

	type entry struct {
		Name     string
		Type     tstypes.Type
		Optional bool
	}

	entries := make([]*entry, 0, len(obj.Entries))

	for k, v := range obj.Entries {
		entries = append(entries, &entry{
			Name:     k,
			Type:     v.Type,
			Optional: v.Optional,
		})
	}

	sort.Slice(entries, func(i, j int) bool {
		return entries[i].Name < entries[j].Name
	})

	buf := bytes.NewBuffer(nil)

	if len(entries) == 0 {
		return "{}"
	}

	buf.WriteString("{\n")

	for _, e := range entries {
		if e.Optional {
			buf.WriteString(fmt.Sprintf("\t%s?: %s;\n", e.Name, g.indent(g.generateTypeSimple(e.Type))))
		} else {
			buf.WriteString(fmt.Sprintf("\t%s: %s;\n", e.Name, g.indent(g.generateTypeSimple(e.Type))))
		}
	}

	buf.WriteString("}")

	return buf.String()
}
