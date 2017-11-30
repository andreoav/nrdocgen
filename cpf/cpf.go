package cpf

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

const (
	REMINDER_1 = 10 + iota
	REMINDER_2
)

// MASK of a valid CPF
const MASK = "%d%d%d.%d%d%d.%d%d%d-%d%d"

// CPF struct creates valid CPF numbers
// has an option to create raw numbers or formatted ones
type CPF struct {
	digits []int
}

// Generate creates a new valid CPF
func (c *CPF) Generate() *CPF {
	c.digits = c.randomDigits()

	c.addVerification(c.verification(REMINDER_1))
	c.addVerification(c.verification(REMINDER_2))

	return c
}

// Value returns a raw CPF number
func (c *CPF) Value() string {
	return c.joinNumbers(c.digits)
}

// Format returns a formatted CPF number
func (c *CPF) Format() string {
	var numbers = make([]interface{}, len(c.digits))

	for i, digit := range c.digits {
		numbers[i] = digit
	}

	return fmt.Sprintf(MASK, numbers...)
}

func (c *CPF) verification(base int) int {
	var sum int

	for index, digit := range c.digits {
		sum += (base - index) * digit
	}

	reminder := sum % 11

	if reminder < 2 {
		return 0
	}

	return 11 - reminder
}

func (c *CPF) addVerification(number int) {
	c.digits = append(c.digits, number)
}

func (c *CPF) joinNumbers(digits []int) string {
	var strNumber = []string{}

	for _, i := range digits {
		strNumber = append(strNumber, strconv.Itoa(i))
	}

	return strings.Join(strNumber, "")
}

func (c *CPF) randomDigits() []int {
	c.makeSeed()

	arr := make([]int, 9)

	for r := 0; r < 9; r++ {
		arr[r] = rand.Intn(9)
	}

	return arr
}

func (c *CPF) makeSeed() {
	rand.Seed(time.Now().UnixNano())
}

// New creates a new CPF Generator
func New() *CPF {
	return &CPF{}
}
