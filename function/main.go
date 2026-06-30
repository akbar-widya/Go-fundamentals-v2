package main

import (
	"fmt"
	"os"

	// Replace "myproject" with the actual module name found in your go.mod file
	"practice/one"
	"practice/two"
	"practice/three"
)

func main() {
	// Check if the user provided a number argument
	if len(os.Args) < 2 {
		fmt.Println("Please provide a package number (e.g., 'go run main.go 1')")
		return
	}

	choice := os.Args[1]

	switch choice {
	case "1":
		fmt.Println("--- Running Example One ---")
		one.Run()
	case "2":
		fmt.Println("--- Running Example Two ---")
		two.Run()
	case "3":
		fmt.Println("--- Running Example Three ---")
		three.Run()
	default:
		fmt.Println("Unknown choice:", choice)
	}
}
