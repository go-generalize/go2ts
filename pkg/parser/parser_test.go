package parser

import (
	"fmt"
	"go/types"
	"testing"

	tstypes "github.com/go-generalize/go2ts/pkg/types"
	"github.com/google/go-cmp/cmp"
	"golang.org/x/tools/go/packages"
)

var (
	successParsedTypes = map[string]tstypes.Type{
		"github.com/go-generalize/go2ts/pkg/parser/testdata/success.Embedded": &tstypes.Object{
			Name: "github.com/go-generalize/go2ts/pkg/parser/testdata/success.Embedded",
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
		"github.com/go-generalize/go2ts/pkg/parser/testdata/success.Status": &tstypes.String{
			Name: "github.com/go-generalize/go2ts/pkg/parser/testdata/success.Status",
			Enum: []string{"Failure", "OK"},
		},
		"github.com/go-generalize/go2ts/pkg/parser/testdata/success.EmbeddedInt": &tstypes.Number{
			Name:    "github.com/go-generalize/go2ts/pkg/parser/testdata/success.EmbeddedInt",
			RawType: types.Int,
		},
		"github.com/go-generalize/go2ts/pkg/parser/testdata/success.Data": &tstypes.Object{
			Name: "github.com/go-generalize/go2ts/pkg/parser/testdata/success.Data",
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
						Name:    "github.com/go-generalize/go2ts/pkg/parser/testdata/success.EmbeddedInt",
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
							Name: "github.com/go-generalize/go2ts/pkg/parser/testdata/success.Status",
							Enum: []string{"Failure", "OK"},
						},
					},
				},
				"Status": {
					RawName: "Status",
					Type: &tstypes.String{
						Name: "github.com/go-generalize/go2ts/pkg/parser/testdata/success.Status",
						Enum: []string{"Failure", "OK"},
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

	successAllExportedParsedTypes = map[string]tstypes.Type{
		"github.com/go-generalize/go2ts/pkg/parser/testdata/success.Embedded": &tstypes.Object{
			Name: "github.com/go-generalize/go2ts/pkg/parser/testdata/success.Embedded",
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
		"github.com/go-generalize/go2ts/pkg/parser/testdata/success.Status": &tstypes.String{
			Name: "github.com/go-generalize/go2ts/pkg/parser/testdata/success.Status",
			Enum: []string{"Failure", "OK"},
		},
		"github.com/go-generalize/go2ts/pkg/parser/testdata/success.foo": &tstypes.Object{
			Name: "github.com/go-generalize/go2ts/pkg/parser/testdata/success.foo",
			Entries: map[string]tstypes.ObjectEntry{
				"V": {
					RawName: "V",
					Type: &tstypes.Number{
						RawType: types.Int,
					},
				},
			},
		},
		"github.com/go-generalize/go2ts/pkg/parser/testdata/success/pkg.Package": &tstypes.Object{
			Name: "github.com/go-generalize/go2ts/pkg/parser/testdata/success/pkg.Package",
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
		"github.com/go-generalize/go2ts/pkg/parser/testdata/success.Unexported": &tstypes.Object{
			Name: "github.com/go-generalize/go2ts/pkg/parser/testdata/success.Unexported",
			Entries: map[string]tstypes.ObjectEntry{
				"Data": {
					RawName: "Data",
					Type: &tstypes.Number{
						RawType: types.Int,
					},
				},
			},
		},
		"github.com/go-generalize/go2ts/pkg/parser/testdata/success.EmbeddedInt": &tstypes.Number{
			Name:    "github.com/go-generalize/go2ts/pkg/parser/testdata/success.EmbeddedInt",
			RawType: types.Int,
		},
		"github.com/go-generalize/go2ts/pkg/parser/testdata/success.Data": &tstypes.Object{
			Name: "github.com/go-generalize/go2ts/pkg/parser/testdata/success.Data",
			Entries: map[string]tstypes.ObjectEntry{
				"Time": {
					RawName: "Time",
					Type:    &tstypes.Date{},
				},
				"Package": {
					RawName: "Package",
					Type: &tstypes.Nullable{
						Inner: &tstypes.Object{
							Name: "github.com/go-generalize/go2ts/pkg/parser/testdata/success/pkg.Package",
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
						Name:    "github.com/go-generalize/go2ts/pkg/parser/testdata/success.EmbeddedInt",
						RawType: types.Int,
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
							Name: "github.com/go-generalize/go2ts/pkg/parser/testdata/success.Status",
							Enum: []string{"Failure", "OK"},
						},
					},
				},
				"Status": {
					RawName: "Status",
					Type: &tstypes.String{
						Name: "github.com/go-generalize/go2ts/pkg/parser/testdata/success.Status",
						Enum: []string{"Failure", "OK"},
					},
				},
				"Foo": {
					RawName:  "Foo",
					RawTag:   `json:",omitempty"`,
					Optional: true,
					Type: &tstypes.Object{
						Name: "github.com/go-generalize/go2ts/pkg/parser/testdata/success.foo",
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
						Name: "github.com/go-generalize/go2ts/pkg/parser/testdata/success.Unexported",
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

	recursiveData = map[string]tstypes.Type{
		"github.com/go-generalize/go2ts/pkg/parser/testdata/recursive.Recursive": &tstypes.Object{
			Name: "github.com/go-generalize/go2ts/pkg/parser/testdata/recursive.Recursive",
			Entries: map[string]tstypes.ObjectEntry{
				"Re": {}, // Overwritten by init()
			},
		},
	}

	conflictingData = map[string]tstypes.Type{
		"github.com/go-generalize/go2ts/pkg/parser/testdata/conflict.Data": &tstypes.Object{
			Name: "github.com/go-generalize/go2ts/pkg/parser/testdata/conflict.Data",
			Entries: map[string]tstypes.ObjectEntry{
				"Hoge": {
					RawName: "Hoge",
					Type: &tstypes.Object{
						Name: "github.com/go-generalize/go2ts/pkg/parser/testdata/conflict.Hoge",
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
				"PkgHoge": {
					RawName: "PkgHoge",
					Type: &tstypes.Object{
						Name: "github.com/go-generalize/go2ts/pkg/parser/testdata/conflict/pkg.Hoge",
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
		"github.com/go-generalize/go2ts/pkg/parser/testdata/conflict.Hoge": &tstypes.Object{
			Name: "github.com/go-generalize/go2ts/pkg/parser/testdata/conflict.Hoge",
			Entries: map[string]tstypes.ObjectEntry{
				"Data": {
					RawName: "Data",
					Type: &tstypes.Number{
						RawType: types.Int,
					},
				},
			},
		},
		"github.com/go-generalize/go2ts/pkg/parser/testdata/conflict/pkg.Hoge": &tstypes.Object{
			Name: "github.com/go-generalize/go2ts/pkg/parser/testdata/conflict/pkg.Hoge",
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
)

func init() {
	//nolint
	re := recursiveData["github.com/go-generalize/go2ts/pkg/parser/testdata/recursive.Recursive"].(*tstypes.Object)

	re.Entries["Re"] = tstypes.ObjectEntry{
		RawName: "Re",
		Type: &tstypes.Nullable{
			Inner: re,
		},
	}
}

func parse(t *testing.T, dir string) []*packages.Package {
	t.Helper()

	parser, err := NewParser(dir, All)

	if err != nil {
		t.Fatalf("failed to initialize parser: %+v", err)
	}

	return parser.pkgs
}

func TestParser_Parse(t *testing.T) {
	base := "github.com/go-generalize/go2ts/pkg/parser/testdata/success"

	type fields struct {
		pkgs        []*packages.Package
		types       map[string]tstypes.Type
		consts      map[string][]interface{}
		basePackage string
		Filter      func(opt *FilterOpt) bool
	}
	tests := []struct {
		name    string
		fields  fields
		wantRes map[string]tstypes.Type
		wantErr bool
	}{
		{
			name: "success",
			fields: fields{
				pkgs: parse(t, "./testdata/success"),
				Filter: func(opt *FilterOpt) bool {
					fmt.Println(opt)
					if !Default(opt) {
						return false
					}
					fmt.Println("ok")

					t.Log("checking export status: ", opt.Name)
					return opt.Name != "Unexported"
				},
				basePackage: base,
			},
			wantRes: successParsedTypes,
			wantErr: false,
		},
		{
			name: "all exported",
			fields: fields{
				pkgs:        parse(t, "./testdata/success"),
				basePackage: base,
				Filter:      All,
			},
			wantRes: successAllExportedParsedTypes,
			wantErr: false,
		},
		{
			name: "recursive",
			fields: fields{
				pkgs:        parse(t, "./testdata/recursive"),
				basePackage: base,
				Filter:      All,
			},
			wantRes: recursiveData,
			wantErr: false,
		},
		{
			name: "conflicting",
			fields: fields{
				pkgs:        parse(t, "./testdata/conflict"),
				basePackage: base,
				Filter:      All,
			},
			wantRes: conflictingData,
			wantErr: false,
		},
		{
			name: "unsupported_map_key",
			fields: fields{
				pkgs: parse(t, "./testdata/unsupported_map_key"),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Parser{
				pkgs:        tt.fields.pkgs,
				types:       tt.fields.types,
				consts:      tt.fields.consts,
				basePackage: tt.fields.basePackage,
				Filter:      tt.fields.Filter,
			}
			gotRes, err := p.Parse()
			if (err != nil) != tt.wantErr {
				t.Errorf("Parser.Parse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err != nil {
				t.Logf("Parser.Parse() returned expected error = %v", err)
			}

			if diff := cmp.Diff(tt.wantRes, gotRes); diff != "" {
				t.Errorf("Parser.Parse() differed: %s", diff)
			}
		})
	}
}
