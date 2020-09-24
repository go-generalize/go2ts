package parser

import (
	"go/types"
	"reflect"
	"strings"

	tstypes "github.com/go-generalize/go2ts/pkg/types"
)

func (p *Parser) parseStruct(strct *types.Struct) tstypes.Type {
	obj := tstypes.Object{
		Entries: map[string]tstypes.ObjectEntry{},
	}

	// embedding
	for i := 0; i < strct.NumFields(); i++ {
		v := strct.Field(i)
		tag := strct.Tag(i)

		if !v.Exported() || !v.Embedded() {
			continue
		}

		tst := p.parseType(v.Type())

		jsonTag := strings.SplitN(reflect.StructTag(tag).Get("json"), ",", 2)
		field := ""
		if len(jsonTag) >= 1 {
			field = jsonTag[0]
		}
		optional := len(jsonTag) >= 2 && jsonTag[1] == "omitempty"

		if len(field) == 0 {
			if o, ok := tst.(*tstypes.Object); ok {
				for k, v := range o.Entries {
					obj.Entries[k] = v
				}

				continue
			}

			field = v.Name()
		}

		obj.Entries[field] = tstypes.ObjectEntry{
			Type:     tst,
			Optional: optional,
		}
	}

	// not embedding
	for i := 0; i < strct.NumFields(); i++ {
		v := strct.Field(i)
		tag := strct.Tag(i)

		if !v.Exported() || v.Embedded() {
			continue
		}

		tst := p.parseType(v.Type())

		jsonTag := strings.SplitN(reflect.StructTag(tag).Get("json"), ",", 2)
		field := ""
		if len(jsonTag) >= 1 {
			field = jsonTag[0]
		}
		optional := len(jsonTag) >= 2 && jsonTag[1] == "omitempty"

		if len(field) == 0 {
			field = v.Name()
		}

		if optional {
			obj.Entries[field] = tstypes.ObjectEntry{
				Type:     p.removeNullable(tst),
				Optional: true,
			}
		} else {
			obj.Entries[field] = tstypes.ObjectEntry{
				Type:     tst,
				Optional: false,
			}
		}

	}

	return &obj
}

func (p *Parser) removeNullable(typ tstypes.Type) tstypes.Type {
	if nullable, ok := typ.(*tstypes.Nullable); ok {
		return nullable.Inner
	}

	return typ
}
