package nrdocgen

import (
	"math/rand"
	"time"
)

// Generator struct used in some generators
type Generator struct{}

// RandomDigits creates a slice of digits using
// an allowed range of digits and a maximum size
func (d Generator) RandomDigits(size, min, max int) []int {
	arr := make([]int, size)

	for r := 0; r < size; r++ {
		arr[r] = rand.Intn(max) - min
	}

	return arr
}

func init() {
	rand.Seed(time.Now().UnixNano())
}
