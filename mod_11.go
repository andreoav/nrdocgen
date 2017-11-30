package generators

// Mod11 struct
type Mod11 struct{}

// Generate creates a verification digit
func (m Mod11) Generate(sum int) int {
	reminder := sum % 11

	if reminder < 2 {
		return 0
	}

	return 11 - reminder
}
