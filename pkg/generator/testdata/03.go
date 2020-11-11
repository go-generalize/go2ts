package testdata

import (
	tstypes "github.com/go-generalize/go2ts/pkg/types"
)

var (
	// Data03 - 03.ts
	Data03 = map[string]tstypes.Type{
		"github.com/go-generalize/go2ts/pkg/parser/testdata/recursive.Recursive": &tstypes.Object{
			Name: "github.com/go-generalize/go2ts/pkg/parser/testdata/recursive.Recursive",
			Entries: map[string]tstypes.ObjectEntry{
				"Re": {}, // Overwritten by init()
			},
		},
	}
)

func init() {
	//nolint
	re := Data03["github.com/go-generalize/go2ts/pkg/parser/testdata/recursive.Recursive"].(*tstypes.Object)

	re.Entries["Re"] = tstypes.ObjectEntry{
		Type: &tstypes.Nullable{
			Inner: re,
		},
	}

}
