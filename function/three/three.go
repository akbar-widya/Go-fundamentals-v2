package three

import (
	"errors"
	"fmt"
)

// ValidatorFunc names a common shape, so a function that accepts
// validation logic reads clearly instead of spelling out func(string) error.
type ValidatorFunc func(string) error

func notEmpty(s string) error {
	if s == "" {
		return errors.New("value cannot be empty")
	}
	return nil
}

func maxLength(n int) ValidatorFunc {
	return func(s string) error {
		if len(s) > n {
			return fmt.Errorf("value exceeds %d characters", n)
		}
		return nil
	}
}

// runValidators is a higher-order function: it accepts functions as parameters
// inside the function, validators has type []ValidatorFunc.
func runValidators(value string, validators ...ValidatorFunc) error {
	for _, validate := range validators {
		if err := validate(value); err != nil {
			return err
		}
	}
	return nil
}

func Run() {
	err := runValidators("alice whitetaker", notEmpty, maxLength(10))
	fmt.Println(err)
}
