package generator

import (
	"fmt"
	"io/ioutil"
	"testing"

	tstypes "github.com/go-generalize/go2ts/pkg/types"
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
		types map[string]tstypes.Type
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
				types: map[string]tstypes.Type{
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
							"Map": {
								Type: &tstypes.Map{
									Key: &tstypes.String{},
									Value: &tstypes.String{
										Enum: []string{"Foo", "Bar"},
									},
								},
							},
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &Generator{
				types: tt.fields.types,
			}
			if got := g.Generate(); got != tt.want {
				t.Errorf("Generator.Generate() = %v, want %v", got, tt.want)

				fmt.Println(got)
			}
		})
	}
}
