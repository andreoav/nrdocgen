package nrdocgen

// Documenter interface used in some generators
type Documenter interface {
	Value() string
	Format() string
	Generate() string
}

// Document struct used in some generators
type Document struct {
	Generator
	Mod11
}
