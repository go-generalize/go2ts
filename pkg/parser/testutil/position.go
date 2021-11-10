package testutil

import (
	"go/token"
	"strconv"
	"strings"
)

// ParsePositionString parses position notations generated by token.Position.String()
func ParsePositionString(s string) *token.Position {
	split := strings.SplitN(s, ":", 3)

	pos := token.Position{}
	switch len(split) {
	default:
		fallthrough
	case 3:
		pos.Column, _ = strconv.Atoi(split[2])
		fallthrough
	case 2:
		pos.Line, _ = strconv.Atoi(split[1])
		fallthrough
	case 1:
		pos.Filename = split[0]
	case 0:
		return nil
	}

	return &pos
}
