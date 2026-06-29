package one

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

// ParseAmount is now exported so it can be accessed if needed
func ParseAmount(raw string) (int, error) {
	dollars, cents, found := strings.Cut(raw, ".")
	if !found {
		return 0, fmt.Errorf("invalid amount format: %q", raw)
	}
	d, err := strconv.Atoi(dollars)
	if err != nil {
		return 0, fmt.Errorf("invalid dollars: %w", err)
	}
	c, err := strconv.Atoi(cents)
	if err != nil {
		return 0, fmt.Errorf("invalid cents: %w", err)
	}
	return d*100 + c, nil
}

// Divide is now exported
func Divide(a, b int) (quotient, remainder int) {
	quotient = a / b
	remainder = a % b
	return
}

// Run replaces the old main() function
func Run() {
	amount, err := ParseAmount("19.99")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(amount)

	q, _ := Divide(17, 5)
	fmt.Println(q)
}