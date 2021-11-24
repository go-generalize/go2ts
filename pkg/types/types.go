// Package types contains structs/interfaces representing TypeScript types
package types

import (
	"github.com/go-generalize/go-easyparser/types"
)

// Common defines common fields in the all types
type Common = types.Common

// Type interface represents all TypeScript types handled by go2ts
type Type = types.Type

// Enumerable interface represents union types
type Enumerable = types.Enumerable

// NamedType interface represents named types
type NamedType = types.NamedType
