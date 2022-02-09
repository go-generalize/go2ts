package testdata

import tstypes "github.com/go-generalize/go2ts/pkg/types"

var (
	// Test05 - 05.ts
	Test05 = map[string]tstypes.Type{
		"github.com/go-generalize/go2ts/pkg/parser/testdata.CustomTest": &tstypes.Object{
			Name: "github.com/go-generalize/go2ts/pkg/parser/testdata.CustomTest",
			Entries: map[string]tstypes.ObjectEntry{
				"C": {
					RawName:    "C",
					FieldIndex: 0,

					Type: &tstypes.Object{
						Name: "github.com/go-generalize/go2ts/pkg/parser/testdata.CustomTestC",
					},
				},
			},
		},
	}
)
