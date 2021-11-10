package success

import (
	"go/types"

	"github.com/go-generalize/go2ts/pkg/parser/testutil"
	tstypes "github.com/go-generalize/go2ts/pkg/types"
)

var NormalType = map[string]tstypes.Type{
	"github.com/go-generalize/go2ts/pkg/parser/testdata/success/base.Embedded": &tstypes.Object{
		Common: tstypes.Common{
			PkgName:  "main",
			Position: testutil.ParsePositionString("testdata/success/base/main.go:9:6"),
		},
		Name: "github.com/go-generalize/go2ts/pkg/parser/testdata/success/base.Embedded",
		Entries: map[string]tstypes.ObjectEntry{
			"foo": {
				RawName:  "Foo",
				RawTag:   `json:"foo,omitempty"`,
				Optional: true,
				Type: &tstypes.Number{
					RawType: types.Int,
				},
				FieldIndex: 0,
				Position:   testutil.ParsePositionString("testdata/success/base/main.go:10:2"),
			},
		},
	},
	"github.com/go-generalize/go2ts/pkg/parser/testdata/success/base.Status": &tstypes.String{
		Common: tstypes.Common{
			PkgName:  "main",
			Position: testutil.ParsePositionString("testdata/success/base/main.go:13:6"),
		},
		Name: "github.com/go-generalize/go2ts/pkg/parser/testdata/success/base.Status",
		Enum: []string{"Failure", "OK"},
		RawEnum: []tstypes.RawStringEnumCandidate{
			{Key: "Failure", Value: "Failure"},
			{Key: "OK", Value: "OK"},
		},
	},
	"github.com/go-generalize/go2ts/pkg/parser/testdata/success/base.EmbeddedInt": &tstypes.Number{
		Common: tstypes.Common{
			PkgName:  "main",
			Position: testutil.ParsePositionString("testdata/success/base/main.go:20:6"),
		},
		Name:    "github.com/go-generalize/go2ts/pkg/parser/testdata/success/base.EmbeddedInt",
		RawType: types.Int,
	},
	"github.com/go-generalize/go2ts/pkg/parser/testdata/success/base.Data": &tstypes.Object{
		Common: tstypes.Common{
			PkgName:  "main",
			Position: testutil.ParsePositionString("testdata/success/base/main.go:22:6"),
		},
		Name: "github.com/go-generalize/go2ts/pkg/parser/testdata/success/base.Data",
		Entries: map[string]tstypes.ObjectEntry{
			"Time": {
				RawName:    "Time",
				Type:       &tstypes.Date{},
				FieldIndex: 0,
				Position:   testutil.ParsePositionString("testdata/success/base/main.go:23:2"),
			},
			"Package": {
				RawName: "Package",
				Type: &tstypes.Nullable{
					Inner: &tstypes.Object{
						Entries: map[string]tstypes.ObjectEntry{
							"data": {
								RawName: "Data",
								RawTag:  `json:"data"`,
								Type: &tstypes.Number{
									RawType: types.Int,
								},
								FieldIndex: 0,
								Position:   testutil.ParsePositionString("testdata/success/base/pkg/pkg.go:4:1"),
							},
						},
					},
				},
				FieldIndex: 1,
				Position:   testutil.ParsePositionString("testdata/success/base/main.go:24:2"),
			},
			"foo": {
				RawName:  "Foo",
				RawTag:   `json:"foo,omitempty"`,
				Optional: true,
				Type: &tstypes.Number{
					RawType: types.Int,
				},
				FieldIndex: 2,
				Position:   testutil.ParsePositionString("testdata/success/base/main.go:10:2"),
			},
			"EmbeddedInt": {
				RawName: "EmbeddedInt",
				Type: &tstypes.Number{
					Common: tstypes.Common{
						PkgName:  "main",
						Position: testutil.ParsePositionString("testdata/success/base/main.go:20:6"),
					},
					RawType: types.Int,
					Name:    "github.com/go-generalize/go2ts/pkg/parser/testdata/success/base.EmbeddedInt",
				},
				FieldIndex: 3,
			},
			"A": {
				RawName: "A",
				Type: &tstypes.Number{
					RawType: types.Int,
				},
				FieldIndex: 4,
				Position:   testutil.ParsePositionString("testdata/success/base/main.go:29:2"),
			},
			"b": {
				RawName:  "B",
				RawTag:   `json:"b,omitempty"`,
				Optional: true,
				Type: &tstypes.Number{
					RawType: types.Int,
				},
				FieldIndex: 5,
				Position:   testutil.ParsePositionString("testdata/success/base/main.go:30:2"),
			},
			"C": {
				RawName:    "C",
				Type:       &tstypes.String{},
				FieldIndex: 6,
				Position:   testutil.ParsePositionString("testdata/success/base/main.go:31:2"),
			},
			"D": {
				RawName: "D",
				Type: &tstypes.Nullable{
					Inner: &tstypes.Number{
						RawType: types.Float32,
					},
				},
				FieldIndex: 7,
				Position:   testutil.ParsePositionString("testdata/success/base/main.go:32:2"),
			},

			"Array": {
				RawName: "Array",
				Type: &tstypes.Nullable{
					Inner: &tstypes.Array{
						Inner: &tstypes.Number{
							RawType: types.Int,
						},
					},
				},
				FieldIndex: 8,
				Position:   testutil.ParsePositionString("testdata/success/base/main.go:34:2"),
			},

			"Status": {
				RawName: "Status",
				Type: &tstypes.String{
					Common: tstypes.Common{
						PkgName:  "main",
						Position: testutil.ParsePositionString("testdata/success/base/main.go:13:6"),
					},
					Name: "github.com/go-generalize/go2ts/pkg/parser/testdata/success/base.Status",
					Enum: []string{"Failure", "OK"},
					RawEnum: []tstypes.RawStringEnumCandidate{
						{Key: "Failure", Value: "Failure"},
						{Key: "OK", Value: "OK"},
					},
				},
				FieldIndex: 9,
				Position:   testutil.ParsePositionString("testdata/success/base/main.go:36:2"),
			},
			"Map": {
				RawName: "Map",
				Type: &tstypes.Map{
					Key: &tstypes.String{},
					Value: &tstypes.String{
						Common: tstypes.Common{
							PkgName:  "main",
							Position: testutil.ParsePositionString("testdata/success/base/main.go:13:6"),
						},
						Name: "github.com/go-generalize/go2ts/pkg/parser/testdata/success/base.Status",
						Enum: []string{"Failure", "OK"},
						RawEnum: []tstypes.RawStringEnumCandidate{
							{Key: "Failure", Value: "Failure"},
							{Key: "OK", Value: "OK"},
						},
					},
				},
				FieldIndex: 10,
				Position:   testutil.ParsePositionString("testdata/success/base/main.go:37:2"),
			},
			"Foo": {
				RawName:  "Foo",
				RawTag:   `json:",omitempty"`,
				Optional: true,
				Type: &tstypes.Object{
					Entries: map[string]tstypes.ObjectEntry{
						"V": {
							RawName: "V",
							Type: &tstypes.Number{
								RawType: types.Int,
							},
							Position: testutil.ParsePositionString("testdata/success/base/main.go:49:2"),
						},
					},
				},
				FieldIndex: 11,
				Position:   testutil.ParsePositionString("testdata/success/base/main.go:39:2"),
			},
			"U": {
				RawName: "U",
				Type: &tstypes.Object{
					Entries: map[string]tstypes.ObjectEntry{
						"Data": {
							RawName: "Data",
							Type: &tstypes.Number{
								RawType: types.Int,
							},
							Position: testutil.ParsePositionString("testdata/success/base/main.go:53:2"),
						},
					},
				},
				FieldIndex: 12,
				Position:   testutil.ParsePositionString("testdata/success/base/main.go:41:2"),
			},
			"ByteSlice": {
				RawName: "ByteSlice",
				Type: &tstypes.Nullable{
					Inner: &tstypes.String{},
				},
				FieldIndex: 13,
				Position:   testutil.ParsePositionString("testdata/success/base/main.go:45:2"),
			},
		},
	},
}
