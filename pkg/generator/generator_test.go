package generator

import (
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/go-generalize/go2ts/pkg/generator/testdata"
	tstypes "github.com/go-generalize/go2ts/pkg/types"
	"github.com/google/go-cmp/cmp"
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
				types:       testdata.Data01,
				altPkgs:     map[string]string{},
				BasePackage: "github.com/go-generalize/go2ts/pkg/parser/testdata/success",
			},
		},
		{
			name: "02",
			want: loadFile(t, "./testdata/02.ts"),
			fields: fields{
				types:       testdata.Data02,
				altPkgs:     map[string]string{},
				BasePackage: "github.com/go-generalize/go2ts/pkg/parser/testdata/conflict",
			},
		},
		{
			name: "03",
			want: loadFile(t, "./testdata/03.ts"),
			fields: fields{
				types:       testdata.Data03,
				altPkgs:     map[string]string{},
				BasePackage: "github.com/go-generalize/go2ts/pkg/parser/testdata/recursive",
			},
		},
		{
			name: "04",
			want: loadFile(t, "./testdata/04.ts"),
			fields: fields{
				types:       testdata.Data04,
				altPkgs:     map[string]string{},
				BasePackage: "github.com/go-generalize/go2ts/pkg/parser/testdata/recursive",
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
