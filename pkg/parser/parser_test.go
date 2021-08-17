package parser

import (
	"fmt"
	"testing"

	"github.com/go-generalize/go2ts/pkg/parser/testdata/conflict"
	"github.com/go-generalize/go2ts/pkg/parser/testdata/recursive"
	"github.com/go-generalize/go2ts/pkg/parser/testdata/success"
	tstypes "github.com/go-generalize/go2ts/pkg/types"
	"github.com/google/go-cmp/cmp"
	"golang.org/x/tools/go/packages"
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
	base := "github.com/go-generalize/go2ts/pkg/parser/testdata/success"

	type fields struct {
		pkgs        []*packages.Package
		types       map[string]tstypes.Type
		consts      map[string][]constCandidate
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
				pkgs: parse(t, "./testdata/success/base"),
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
			wantRes: success.NormalType,
			wantErr: false,
		},
		{
			name: "all exported",
			fields: fields{
				pkgs:        parse(t, "./testdata/success/base"),
				basePackage: base,
				Filter:      All,
			},
			wantRes: success.AllExportedType,
			wantErr: false,
		},
		{
			name: "recursive",
			fields: fields{
				pkgs:        parse(t, "./testdata/recursive/base"),
				basePackage: base,
				Filter:      All,
			},
			wantRes: recursive.Type,
			wantErr: false,
		},
		{
			name: "conflicting",
			fields: fields{
				pkgs:        parse(t, "./testdata/conflict/base"),
				basePackage: base,
				Filter:      All,
			},
			wantRes: conflict.Type,
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
