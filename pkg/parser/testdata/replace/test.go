package replace

import (
	tstypes "github.com/go-generalize/go2ts/pkg/types"
)

var Type = map[string]tstypes.Type{
	"github.com/go-generalize/go2ts/pkg/parser/testdata/replace/base.Struct": &tstypes.Object{
		Common: tstypes.Common{PkgName: "main"},
		Name:   "github.com/go-generalize/go2ts/pkg/parser/testdata/replace/base.Struct",
		Entries: map[string]tstypes.ObjectEntry{
			"A": {
				RawName:    "A",
				FieldIndex: 0,

				Type:     &tstypes.Number{},
				Optional: false,
			},
		},
	},
}
