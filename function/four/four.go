package four

import (
	"fmt"
	"log"
)

type User struct {
	Name string
	Age  int
}

func (u User) Validate() error {
	if u.Age < 0 {
		return fmt.Errorf("invalid age for %s", u.Name)
	}
	return nil
}

func Run() {
	u := User{Name: "bob", Age: -1}

	validate := u.Validate

	if err := validate(); err != nil {
		log.Println(err)
	}

	checks := []func() error{validate}
	for _, check := range checks {
		_ = check()
	}
}
