package parser

import (
	"fmt"
	"testing"

	tstypes "github.com/go-generalize/go2ts/v2/pkg/types"
	"github.com/google/go-cmp/cmp"
	"golang.org/x/tools/go/packages"
)

var (
	successParsedTypes = map[string]tstypes.Type{
		"github.com/go-generalize/go2ts/v2/pkg/parser/testdata/success.Embedded": &tstypes.Object{
			Name: "github.com/go-generalize/go2ts/v2/pkg/parser/testdata/success.Embedded",
			Entries: map[string]tstypes.ObjectEntry{
				"foo": {
					Optional: true,
					Type:     &tstypes.Number{},
				},
			},
		},
		"github.com/go-generalize/go2ts/v2/pkg/parser/testdata/success.Status": &tstypes.String{
			Name: "github.com/go-generalize/go2ts/v2/pkg/parser/testdata/success.Status",
			Enum: []string{"Failure", "OK"},
		},
		"github.com/go-generalize/go2ts/v2/pkg/parser/testdata/success.EmbeddedInt": &tstypes.Number{
			Name: "github.com/go-generalize/go2ts/v2/pkg/parser/testdata/success.EmbeddedInt",
		},
		"github.com/go-generalize/go2ts/v2/pkg/parser/testdata/success.Data": &tstypes.Object{
			Name: "github.com/go-generalize/go2ts/v2/pkg/parser/testdata/success.Data",
			Entries: map[string]tstypes.ObjectEntry{
				"Time": {
					Type: &tstypes.Date{},
				},
				"Package": {
					Type: &tstypes.Nullable{
						Inner: &tstypes.Object{
							Entries: map[string]tstypes.ObjectEntry{
								"data": {
									Type: &tstypes.Number{},
								},
							},
						},
					},
				},
				"foo": {
					Optional: true,
					Type:     &tstypes.Number{},
				},
				"A": {
					Type: &tstypes.Number{},
				},
				"b": {
					Optional: true,
					Type:     &tstypes.Number{},
				},
				"C": {
					Type: &tstypes.String{},
				},
				"D": {
					Type: &tstypes.Nullable{
						Inner: &tstypes.Number{},
					},
				},
				"EmbeddedInt": {
					Type: &tstypes.Number{
						Name: "github.com/go-generalize/go2ts/v2/pkg/parser/testdata/success.EmbeddedInt",
					},
				},
				"Array": {
					Type: &tstypes.Nullable{
						Inner: &tstypes.Array{
							Inner: &tstypes.Number{},
						},
					},
				},
				"Map": {
					Type: &tstypes.Map{
						Key: &tstypes.String{},
						Value: &tstypes.String{
							Name: "github.com/go-generalize/go2ts/v2/pkg/parser/testdata/success.Status",
							Enum: []string{"Failure", "OK"},
						},
					},
				},
				"Status": {
					Type: &tstypes.String{
						Name: "github.com/go-generalize/go2ts/v2/pkg/parser/testdata/success.Status",
						Enum: []string{"Failure", "OK"},
					},
				},
				"Foo": {
					Optional: true,
					Type: &tstypes.Object{
						Entries: map[string]tstypes.ObjectEntry{
							"V": {
								Type: &tstypes.Number{},
							},
						},
					},
				},
				"U": {
					Type: &tstypes.Object{
						Entries: map[string]tstypes.ObjectEntry{
							"Data": {
								Type: &tstypes.Number{},
							},
						},
					},
				},
			},
		},
	}

	successAllExportedParsedTypes = map[string]tstypes.Type{
		"github.com/go-generalize/go2ts/v2/pkg/parser/testdata/success.Embedded": &tstypes.Object{
			Name: "github.com/go-generalize/go2ts/v2/pkg/parser/testdata/success.Embedded",
			Entries: map[string]tstypes.ObjectEntry{
				"foo": {
					Optional: true,
					Type:     &tstypes.Number{},
				},
			},
		},
		"github.com/go-generalize/go2ts/v2/pkg/parser/testdata/success.Status": &tstypes.String{
			Name: "github.com/go-generalize/go2ts/v2/pkg/parser/testdata/success.Status",
			Enum: []string{"Failure", "OK"},
		},
		"github.com/go-generalize/go2ts/v2/pkg/parser/testdata/success.foo": &tstypes.Object{
			Name: "github.com/go-generalize/go2ts/v2/pkg/parser/testdata/success.foo",
			Entries: map[string]tstypes.ObjectEntry{
				"V": {
					Type: &tstypes.Number{},
				},
			},
		},
		"github.com/go-generalize/go2ts/v2/pkg/parser/testdata/success/pkg.Package": &tstypes.Object{
			Name: "github.com/go-generalize/go2ts/v2/pkg/parser/testdata/success/pkg.Package",
			Entries: map[string]tstypes.ObjectEntry{
				"data": {
					Type: &tstypes.Number{},
				},
			},
		},
		"github.com/go-generalize/go2ts/v2/pkg/parser/testdata/success.Unexported": &tstypes.Object{
			Name: "github.com/go-generalize/go2ts/v2/pkg/parser/testdata/success.Unexported",
			Entries: map[string]tstypes.ObjectEntry{
				"Data": {
					Type: &tstypes.Number{},
				},
			},
		},
		"github.com/go-generalize/go2ts/v2/pkg/parser/testdata/success.EmbeddedInt": &tstypes.Number{
			Name: "github.com/go-generalize/go2ts/v2/pkg/parser/testdata/success.EmbeddedInt",
		},
		"github.com/go-generalize/go2ts/v2/pkg/parser/testdata/success.Data": &tstypes.Object{
			Name: "github.com/go-generalize/go2ts/v2/pkg/parser/testdata/success.Data",
			Entries: map[string]tstypes.ObjectEntry{
				"Time": {
					Type: &tstypes.Date{},
				},
				"Package": {
					Type: &tstypes.Nullable{
						Inner: &tstypes.Object{
							Name: "github.com/go-generalize/go2ts/v2/pkg/parser/testdata/success/pkg.Package",
							Entries: map[string]tstypes.ObjectEntry{
								"data": {
									Type: &tstypes.Number{},
								},
							},
						},
					},
				},
				"foo": {
					Optional: true,
					Type:     &tstypes.Number{},
				},
				"A": {
					Type: &tstypes.Number{},
				},
				"b": {
					Optional: true,
					Type:     &tstypes.Number{},
				},
				"C": {
					Type: &tstypes.String{},
				},
				"D": {
					Type: &tstypes.Nullable{
						Inner: &tstypes.Number{},
					},
				},
				"EmbeddedInt": {
					Type: &tstypes.Number{
						Name: "github.com/go-generalize/go2ts/v2/pkg/parser/testdata/success.EmbeddedInt",
					},
				},
				"Array": {
					Type: &tstypes.Nullable{
						Inner: &tstypes.Array{
							Inner: &tstypes.Number{},
						},
					},
				},
				"Map": {
					Type: &tstypes.Map{
						Key: &tstypes.String{},
						Value: &tstypes.String{
							Name: "github.com/go-generalize/go2ts/v2/pkg/parser/testdata/success.Status",
							Enum: []string{"Failure", "OK"},
						},
					},
				},
				"Status": {
					Type: &tstypes.String{
						Name: "github.com/go-generalize/go2ts/v2/pkg/parser/testdata/success.Status",
						Enum: []string{"Failure", "OK"},
					},
				},
				"Foo": {
					Optional: true,
					Type: &tstypes.Object{
						Name: "github.com/go-generalize/go2ts/v2/pkg/parser/testdata/success.foo",
						Entries: map[string]tstypes.ObjectEntry{
							"V": {
								Type: &tstypes.Number{},
							},
						},
					},
				},
				"U": {
					Type: &tstypes.Object{
						Name: "github.com/go-generalize/go2ts/v2/pkg/parser/testdata/success.Unexported",
						Entries: map[string]tstypes.ObjectEntry{
							"Data": {
								Type: &tstypes.Number{},
							},
						},
					},
				},
			},
		},
	}

	conflictingData = map[string]tstypes.Type{
		"github.com/go-generalize/go2ts/v2/pkg/parser/testdata/conflict.Data": &tstypes.Object{
			Name: "github.com/go-generalize/go2ts/v2/pkg/parser/testdata/conflict.Data",
			Entries: map[string]tstypes.ObjectEntry{
				"Hoge": {
					Type: &tstypes.Object{
						Name: "github.com/go-generalize/go2ts/v2/pkg/parser/testdata/conflict.Hoge",
						Entries: map[string]tstypes.ObjectEntry{
							"Data": {
								Type: &tstypes.Number{},
							},
						},
					},
				},
				"PkgHoge": {
					Type: &tstypes.Object{
						Name: "github.com/go-generalize/go2ts/v2/pkg/parser/testdata/conflict/pkg.Hoge",
						Entries: map[string]tstypes.ObjectEntry{
							"Data": {
								Type: &tstypes.Number{},
							},
						},
					},
				},
			},
		},
		"github.com/go-generalize/go2ts/v2/pkg/parser/testdata/conflict.Hoge": &tstypes.Object{
			Name: "github.com/go-generalize/go2ts/v2/pkg/parser/testdata/conflict.Hoge",
			Entries: map[string]tstypes.ObjectEntry{
				"Data": {
					Type: &tstypes.Number{},
				},
			},
		},
		"github.com/go-generalize/go2ts/v2/pkg/parser/testdata/conflict/pkg.Hoge": &tstypes.Object{
			Name: "github.com/go-generalize/go2ts/v2/pkg/parser/testdata/conflict/pkg.Hoge",
			Entries: map[string]tstypes.ObjectEntry{
				"Data": {
					Type: &tstypes.Number{},
				},
			},
		},
	}
)

func parse(t *testing.T, dir string) []*packages.Package {
	t.Helper()

	parser, err := NewParser(dir, All)

	if err != nil {
		t.Fatalf("failed to initialize parser: %+v", err)
	}

	return parser.pkgs
}

func TestParser_Parse(t *testing.T) {
	base := "github.com/go-generalize/go2ts/v2/pkg/parser/testdata/success"

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
