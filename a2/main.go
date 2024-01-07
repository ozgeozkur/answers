package main

import "fmt"

func recursiveFunction(n int) {
	// Your recursive logic goes here
	// This is where you implement the algorithm

	// Just for illustration, a simple recursive example is provided:
	if n > 0 {
		fmt.Println(n)
		recursiveFunction(n - 1)
	}
}

func main() {
	// Example usage:
	recursiveFunction(5)
}
