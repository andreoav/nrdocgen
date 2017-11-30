package generators

// Document interface
type Document interface {
	Value() string
	Format() string
	Generate() string
}
