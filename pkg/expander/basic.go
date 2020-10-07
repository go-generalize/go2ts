package expander

import (
	tstypes "github.com/go-generalize/go2ts/pkg/types"
)

func (g *expander) expandString(t *tstypes.String) *tstypes.String {
	return t
}

func (g *expander) expandNumber(t *tstypes.Number) *tstypes.Number {
	return t
}

func (g *expander) expandBoolean(t *tstypes.Boolean) *tstypes.Boolean {
	return t
}

func (g *expander) expandDate(t *tstypes.Date) *tstypes.Date {
	return t
}

func (g *expander) expandNullable(t *tstypes.Nullable) *tstypes.Nullable {
	return t
}

func (g *expander) expandAny(t *tstypes.Any) *tstypes.Any {
	return t
}
