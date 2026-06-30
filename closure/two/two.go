package two

import (
	"fmt"
	"sync"
)

// 1. Closures creates a bug in a loop (Turns out it got fix post Go 1.22)
func dispatchBuggy(ids []int) {
	var wg sync.WaitGroup
	for _, id := range ids {
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Println("processing:", id) // problem: capture sthe loop variable itself
		}()
	}
	wg.Wait()
}

func Run() {
	ids := []int{1, 2, 3, 4, 5}

	fmt.Println("Running buggy dispatch:")
	dispatchBuggy(ids)
}
