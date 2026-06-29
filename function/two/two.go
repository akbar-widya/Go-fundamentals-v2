package two

import "fmt"

func double(n int) int { return n * 2 }

func Run() {
	var transform func(int) int = double
	fmt.Println(transform(21))
}
