package parser

import (
	"reflect"
	"testing"

	tstypes "github.com/go-generalize/go2ts/pkg/types"
	"golang.org/x/tools/go/packages"
)

var (
	successParsedTypes = map[string]tstypes.Type{
		"github.com/go-generalize/go2ts/pkg/parser/testdata/success.Embedded": &tstypes.Object{
			Name: "Embedded",
			Entries: map[string]tstypes.ObjectEntry{
				"foo": {
					Optional: true,
					Type:     &tstypes.Number{},
				},
			},
		},
		"github.com/go-generalize/go2ts/pkg/parser/testdata/success.Status": &tstypes.String{
			Name: "Status",
			Enum: []string{"Failure", "OK"},
		},
		"github.com/go-generalize/go2ts/pkg/parser/testdata/success.Data": &tstypes.Object{
			Name: "Data",
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
				"Array": {
					Type: &tstypes.Nullable{
						Inner: &tstypes.Array{
							Inner: &tstypes.Number{},
						},
					},
				},
				"Status": {
					Type: &tstypes.String{
						Name: "Status",
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
)

func parse(t *testing.T, dir string) []*packages.Package {
	t.Helper()

	parser, err := NewParser(dir)

	if err != nil {
		t.Fatalf("failed to initialize parser: %+v", err)
	}

	return parser.pkgs
}

func TestParser_Parse(t *testing.T) {
	type fields struct {
		pkgs   []*packages.Package
		types  map[string]tstypes.Type
		consts map[string][]interface{}
		Filter func(name string) bool
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
				Filter: func(name string) bool {
					t.Log("checking export status: ", name)
					return name != "Unexported"
				},
			},
			wantRes: successParsedTypes,
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
				pkgs:   tt.fields.pkgs,
				types:  tt.fields.types,
				consts: tt.fields.consts,
				Filter: tt.fields.Filter,
			}
			gotRes, err := p.Parse()
			if (err != nil) != tt.wantErr {
				t.Errorf("Parser.Parse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err != nil {
				t.Logf("Parser.Parse() returned expected error = %v", err)
			}

			if !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("Parser.Parse() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}
