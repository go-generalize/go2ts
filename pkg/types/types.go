// Package types contains structs/interfaces representing TypeScript types
package types

// Common defines common fields in the all types
type Common struct {
	// PkgName is the package name declared at the beggining of .go files.
	// Currently, only exported types in the root package is available.
	PkgName string
}

// SetPackageName sets PkgName in Common
func (c *Common) SetPackageName(pkgName string) {
	c.PkgName = pkgName
}

// GetPackageName returns PkgName in Common
func (c *Common) GetPackageName() string {
	return c.PkgName
}

// Type interface represents all TypeScript types handled by go2ts
type Type interface {
	SetPackageName(pkgName string)
	GetPackageName() string
	UsedAsMapKey() bool
	String() string
}

// Enumerable interface represents union types
type Enumerable interface {
	Type

	// AddCandidates adds a candidate for enum
	AddCandidates(key string, v interface{})
}

// NamedType interface represents named types
type NamedType interface {
	Type

	// SetName sets an alternative name
	SetName(name string)
}
