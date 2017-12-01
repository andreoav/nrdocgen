package cpf

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/andreoav/nrdocgen"
)

const (
	// FACTOR1 used for the first verification number
	FACTOR1 = 10 + iota

	// FACTOR2 used for the second verification number
	FACTOR2

	// SIZE of the slice of digits
	SIZE = 9

	// MIN digit allowed
	MIN = 0

	// MAX digit allowed
	MAX = 9
)

// MASK CPF for formatting output
const MASK = "%d%d%d.%d%d%d.%d%d%d-%d%d"

// CPF struct creates valid CPF numbers and has an
// option to create raw numbers or formatted ones
type CPF struct {
	nrdocgen.Document
	digits []int
}

// Generate creates a new valid CPF
func (c *CPF) Generate() *CPF {
	c.digits = c.RandomDigits(SIZE, MIN, MAX)

	c.addVerification(c.verification(FACTOR1))
	c.addVerification(c.verification(FACTOR2))

	return c
}

// Value returns a raw CPF number
func (c CPF) Value() string {
	return c.joinNumbers(c.digits)
}

// Format returns a formatted CPF number
func (c CPF) Format() string {
	var numbers = make([]interface{}, len(c.digits))

	for i := range c.digits {
		numbers[i] = c.digits[i]
	}

	return fmt.Sprintf(MASK, numbers...)
}

func (c CPF) verification(base int) int {
	var sum int

	for i := range c.digits {
		sum += (base - i) * c.digits[i]
	}

	return c.Mod11.Generate(sum)
}

func (c *CPF) addVerification(number int) {
	c.digits = append(c.digits, number)
}

func (c CPF) joinNumbers(digits []int) string {
	var strNumber = []string{}

	for i := range digits {
		strNumber = append(strNumber, strconv.Itoa(digits[i]))
	}

	return strings.Join(strNumber, "")
}

// New creates a new CPF Generator
func New() *CPF {
	return &CPF{}
}
