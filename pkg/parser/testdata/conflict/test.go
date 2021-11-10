package conflict

import (
	"go/types"

	"github.com/go-generalize/go2ts/pkg/parser/testutil"
	tstypes "github.com/go-generalize/go2ts/pkg/types"
)

var Type = map[string]tstypes.Type{
	"github.com/go-generalize/go2ts/pkg/parser/testdata/conflict/base.Data": &tstypes.Object{
		Common: tstypes.Common{
			PkgName:  "main",
			Position: testutil.ParsePositionString("testdata/conflict/base/main.go:7:6"),
		},
		Name: "github.com/go-generalize/go2ts/pkg/parser/testdata/conflict/base.Data",
		Entries: map[string]tstypes.ObjectEntry{
			"Hoge": {
				RawName: "Hoge",
				Type: &tstypes.Object{
					Common: tstypes.Common{
						PkgName:  "main",
						Position: testutil.ParsePositionString("testdata/conflict/base/main.go:12:6"),
					},
					Name: "github.com/go-generalize/go2ts/pkg/parser/testdata/conflict/base.Hoge",
					Entries: map[string]tstypes.ObjectEntry{
						"Data": {
							RawName: "Data",
							Type: &tstypes.Number{
								RawType: types.Int,
							},
							Position: testutil.ParsePositionString("testdata/conflict/base/main.go:13:2"),
						},
					},
				},
				FieldIndex: 0,
				Position:   testutil.ParsePositionString("testdata/conflict/base/main.go:8:2"),
			},
			"PkgHoge": {
				RawName: "PkgHoge",
				Type: &tstypes.Object{
					Common: tstypes.Common{
						Position: testutil.ParsePositionString("testdata/conflict/base/pkg/pkg.go:3:1"),
					},
					Name: "github.com/go-generalize/go2ts/pkg/parser/testdata/conflict/base/pkg.Hoge",
					Entries: map[string]tstypes.ObjectEntry{
						"Data": {
							RawName: "Data",
							Type: &tstypes.Number{
								RawType: types.Int,
							},
							Position: testutil.ParsePositionString("testdata/conflict/base/pkg/pkg.go:4:1"),
						},
					},
				},
				FieldIndex: 1,
				Position:   testutil.ParsePositionString("testdata/conflict/base/main.go:9:2"),
			},
		},
	},
	"github.com/go-generalize/go2ts/pkg/parser/testdata/conflict/base.Hoge": &tstypes.Object{
		Common: tstypes.Common{
			PkgName:  "main",
			Position: testutil.ParsePositionString("testdata/conflict/base/main.go:12:6"),
		},
		Name: "github.com/go-generalize/go2ts/pkg/parser/testdata/conflict/base.Hoge",
		Entries: map[string]tstypes.ObjectEntry{
			"Data": {
				RawName: "Data",
				Type: &tstypes.Number{
					RawType: types.Int,
				},
				Position: testutil.ParsePositionString("testdata/conflict/base/main.go:13:2"),
			},
		},
	},
	"github.com/go-generalize/go2ts/pkg/parser/testdata/conflict/base/pkg.Hoge": &tstypes.Object{
		Common: tstypes.Common{
			Position: testutil.ParsePositionString("testdata/conflict/base/pkg/pkg.go:3:1"),
		},
		Name: "github.com/go-generalize/go2ts/pkg/parser/testdata/conflict/base/pkg.Hoge",
		Entries: map[string]tstypes.ObjectEntry{
			"Data": {
				RawName: "Data",
				Type: &tstypes.Number{
					RawType: types.Int,
				},
				Position: testutil.ParsePositionString("testdata/conflict/base/pkg/pkg.go:4:1"),
			},
		},
	},
}
