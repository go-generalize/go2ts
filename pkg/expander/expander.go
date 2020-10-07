package expander

import (
	"strings"

	tstypes "github.com/go-generalize/go2ts/pkg/types"
)

type expander struct {
	data map[string]tstypes.Type

	basePackage string
	rootMode    bool
}

func (e *expander) expandType(typ tstypes.Type) tstypes.Type {
	switch v := typ.(type) {
	case *tstypes.Array:
		return e.expandArray(v)
	case *tstypes.Object:
		return e.expandObject(v)
	case *tstypes.String:
		return e.expandString(v)
	case *tstypes.Number:
		return e.expandNumber(v)
	case *tstypes.Boolean:
		return e.expandBoolean(v)
	case *tstypes.Date:
		return e.expandDate(v)
	case *tstypes.Nullable:
		return e.expandNullable(v)
	case *tstypes.Any:
		return e.expandAny(v)
	case *tstypes.Map:
		return e.expandMap(v)
	default:
		panic("unsupported")
	}
}

func ExpandUnexported(data map[string]tstypes.Type) map[string]tstypes.Type {
	e := expander{
		data: data,
	}

	if len(data) == 0 {
		return data
	}

	for _, v := range data {
		if v, ok := v.(*tstypes.Object); ok {
			s := strings.Split(v.Package, ".")

			if len(s) == 0 {
				continue
			}

			e.basePackage = strings.Join(s[:len(s)-1], ".")
			break
		}
	}

	for k := range data {
		data[k] = e.expandType(data[k])
	}

	return data
}
