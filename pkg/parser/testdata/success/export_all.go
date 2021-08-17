package success

import (
	"go/types"

	tstypes "github.com/go-generalize/go2ts/pkg/types"
)

var AllExportedType = map[string]tstypes.Type{
	"github.com/go-generalize/go2ts/pkg/parser/testdata/success/base.Embedded": &tstypes.Object{
		Name: "github.com/go-generalize/go2ts/pkg/parser/testdata/success/base.Embedded",
		Entries: map[string]tstypes.ObjectEntry{
			"foo": {
				RawName:  "Foo",
				RawTag:   `json:"foo,omitempty"`,
				Optional: true,
				Type: &tstypes.Number{
					RawType: types.Int,
				},
			},
		},
	},
	"github.com/go-generalize/go2ts/pkg/parser/testdata/success/base.Status": &tstypes.String{
		Name: "github.com/go-generalize/go2ts/pkg/parser/testdata/success/base.Status",
		Enum: []string{"Failure", "OK"},
		RawEnum: []tstypes.RawStringEnumCandidate{
			{Key: "Failure", Value: "Failure"},
			{Key: "OK", Value: "OK"},
		},
	},
	"github.com/go-generalize/go2ts/pkg/parser/testdata/success/base.foo": &tstypes.Object{
		Name: "github.com/go-generalize/go2ts/pkg/parser/testdata/success/base.foo",
		Entries: map[string]tstypes.ObjectEntry{
			"V": {
				RawName: "V",
				Type: &tstypes.Number{
					RawType: types.Int,
				},
			},
		},
	},
	"github.com/go-generalize/go2ts/pkg/parser/testdata/success/base/pkg.Package": &tstypes.Object{
		Name: "github.com/go-generalize/go2ts/pkg/parser/testdata/success/base/pkg.Package",
		Entries: map[string]tstypes.ObjectEntry{
			"data": {
				RawName: "Data",
				RawTag:  `json:"data"`,
				Type: &tstypes.Number{
					RawType: types.Int,
				},
			},
		},
	},
	"github.com/go-generalize/go2ts/pkg/parser/testdata/success/base.Unexported": &tstypes.Object{
		Name: "github.com/go-generalize/go2ts/pkg/parser/testdata/success/base.Unexported",
		Entries: map[string]tstypes.ObjectEntry{
			"Data": {
				RawName: "Data",
				Type: &tstypes.Number{
					RawType: types.Int,
				},
			},
		},
	},
	"github.com/go-generalize/go2ts/pkg/parser/testdata/success/base.EmbeddedInt": &tstypes.Number{
		Name:    "github.com/go-generalize/go2ts/pkg/parser/testdata/success/base.EmbeddedInt",
		RawType: types.Int,
	},
	"github.com/go-generalize/go2ts/pkg/parser/testdata/success/base.Data": &tstypes.Object{
		Name: "github.com/go-generalize/go2ts/pkg/parser/testdata/success/base.Data",
		Entries: map[string]tstypes.ObjectEntry{
			"Time": {
				RawName:    "Time",
				Type:       &tstypes.Date{},
				FieldIndex: 0,
			},
			"Package": {
				RawName: "Package",
				Type: &tstypes.Nullable{
					Inner: &tstypes.Object{
						Name: "github.com/go-generalize/go2ts/pkg/parser/testdata/success/base/pkg.Package",
						Entries: map[string]tstypes.ObjectEntry{
							"data": {
								RawName: "Data",
								RawTag:  `json:"data"`,
								Type: &tstypes.Number{
									RawType: types.Int,
								},
							},
						},
					},
				},
				FieldIndex: 1,
			},
			"foo": {
				RawName:  "Foo",
				RawTag:   `json:"foo,omitempty"`,
				Optional: true,
				Type: &tstypes.Number{
					RawType: types.Int,
				},
				FieldIndex: 2,
			},
			"EmbeddedInt": {
				RawName: "EmbeddedInt",
				Type: &tstypes.Number{
					Name:    "github.com/go-generalize/go2ts/pkg/parser/testdata/success/base.EmbeddedInt",
					RawType: types.Int,
				},
				FieldIndex: 3,
			},
			"A": {
				RawName: "A",
				Type: &tstypes.Number{
					RawType: types.Int,
				},
				FieldIndex: 4,
			},
			"b": {
				RawName:  "B",
				RawTag:   `json:"b,omitempty"`,
				Optional: true,
				Type: &tstypes.Number{
					RawType: types.Int,
				},
				FieldIndex: 5,
			},
			"C": {
				RawName:    "C",
				Type:       &tstypes.String{},
				FieldIndex: 6,
			},
			"D": {
				RawName: "D",
				Type: &tstypes.Nullable{
					Inner: &tstypes.Number{
						RawType: types.Float32,
					},
				},
				FieldIndex: 7,
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
			},
			"Status": {
				RawName: "Status",
				Type: &tstypes.String{
					Name: "github.com/go-generalize/go2ts/pkg/parser/testdata/success/base.Status",
					Enum: []string{"Failure", "OK"},
					RawEnum: []tstypes.RawStringEnumCandidate{
						{Key: "Failure", Value: "Failure"},
						{Key: "OK", Value: "OK"},
					},
				},
				FieldIndex: 9,
			},
			"Map": {
				RawName: "Map",
				Type: &tstypes.Map{
					Key: &tstypes.String{},
					Value: &tstypes.String{
						Name: "github.com/go-generalize/go2ts/pkg/parser/testdata/success/base.Status",
						Enum: []string{"Failure", "OK"},
						RawEnum: []tstypes.RawStringEnumCandidate{
							{Key: "Failure", Value: "Failure"},
							{Key: "OK", Value: "OK"},
						},
					},
				},
				FieldIndex: 10,
			},
			"Foo": {
				RawName:  "Foo",
				RawTag:   `json:",omitempty"`,
				Optional: true,
				Type: &tstypes.Object{
					Name: "github.com/go-generalize/go2ts/pkg/parser/testdata/success/base.foo",
					Entries: map[string]tstypes.ObjectEntry{
						"V": {
							RawName: "V",
							Type: &tstypes.Number{
								RawType: types.Int,
							},
						},
					},
				},
				FieldIndex: 11,
			},
			"U": {
				RawName: "U",
				Type: &tstypes.Object{
					Name: "github.com/go-generalize/go2ts/pkg/parser/testdata/success/base.Unexported",
					Entries: map[string]tstypes.ObjectEntry{
						"Data": {
							RawName: "Data",
							Type: &tstypes.Number{
								RawType: types.Int,
							},
						},
					},
				},
				FieldIndex: 12,
			},
		},
	},
}
