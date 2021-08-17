package success

import (
	"go/types"

	tstypes "github.com/go-generalize/go2ts/pkg/types"
)

var NormalType = map[string]tstypes.Type{
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
	"github.com/go-generalize/go2ts/pkg/parser/testdata/success/base.EmbeddedInt": &tstypes.Number{
		Name:    "github.com/go-generalize/go2ts/pkg/parser/testdata/success/base.EmbeddedInt",
		RawType: types.Int,
	},
	"github.com/go-generalize/go2ts/pkg/parser/testdata/success/base.Data": &tstypes.Object{
		Name: "github.com/go-generalize/go2ts/pkg/parser/testdata/success/base.Data",
		Entries: map[string]tstypes.ObjectEntry{
			"Time": {
				RawName: "Time",
				Type:    &tstypes.Date{},
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
							},
						},
					},
				},
			},
			"foo": {
				RawName:  "Foo",
				RawTag:   `json:"foo,omitempty"`,
				Optional: true,
				Type: &tstypes.Number{
					RawType: types.Int,
				},
			},
			"A": {
				RawName: "A",
				Type: &tstypes.Number{
					RawType: types.Int,
				},
			},
			"b": {
				RawName:  "B",
				RawTag:   `json:"b,omitempty"`,
				Optional: true,
				Type: &tstypes.Number{
					RawType: types.Int,
				},
			},
			"C": {
				RawName: "C",
				Type:    &tstypes.String{},
			},
			"D": {
				RawName: "D",
				Type: &tstypes.Nullable{
					Inner: &tstypes.Number{
						RawType: types.Float32,
					},
				},
			},
			"EmbeddedInt": {
				RawName: "EmbeddedInt",
				Type: &tstypes.Number{
					RawType: types.Int,
					Name:    "github.com/go-generalize/go2ts/pkg/parser/testdata/success/base.EmbeddedInt",
				},
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
						},
					},
				},
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
						},
					},
				},
			},
		},
	},
}
