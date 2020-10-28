// Package generator is a generator for TypeScript interfaces
package generator

import (
	"bytes"
	"strconv"
	"strings"

	tstypes "github.com/go-generalize/go2ts/pkg/types"
)

const (
	tsString  = "string"
	tsNumber  = "number"
	tsBoolean = "boolean"
	tsAny     = "any"
	tsNull    = "null"
)

func quoteAll(s []string) []string {
	res := make([]string, len(s))

	for i := range res {
		res[i] = strconv.Quote(s[i])
	}

	return res
}

func (g *Generator) generateString(t *tstypes.String) string {
	if len(t.Enum) != 0 {
		return strings.Join(quoteAll(t.Enum), " | ")
	}

	return tsString
}

func (g *Generator) generateNumber(t *tstypes.Number) string {
	if len(t.Enum) != 0 {
		buf := bytes.NewBuffer(nil)

		for i := range t.Enum {
			if i != 0 {
				buf.WriteString(" | ")
			}

			buf.WriteString(strconv.FormatInt(t.Enum[i], 10))
		}

		return buf.String()
	}

	return tsNumber
}

func (g *Generator) generateBoolean(t *tstypes.Boolean) string {
	return tsBoolean
}

func (g *Generator) generateDate(t *tstypes.Date) string {
	return tsString
}

func (g *Generator) generateNullable(t *tstypes.Nullable) string {
	return g.generateType(t.Inner) + " | " + tsNull
}

func (g *Generator) generateAny(t *tstypes.Any) string {
	return tsAny
}
