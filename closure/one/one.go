package one

import "fmt"

func newCounterPair() (increment func() int, current func() int) {
	count := 0		// captured by reference by both closures below - not copied into either

	increment = func() int {
		count++
		return count
	}
	current = func() int {
		return count
	}
	return increment, current
}

func Run() {
	inc, get := newCounterPair()

	inc()
	inc()
	fmt.Println(get())	// both closures share the same count

	inc()
	fmt.Println(get())
}