package parser

import (
	"go/types"
	"reflect"
	"testing"

	tstypes "github.com/go-generalize/go2ts/pkg/types"
	"golang.org/x/tools/go/packages"
)

func TestParser_parseType(t *testing.T) {
	type fields struct {
		pkgs  []*packages.Package
		types map[string]tstypes.Type
	}
	type args struct {
		u types.Type
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   tstypes.Type
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Parser{
				pkgs:  tt.fields.pkgs,
				types: tt.fields.types,
			}
			if got := p.parseType(tt.args.u); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Parser.parseType() = %v, want %v", got, tt.want)
			}
		})
	}
}
