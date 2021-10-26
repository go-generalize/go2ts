package conflict

import (
	"go/types"

	tstypes "github.com/go-generalize/go2ts/pkg/types"
)

var Type = map[string]tstypes.Type{
	"github.com/go-generalize/go2ts/pkg/parser/testdata/conflict/base.Data": &tstypes.Object{
		Common: tstypes.Common{PkgName: "main"},
		Name:   "github.com/go-generalize/go2ts/pkg/parser/testdata/conflict/base.Data",
		Entries: map[string]tstypes.ObjectEntry{
			"Hoge": {
				RawName: "Hoge",
				Type: &tstypes.Object{
					Common: tstypes.Common{PkgName: "main"},
					Name:   "github.com/go-generalize/go2ts/pkg/parser/testdata/conflict/base.Hoge",
					Entries: map[string]tstypes.ObjectEntry{
						"Data": {
							RawName: "Data",
							Type: &tstypes.Number{
								RawType: types.Int,
							},
						},
					},
				},
				FieldIndex: 0,
			},
			"PkgHoge": {
				RawName: "PkgHoge",
				Type: &tstypes.Object{
					Name: "github.com/go-generalize/go2ts/pkg/parser/testdata/conflict/base/pkg.Hoge",
					Entries: map[string]tstypes.ObjectEntry{
						"Data": {
							RawName: "Data",
							Type: &tstypes.Number{
								RawType: types.Int,
							},
						},
					},
				},
				FieldIndex: 1,
			},
		},
	},
	"github.com/go-generalize/go2ts/pkg/parser/testdata/conflict/base.Hoge": &tstypes.Object{
		Common: tstypes.Common{PkgName: "main"},
		Name:   "github.com/go-generalize/go2ts/pkg/parser/testdata/conflict/base.Hoge",
		Entries: map[string]tstypes.ObjectEntry{
			"Data": {
				RawName: "Data",
				Type: &tstypes.Number{
					RawType: types.Int,
				},
			},
		},
	},
	"github.com/go-generalize/go2ts/pkg/parser/testdata/conflict/base/pkg.Hoge": &tstypes.Object{
		Name: "github.com/go-generalize/go2ts/pkg/parser/testdata/conflict/base/pkg.Hoge",
		Entries: map[string]tstypes.ObjectEntry{
			"Data": {
				RawName: "Data",
				Type: &tstypes.Number{
					RawType: types.Int,
				},
			},
		},
	},
}
