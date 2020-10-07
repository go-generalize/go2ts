package generator

import (
	"fmt"
	"io/ioutil"
	"testing"

	tstypes "github.com/go-generalize/go2ts/pkg/types"
	"github.com/google/go-cmp/cmp"
)

var (
	successParsedTypes = map[string]tstypes.Type{
		"github.com/go-generalize/go2ts/pkg/parser/testdata/success.Embedded": &tstypes.Object{
			Name: "github.com/go-generalize/go2ts/pkg/parser/testdata/success.Embedded",
			Entries: map[string]tstypes.ObjectEntry{
				"foo": {
					Optional: true,
					Type:     &tstypes.Number{},
				},
			},
		},
		"github.com/go-generalize/go2ts/pkg/parser/testdata/success.Status": &tstypes.String{
			Name: "github.com/go-generalize/go2ts/pkg/parser/testdata/success.Status",
			Enum: []string{"Failure", "OK"},
		},
		"github.com/go-generalize/go2ts/pkg/parser/testdata/success.Data": &tstypes.Object{
			Name: "github.com/go-generalize/go2ts/pkg/parser/testdata/success.Data",
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
				"Map": {
					Type: &tstypes.Map{
						Key: &tstypes.String{},
						Value: &tstypes.String{
							Name: "github.com/go-generalize/go2ts/pkg/parser/testdata/success.Status",
							Enum: []string{"Failure", "OK"},
						},
					},
				},
				"Status": {
					Type: &tstypes.String{
						Name: "github.com/go-generalize/go2ts/pkg/parser/testdata/success.Status",
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

	conflictingData = map[string]tstypes.Type{
		"github.com/go-generalize/go2ts/pkg/parser/testdata/conflict.Data": &tstypes.Object{
			Name: "github.com/go-generalize/go2ts/pkg/parser/testdata/conflict.Data",
			Entries: map[string]tstypes.ObjectEntry{
				"Hoge": {
					Type: &tstypes.Object{
						Name: "github.com/go-generalize/go2ts/pkg/parser/testdata/conflict.Hoge",
						Entries: map[string]tstypes.ObjectEntry{
							"Data": {
								Type: &tstypes.Number{},
							},
						},
					},
				},
				"PkgHoge": {
					Type: &tstypes.Object{
						Name: "github.com/go-generalize/go2ts/pkg/parser/testdata/conflict/pkg.Hoge",
						Entries: map[string]tstypes.ObjectEntry{
							"Data": {
								Type: &tstypes.Number{},
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
					Type: &tstypes.Number{},
				},
			},
		},
		"github.com/go-generalize/go2ts/pkg/parser/testdata/conflict/pkg.Hoge": &tstypes.Object{
			Name: "github.com/go-generalize/go2ts/pkg/parser/testdata/conflict/pkg.Hoge",
			Entries: map[string]tstypes.ObjectEntry{
				"Data": {
					Type: &tstypes.Number{},
				},
			},
		},
	}
)

func loadFile(t *testing.T, name string) string {
	t.Helper()

	b, err := ioutil.ReadFile(name)
	if err != nil {
		t.Fatalf("failed to open file(%s): %+v", name, err)
	}

	return string(b)
}

func TestGenerator_Generate(t *testing.T) {
	type fields struct {
		types       map[string]tstypes.Type
		altPkgs     map[string]string
		BasePackage string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "01",
			want: loadFile(t, "./testdata/01.ts"),
			fields: fields{
				types:       successParsedTypes,
				altPkgs:     map[string]string{},
				BasePackage: "github.com/go-generalize/go2ts/pkg/parser/testdata/success",
			},
		},
		{
			name: "02",
			want: loadFile(t, "./testdata/02.ts"),
			fields: fields{
				types:       conflictingData,
				altPkgs:     map[string]string{},
				BasePackage: "github.com/go-generalize/go2ts/pkg/parser/testdata/conflict",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &Generator{
				types:       tt.fields.types,
				BasePackage: tt.fields.BasePackage,
				altPkgs:     tt.fields.altPkgs,
			}
			got := g.Generate()
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("Generator.Generate() differed: %s", diff)

				fmt.Println(got)
				fmt.Println(tt.want)
			}
		})
	}
}
