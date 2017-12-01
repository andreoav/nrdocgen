package nrdocgen_test

import (
	"testing"

	"github.com/andreoav/nrdocgen"
	"github.com/stretchr/testify/assert"
)

var mod11 = &nrdocgen.Mod11{}

func TestMod11_Generate(t *testing.T) {
	tests := []struct {
		input    int
		expected int
	}{
		{0, 0},
		{1, 0},
		{2, 9},
		{50, 5},
		{100, 0},
		{150, 4},
		{200, 9},
		{500, 6},
		{1000, 1},
	}
	for _, tt := range tests {
		assert.Equal(t, mod11.Generate(tt.input), tt.expected)
	}
}
