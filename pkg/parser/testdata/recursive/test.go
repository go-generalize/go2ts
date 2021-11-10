package recursive

import (
	"github.com/go-generalize/go2ts/pkg/parser/testutil"
	tstypes "github.com/go-generalize/go2ts/pkg/types"
)

var Type = map[string]tstypes.Type{
	"github.com/go-generalize/go2ts/pkg/parser/testdata/recursive/base.Recursive": &tstypes.Object{
		Common: tstypes.Common{
			PkgName:  "recursive",
			Position: testutil.ParsePositionString("testdata/recursive/base/main.go:4:6"),
		},
		Name: "github.com/go-generalize/go2ts/pkg/parser/testdata/recursive/base.Recursive",
		Entries: map[string]tstypes.ObjectEntry{
			"Re": {}, // Overwritten by init()
		},
	},
}

func init() {
	//nolint
	re := Type["github.com/go-generalize/go2ts/pkg/parser/testdata/recursive/base.Recursive"].(*tstypes.Object)

	re.Entries["Re"] = tstypes.ObjectEntry{
		RawName: "Re",
		Type: &tstypes.Nullable{
			Inner: re,
		},
		Position: testutil.ParsePositionString("testdata/recursive/base/main.go:5:2"),
	}
}
