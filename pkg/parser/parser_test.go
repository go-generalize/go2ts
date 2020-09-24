package parser

import (
	"testing"

	tstypes "github.com/go-generalize/go2ts/pkg/types"
	"github.com/google/go-cmp/cmp"
)

func parse(t *testing.T, dir string, filter func(name string) bool) map[string]tstypes.Type {
	t.Helper()

	parser, err := NewParser(dir)

	if err != nil {
		t.Fatalf("failed to initialize parser: %+v", err)
	}
	parser.Filter = filter

	types, err := parser.Parse()

	if err != nil {
		t.Fatalf("failed to parse: %+v", err)
	}

	return types
}

func TestParser(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		types := parse(t, "./testdata/success", func(name string) bool {
			t.Log("checking export status: ", name)
			return name != "Unexported"
		})

		expected := map[string]tstypes.Type{
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

		if diff := cmp.Diff(expected, types); diff != "" {
			t.Errorf("unexpected: %s", diff)
		}
	})
}
