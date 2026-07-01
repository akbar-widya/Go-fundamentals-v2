package four

import "fmt"

// newPrefixer returns a configured fun(string) string - the prefix is
// captured once, at creation, and reused on every call.
func newPrefixer(prefix string) func(string) string {
	return func(s string) string {
		return prefix + s
	}
}

// newLengthValidator returns a func(string) error configured with a
// specific rule, without the caller needing to know what that rule is.
func newLengthValidator(max int) func(string) error {
	return func(s string) error {
		if len(s) > max {
			return fmt.Errorf("exceeds %d characters", max)
		}
		return nil
	}
}

func Run() {
	logPrefix := newPrefixer("[INFO] ")
	fmt.Println(logPrefix("server started"))

	validateUsername := newLengthValidator(20)
	fmt.Println(validateUsername("alice"))
}