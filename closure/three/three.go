package three

import "fmt"

func Run() {
	ids := []int{1, 2, 3}
	fmt.Println("local copy - id shadows the loop variable wth a new one")
	for _, id := range ids {
		id := id
		go func() {
			fmt.Println("processing:", id)
		}()
	}

	fmt.Println("function argument - id is evaluated immediately when go is called")
	for _, id := range ids {
		go func(id int) {
			fmt.Println("processing:", id)
		}(id)
	}
}