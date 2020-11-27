package testdata

import (
	tstypes "github.com/go-generalize/go2ts/pkg/types"
)

var (
	// Data01 - 01.ts
	Data01 = map[string]tstypes.Type{
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
				"EnumArray": {
					Type: &tstypes.Array{
						Inner: &tstypes.String{
							Enum: []string{"a", "b", "c"},
						},
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
				"OptionalArray": {
					Type: &tstypes.Array{
						Inner: &tstypes.Nullable{
							Inner: &tstypes.String{},
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
)
